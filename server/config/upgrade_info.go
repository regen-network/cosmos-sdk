package config

import (
	"encoding/json"
	store "github.com/cosmos/cosmos-sdk/store/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"io/ioutil"
	"os"
)

func GenerateOrLoadUpgradeInfoFile(defaultUpgradeFilePath string) error {

	if _, err := os.Stat(defaultUpgradeFilePath); err == nil {
		return nil
	} else if os.IsNotExist(err) {
		var upgradeFile store.UpgradeFile
		upgradeFile.Height = 0
		upgradeFile.StoreUpgrades = storetypes.StoreUpgrades{
			Renamed: []storetypes.StoreRename{{
				OldKey: "",
				NewKey: "",
			}},
			Deleted: []string{""},
		}
		info, err := json.Marshal(upgradeFile)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(defaultUpgradeFilePath, info, 0644)
	}

	return nil
}
