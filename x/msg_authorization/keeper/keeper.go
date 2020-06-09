package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler
	router   sdk.Router
}

// NewKeeper constructs a message authorization Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.Marshaler, router sdk.Router) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		router:   router,
	}
}

func (k Keeper) getActorAuthorizationKey(grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) []byte {
	return []byte(fmt.Sprintf("c/%x/%x/%s", grantee, granter, msgType))
}

func (k Keeper) getAuthorizationGrant(ctx sdk.Context, actor []byte) (grant types.AuthorizationGrant, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(actor)
	if bz == nil {
		return grant, false
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &grant)
	return grant, true
}

func (k Keeper) update(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, updated *codectypes.Any) {
	var authorization types.AuthorizationI
	err := k.cdc.UnpackAny(updated, &authorization)
	if err != nil {
		return
	}

	actor := k.getActorAuthorizationKey(grantee, granter, authorization.MsgType())
	grant, found := k.getAuthorizationGrant(ctx, actor)
	if !found {
		return
	}

	grant.Authorization = updated
	store := ctx.KVStore(k.storeKey)
	store.Set(actor, k.cdc.MustMarshalBinaryBare(&grant))
}

// DispatchActions attempts to execute the provided messages via authorization
// grants from the message signer to the grantee.
func (k Keeper) DispatchActions(ctx sdk.Context, grantee sdk.AccAddress, msgs []sdk.Msg) (*sdk.Result, error) {
	var msgResult *sdk.Result
	var err error
	for _, msg := range msgs {
		signers := msg.GetSigners()
		if len(signers) != 1 {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "authorization can be given to msg with only one signer")
		}
		granter := signers[0]
		if !bytes.Equal(granter, grantee) {
			authorization, _ := k.GetAuthorization(ctx, grantee, granter, msg.Type())
			if authorization == nil {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "authorization not found")
			}
			allow, updated, del := authorization.Accept(msg, ctx.BlockHeader())
			if !allow {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "authorization not found")
			}
			if del {
				k.Revoke(ctx, grantee, granter, msg.Type())
			} else if updated != nil {
				k.update(ctx, grantee, granter, updated)
			}
		}
		handler := k.router.Route(ctx, msg.Route())
		if handler == nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message route: %s", msg.Route())
		}

		msgResult, err = handler(ctx, msg)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "failed to execute message; message %s", msg.Type())
		}
	}

	return msgResult, nil
}

// Grant method grants the provided authorization to the grantee on the granter's account with the provided expiration
// time. If there is an existing authorization grant for the same `sdk.Msg` type, this grant
// overwrites that.
func (k Keeper) Grant(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, authorization *codectypes.Any, expiration int64) {
	store := ctx.KVStore(k.storeKey)

	var authorization1 types.AuthorizationI
	err := k.cdc.UnpackAny(authorization, &authorization1)
	if err != nil {
		return
	}

	bz := k.cdc.MustMarshalBinaryBare(&types.AuthorizationGrant{Authorization: authorization, Expiration: expiration})
	actor := k.getActorAuthorizationKey(grantee, granter, authorization1.MsgType())
	store.Set(actor, bz)
}

// Revoke method revokes any authorization for the provided message type granted to the grantee by the granter.
func (k Keeper) Revoke(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) error {
	store := ctx.KVStore(k.storeKey)
	actor := k.getActorAuthorizationKey(grantee, granter, msgType)
	_, found := k.getAuthorizationGrant(ctx, actor)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "authorization not found")
	}
	store.Delete(actor)

	return nil
}

// GetAuthorization Returns any `AuthorizationI` (or `nil`), with the expiration time,
// granted to the grantee by the granter for the provided msg type.
func (k Keeper) GetAuthorization(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) (cap types.AuthorizationI, expiration int64) {
	grant, found := k.getAuthorizationGrant(ctx, k.getActorAuthorizationKey(grantee, granter, msgType))
	if !found {
		return nil, 0
	}

	if grant.Expiration != 0 && grant.Expiration < (ctx.BlockHeader().Time.Unix()) {
		k.Revoke(ctx, grantee, granter, msgType)
		return nil, 0
	}

	var authorization types.AuthorizationI
	err := k.cdc.UnpackAny(grant.Authorization, &authorization)
	fmt.Println("err:", err)
	if err != nil {
		return nil, 0
	}

	return authorization, grant.Expiration
}

// UnmarshalAuthorization returns an Authorization interface from raw encoded authorization
// bytes of a Proto-based Authorization type. An error is returned upon decoding
// failure.
func (k Keeper) UnmarshalAuthorization(bz []byte) (types.AuthorizationI, error) {
	var authorization types.AuthorizationI

	if err := codec.UnmarshalAny(k.cdc, &authorization, bz); err != nil {
		return nil, err
	}

	return authorization, nil
}

// MarshalAuthorization marshals an Authorization interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalAuthorization(authorizationI types.AuthorizationI) ([]byte, error) {
	return codec.MarshalAny(k.cdc, authorizationI)
}
