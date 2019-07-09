package v036

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dist036 "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v0_36"
	gov036 "github.com/cosmos/cosmos-sdk/x/gov/legacy/v036"
)

// This can be registered as a handler for the x/upgrade module
// (once that is merged in a separate PR)
// Should be done in the app setup, eg.
//
//   app.upgradeKeeper.SetUpgradeHandler("v0_36", func(ctx sdk.Context, plan upgrade.Plan) {
//     InplaceMigrationFromv034(ctx, app.govKey, app.distrKey)
//   })
func InplaceMigrationFromv034(ctx sdk.Context, govKey sdk.StoreKey, distKey sdk.StoreKey) error {
	err := dist036.InplaceMigrationFromv034(ctx, distKey)
	if err != nil {
		return err
	}
	err = gov036.InplaceMigrationFromv034(ctx, govKey)
	return err
}
