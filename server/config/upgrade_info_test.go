package config

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGenerateOrLoadUpgradeInfoFile(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "example")
	require.NoError(t, err)
	defaultUpgradeFileName := "upgrade-file.json"
	defaultUpgradeFilePath := filepath.Join(tmpDir, defaultUpgradeFileName)
	err = GenerateOrLoadUpgradeInfoFile(defaultUpgradeFilePath)
	require.Nil(t, err)
}
