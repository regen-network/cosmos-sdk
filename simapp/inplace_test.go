package simapp

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

func TestInplaceMigrationFrom034(t *testing.T) {
	dbName := "application_034"
	var blockHeight int64 = 525653
	tmpDir, cleanup := CopyTestdata(t, dbName+".db")
	defer cleanup()

	ldb, err := db.NewGoLevelDB(dbName, tmpDir)
	require.Nil(t, err)
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewSimApp(logger, ldb, nil, true, 0)

	info := app.Info(abci.RequestInfo{})
	require.Equal(t, blockHeight, info.GetLastBlockHeight())
	require.NotNil(t, info.GetLastBlockAppHash())

	err = app.inplaceMigration("foobar")
	require.NoError(t, err)
	res := app.Commit()
	require.NotEmpty(t, res.Data)
	require.NotEqual(t, info.GetLastBlockAppHash(), res.Data)
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
