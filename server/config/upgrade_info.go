package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	store "github.com/cosmos/cosmos-sdk/store/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

func GenerateOrLoadUpgradeInfoFile() error {
	home := viper.GetString(cli.HomeFlag)
	upgradeFilePath := filepath.Join(home, "upgrade-info.json")
	if _, err := os.Stat(upgradeFilePath); err == nil {
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
		err = ioutil.WriteFile(upgradeFilePath, info, 0644)
	}

	return nil
}
