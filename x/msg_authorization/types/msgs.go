package types

import (
	"fmt"
	"time"

	proto "github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgGrantAuthorization(granter sdk.AccAddress, grantee sdk.AccAddress, authorization AuthorizationI, expiration time.Time) (*MsgGrantAuthorization, error) {
	m := &MsgGrantAuthorization{
		Granter:    granter,
		Grantee:    grantee,
		Expiration: expiration,
	}

	err := m.SetAuthorization(authorization)
	return m, err
}

func (msg MsgGrantAuthorization) Route() string { return RouterKey }
func (msg MsgGrantAuthorization) Type() string  { return "grant_authorization" }

func (msg MsgGrantAuthorization) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Granter}
}

func (msg MsgGrantAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgGrantAuthorization) ValidateBasic() error {
	if msg.Granter.Empty() {
		return sdkerrors.Wrap(ErrInvalidGranter, "missing granter address")
	}
	if msg.Grantee.Empty() {
		return sdkerrors.Wrap(ErrInvalidGranter, "missing grantee address")
	}
	if msg.Expiration.Unix() < time.Now().Unix() {
		return sdkerrors.Wrap(ErrInvalidGranter, "Time can't be in the past")
	}

	return nil
}

func (m *MsgGrantAuthorization) SetAuthorization(authorization AuthorizationI) error {
	msg, ok := authorization.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.Authorization = any
	return nil
}

func (m *MsgGrantAuthorization) GetAuthorization() AuthorizationI {
	// var autorization AuthorizationI
	autorization, ok := m.Authorization.GetCachedValue().(AuthorizationI)
	// err := ModuleCdc.UnpackAny(m.Authorization, &autorization)
	if !ok {
		return nil
	}
	return autorization
}

func NewMsgRevokeAuthorization(granter sdk.AccAddress, grantee sdk.AccAddress, authorizationMsgType string) *MsgRevokeAuthorization {
	return &MsgRevokeAuthorization{
		Granter:              granter,
		Grantee:              grantee,
		AuthorizationMsgType: authorizationMsgType,
	}
}

func (msg MsgRevokeAuthorization) Route() string { return RouterKey }
func (msg MsgRevokeAuthorization) Type() string  { return "revoke_authorization" }

func (msg MsgRevokeAuthorization) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Granter}
}

func (msg MsgRevokeAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgRevokeAuthorization) ValidateBasic() error {
	if msg.Granter.Empty() {
		return sdkerrors.Wrap(ErrInvalidGranter, "missing granter address")
	}
	if msg.Grantee.Empty() {
		return sdkerrors.Wrap(ErrInvalidGranter, "missing grantee address")
	}
	return nil
}

// MsgExecAuthorized attempts to execute the provided messages using
// authorizations granted to the grantee. Each message should have only
// one signer corresponding to the granter of the authorization.
// type MsgExecAuthorized struct {
// 	Grantee sdk.AccAddress `json:"grantee"`
// 	Msgs    []sdk.Msg      `json:"msgs"`
// }

func NewMsgExecAuthorized(grantee sdk.AccAddress, msgs []sdk.Msg) (MsgExecAuthorized, error) {
	// TODO change to pointer address
	m := MsgExecAuthorized{
		Grantee: grantee,
	}

	err := m.SetMsgs(msgs)

	return m, err
}

func (msg MsgExecAuthorized) Route() string { return RouterKey }
func (msg MsgExecAuthorized) Type() string  { return "exec_delegated" }

func (msg MsgExecAuthorized) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Grantee}
}

func (msg MsgExecAuthorized) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgExecAuthorized) ValidateBasic() error {
	if msg.Grantee.Empty() {
		return sdkerrors.Wrap(ErrInvalidGranter, "missing grantee address")
	}
	return nil
}

func (m *MsgExecAuthorized) SetMsgs(msgs []sdk.Msg) error {
	for _, msg := range msgs {
		msg1, ok := msg.(proto.Message)
		if !ok {
			return fmt.Errorf("can't proto marshal %T", msg1)
		}
		any, err := types.NewAnyWithValue(msg1)
		if err != nil {
			return err
		}
		m.Msgs = append(m.Msgs, any)
	}

	return nil
}

// GetMsgs return unpacked interfaces from any
func (m *MsgExecAuthorized) GetMsgs() ([]sdk.Msg, error) {
	var msgs []sdk.Msg
	for _, msgItem := range m.Msgs {
		var msgInfo sdk.Msg
		err := ModuleCdc.UnpackAny(msgItem, &msgInfo)
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, msgInfo)
	}

	return msgs, nil
}
