package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reactions/internal/keeper"
	"math/rand"
)

// Simulation operation weights constants
const (
	OpWeightMsgAddReaction    = "op_weight_msg_add_reaction"
	OpWeightMsgRemoveReaction = "op_weight_msg_remove_reaction"

	DefaultGasValue = 5000000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper, pk posts.Keeper) sim.WeightedOperations {

	var weightMsgAddReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddReaction, &weightMsgAddReaction, nil,
		func(_ *rand.Rand) {
			weightMsgAddReaction = params.DefaultWeightMsgAddReaction
		},
	)

	var weightMsgRemoveReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveReaction, &weightMsgRemoveReaction, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveReaction = params.DefaultWeightMsgRemoveReaction
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgAddReaction,
			SimulateMsgAddReaction(k, ak, pk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveReaction,
			SimulateMsgRemoveReaction(k, ak),
		),
	}
}
