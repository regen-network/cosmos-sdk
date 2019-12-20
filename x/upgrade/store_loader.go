package upgrade

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"io/ioutil"
)

func UpgradeableStoreLoader(upgradeInfoPath string) baseapp.StoreLoader {
	return func(ms sdk.CommitMultiStore) error {

		data, err := ioutil.ReadFile(upgradeInfoPath)
		if err != nil {
			return fmt.Errorf("cannot read upgrade file %s: %v", upgradeInfoPath, err)
		}

		var upgrades storetypes.UpgradeFile
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
