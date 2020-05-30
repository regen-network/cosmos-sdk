package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
)

var msg_auth = "msgauth"

type genesisState struct {
	Authorizations []types.AuthorizationI
}

func genAuthorizations(_ *rand.Rand, _ []simtypes.Account) []types.AuthorizationI {
	return []types.AuthorizationI{}
}

// RandomizedGenState generates a random GenesisState for msg_auth
func RandomizedGenState(simState *module.SimulationState) {
	var authorizations []types.AuthorizationI

	simState.AppParams.GetOrGenerate(
		simState.Cdc, msg_auth, &authorizations, simState.Rand,
		func(r *rand.Rand) { authorizations = genAuthorizations(r, simState.Accounts) },
	)

	// TODO: Replace GenesisState of genesis.go file
	authorizationsGenesis := genesisState{
		Authorizations: authorizations,
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, codec.MustMarshalJSONIndent(simState.Cdc, authorizationsGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(authorizationsGenesis)
}
