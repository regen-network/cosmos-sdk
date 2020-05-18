package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdkstd "github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

const (
	holder     = "holder"
	multiPerm  = "multiple permissions account"
	randomPerm = "random permission"
)

func makeTestCodec() codec.Marshaler {
	var (
		amino = codec.New()

		ModuleCdc = codec.NewHybridCodec(amino, codectypes.NewInterfaceRegistry())
	)
	auth.RegisterCodec(amino)
	types.RegisterCodec(amino)
	staking.RegisterCodec(amino)
	sdk.RegisterCodec(amino)
	codec.RegisterCrypto(amino)

	return ModuleCdc
}
func SetupTestInput() (sdk.Context, auth.AccountKeeper, params.Keeper, bank.BaseKeeper, Keeper, baseapp.Router) {
	db := dbm.NewMemDB()

	amino := codec.New()
	auth.RegisterCodec(amino)
	bank.RegisterCodec(amino)
	sdk.RegisterCodec(amino)
	codec.RegisterCrypto(amino)

	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	keyBank := sdk.NewKVStoreKey(bank.StoreKey)
	keyAuthorization := sdk.NewKVStoreKey(types.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)

	maccPerms := simapp.GetMaccPerms()
	maccPerms[holder] = nil
	maccPerms[authtypes.Burner] = []string{authtypes.Burner}
	maccPerms[auth.Minter] = []string{authtypes.Minter}
	maccPerms[multiPerm] = []string{authtypes.Burner, authtypes.Minter, authtypes.Staking}
	maccPerms[randomPerm] = []string{"random"}

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAuthorization, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)

	ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{Time: time.Unix(0, 0)}, false, log.NewNopLogger())
	ModuleCdc := makeTestCodec()

	blacklistedAddrs := make(map[string]bool)

	paramsKeeper := params.NewKeeper(ModuleCdc, keyParams, tkeyParams)
	authKeeper := auth.NewAccountKeeper(sdkstd.NewAppCodec(amino, codectypes.NewInterfaceRegistry()),
		keyAcc, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount, maccPerms)
	bankKeeper := bank.NewBaseKeeper(ModuleCdc, keyBank, authKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)
	bankKeeper.SetSendEnabled(ctx, true)

	router := *baseapp.NewRouter()
	router.AddRoute("bank", bank.NewHandler(bankKeeper))

	authorizationKeeper := NewKeeper(keyAuthorization, ModuleCdc, router)
	authKeeper.SetParams(ctx, auth.DefaultParams())

	return ctx, authKeeper, paramsKeeper, bankKeeper, authorizationKeeper, router
}

var (
	granteePub    = ed25519.GenPrivKey().PubKey()
	granterPub    = ed25519.GenPrivKey().PubKey()
	recipientPub  = ed25519.GenPrivKey().PubKey()
	granteeAddr   = sdk.AccAddress(granteePub.Address())
	granterAddr   = sdk.AccAddress(granterPub.Address())
	recipientAddr = sdk.AccAddress(recipientPub.Address())
)
