package upgrade

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"io/ioutil"
)

// UpgradeableStoreLoader can be configured by SetStoreLoader() to check for the
// existence of a given upgrade file - json encoded StoreUpgrades data.
//
// If not file is present, it will perform the default load (no upgrades to store).
//
// If the file is present, it will parse the file and execute those upgrades
// (rename or delete stores), while loading the data. It will also delete the
// upgrade file upon successful load, so that the upgrade is only applied once,
// and not re-applied on next restart
//
// This is useful for in place migrations when a store key is renamed between
// two versions of the software.
func UpgradeableStoreLoader(upgradeInfoPath string) baseapp.StoreLoader {
	return func(ms sdk.CommitMultiStore) error {

		data, err := ioutil.ReadFile(upgradeInfoPath)
		if err != nil {
			return fmt.Errorf("cannot read upgrade file %s: %v", upgradeInfoPath, err)
		}

		var upgrades storetypes.UpgradeInfo
		err = json.Unmarshal(data, &upgrades)
		if err != nil {
			return fmt.Errorf("cannot parse upgrade file: %v", err)
		}

		err = ms.LoadLatestVersionAndUpgrade(&upgrades.StoreUpgrades)
		if err != nil {
			return fmt.Errorf("load and upgrade database: %v", err)
		}

		// if we have a successful load, we set the values to default
		upgrades.Height = 0
		upgrades.StoreUpgrades = storetypes.StoreUpgrades{
			Renamed: []storetypes.StoreRename{{
				OldKey: "",
				NewKey: "",
			}},
			Deleted: []string{""},
		}
		return nil
	}
}
