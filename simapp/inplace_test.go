package simapp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

func TestInplaceMigrationFrom034(t *testing.T) {
	db := db.NewMemDB()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewSimApp(logger, db, nil, true, 0)

	info := app.Info(abci.RequestInfo{})
	require.Equal(t, 1000, info.GetLastBlockHeight())
	require.NotNil(t, info.GetLastBlockAppHash())

	err := app.inplaceMigration("foobar")
	require.NoError(t, err)
	res := app.Commit()
	require.NotEmpty(t, res.Data)
	require.NotEqual(t, info.GetLastBlockAppHash(), res.Data)
}
