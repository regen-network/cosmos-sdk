package simapp

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

func TestInplaceMigrationFrom034(t *testing.T) {
	dbName := "application_034"
	var blockHeight int64 = 525653
	chainID := "regen-test-1001"
	tmpDir, cleanup := CopyTestdata(t, dbName+".db")
	defer cleanup()

	ldb, err := db.NewGoLevelDB(dbName, tmpDir)
	require.Nil(t, err)
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewSimApp(logger, ldb, nil, true, 0)

	info := app.Info(abci.RequestInfo{})
	require.Equal(t, blockHeight, info.GetLastBlockHeight())
	require.NotNil(t, info.GetLastBlockAppHash())

	err = app.inplaceMigration(chainID)
	require.NoError(t, err)
	res := app.Commit()
	require.NotEmpty(t, res.Data)
	require.NotEqual(t, info.GetLastBlockAppHash(), res.Data)

	checkGovStore(t, app, chainID)
	checkDistrStore(t, app, chainID)
}

// checkGovStore will try to load objects that have changed, should fail on unmigrated data
func checkGovStore(t *testing.T, app *SimApp, chainID string) {
	now := time.Now().UTC()
	header := abci.Header{ChainID: chainID, Time: now}
	ctx := app.GetDeliverState(header)

	app.govKeeper.IterateProposals(ctx, func(proposal govtypes.Proposal) (stop bool) {
		t.Logf("-> P: %s", proposal)
		return false
	})
	t.Logf("Read all proposals")
}

// checkDistrStore will try to load objects that have changed, should fail on unmigrated data
func checkDistrStore(t *testing.T, app *SimApp, chainID string) {
	now := time.Now().UTC()
	header := abci.Header{ChainID: chainID, Time: now}
	ctx := app.GetDeliverState(header)

	rewards := app.distrKeeper.GetTotalRewards(ctx)
	t.Logf("Rewards: %#v", rewards)
}


func CopyTestdata(t *testing.T, subDir string) (string, func()) {
	sourceDir := "testdata"
	destDir, err := ioutil.TempDir("", "test-migration")
	require.Nil(t, err)
	cleanup := func() { os.RemoveAll(destDir) }

	err = copyFiles(t, sourceDir, destDir, subDir)
	if err != nil {
		cleanup()
		t.Fatalf("Cannot copy config files: %+v", err)
	}
	return destDir, cleanup
}

func copyFiles(t *testing.T, sourceDir, rootDir, subDir string) error {
	// make the output dir
	outDir := filepath.Join(rootDir, subDir)
	err := os.Mkdir(outDir, 0755)
	if err != nil {
		return err
	}

	// copy everything over from testdata
	inDir := filepath.Join(sourceDir, subDir)
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		input := filepath.Join(inDir, f.Name())
		output := filepath.Join(outDir, f.Name())
		// t.Logf("Copying %s to %s", input, output)
		err = fileCopy(input, output, f.Mode())
		if err != nil {
			return err
		}
	}

	return nil
}

func fileCopy(input, output string, mode os.FileMode) error {
	from, err := os.Open(input)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, mode)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}
