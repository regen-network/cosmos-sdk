package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/tests/cli/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/types"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGaiadCollectGentxs(t *testing.T) {
	t.Parallel()
	var customMaxBytes, customMaxGas int64 = 99999999, 1234567
	f := helpers.NewFixtures(t)

	// Initialise temporary directories
	gentxDir, err := ioutil.TempDir("", "")
	gentxDoc := filepath.Join(gentxDir, "gentx.json")
	require.NoError(t, err)

	// Reset testing path
	f.UnsafeResetAll()

	// Initialize keys
	f.KeysAdd(helpers.KeyFoo)

	// Configure json output
	f.CLIConfig("output", "json")

	// Run init
	f.SDInit(helpers.KeyFoo)

	// Customise genesis.json

	genFile := f.GenesisFile()
	genDoc, err := tmtypes.GenesisDocFromFile(genFile)
	require.NoError(t, err)
	genDoc.ConsensusParams.Block.MaxBytes = customMaxBytes
	genDoc.ConsensusParams.Block.MaxGas = customMaxGas
	genDoc.SaveAs(genFile)

	// Add account to genesis.json
	f.AddGenesisAccount(f.KeyAddress(helpers.KeyFoo), helpers.StartCoins)

	// Write gentx file
	f.GenTx(helpers.KeyFoo, fmt.Sprintf("--output-document=%s", gentxDoc))

	// Collect gentxs from a custom directory
	f.CollectGenTxs(fmt.Sprintf("--gentx-dir=%s", gentxDir))

	genDoc, err = tmtypes.GenesisDocFromFile(genFile)
	require.NoError(t, err)
	require.Equal(t, genDoc.ConsensusParams.Block.MaxBytes, customMaxBytes)
	require.Equal(t, genDoc.ConsensusParams.Block.MaxGas, customMaxGas)

	f.Cleanup(gentxDir)
}

func TestGaiadAddGenesisAccount(t *testing.T) {
	t.Parallel()
	f := helpers.NewFixtures(t)

	// Reset testing path
	f.UnsafeResetAll()

	// Initialize keys
	f.KeysDelete(helpers.KeyFoo)
	f.KeysDelete(helpers.KeyBar)
	f.KeysDelete(helpers.KeyBaz)
	f.KeysAdd(helpers.KeyFoo)
	f.KeysAdd(helpers.KeyBar)
	f.KeysAdd(helpers.KeyBaz)

	// Configure json output
	f.CLIConfig("output", "json")

	// Run init
	f.SDInit(helpers.KeyFoo)

	// Add account to genesis.json
	bazCoins := sdk.Coins{
		sdk.NewInt64Coin("acoin", 1000000),
		sdk.NewInt64Coin("bcoin", 1000000),
	}

	f.AddGenesisAccount(f.KeyAddress(helpers.KeyFoo), helpers.StartCoins)
	f.AddGenesisAccount(f.KeyAddress(helpers.KeyBar), bazCoins)

	genesisState := f.GenesisState()

	appCodec := std.NewAppCodec(f.Cdc)

	accounts := auth.GetGenesisStateFromAppState(appCodec, genesisState).Accounts
	balances := bank.GetGenesisStateFromAppState(f.Cdc, genesisState).Balances
	balancesSet := make(map[string]sdk.Coins)

	for _, b := range balances {
		balancesSet[b.GetAddress().String()] = b.Coins
	}

	require.Equal(t, accounts[0].GetAddress(), f.KeyAddress(helpers.KeyFoo))
	require.Equal(t, accounts[1].GetAddress(), f.KeyAddress(helpers.KeyBar))
	require.True(t, balancesSet[accounts[0].GetAddress().String()].IsEqual(helpers.StartCoins))
	require.True(t, balancesSet[accounts[1].GetAddress().String()].IsEqual(bazCoins))

	// Cleanup testing directories
	f.Cleanup()
}

func TestValidateGenesis(t *testing.T) {
	t.Parallel()
	f := helpers.InitFixtures(t)

	// start gaiad server
	proc := f.SDStart()
	defer proc.Stop(false)

	f.ValidateGenesis()

	// Cleanup testing directories
	f.Cleanup()
}
