package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

func GenerateOrLoadUpgradeInfoFile() error {
	home := viper.GetString(cli.HomeFlag)
	upgradeFilePath := home + "upgrade-info.json"
	if _, err := os.Stat(upgradeFilePath); err == nil {
		return nil
	} else if os.IsNotExist(err) {
		var upgradeFile store.UpgradeFile
		upgradeFile.Height = 0

		info, err := json.Marshal(upgradeFile)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(upgradeFilePath, info, 0644)
	}

	return nil
}
