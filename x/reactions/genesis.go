package reactions

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	abci "github.com/tendermint/tendermint/abci/types"
)

func convertReactionsMap(reactions map[posts.PostID]PostReactions) map[string]PostReactions {
	reactionsMap := make(map[string]PostReactions, len(reactions))
	for key, value := range reactions {
		reactionsMap[key.String()] = value
	}
	return reactionsMap
}

func convertGenesisReactions(reactions map[string]PostReactions) map[posts.PostID]PostReactions {
	reactionsMap := make(map[posts.PostID]PostReactions, len(reactions))
	for key, value := range reactions {
		postID, err := posts.ParsePostID(key)
		if err != nil {
			panic(err)
		}
		reactionsMap[postID] = value
	}
	return reactionsMap
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Reactions: convertReactionsMap(k.GetReactions(ctx)),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	reactionsMap := convertGenesisReactions(data.Reactions)
	for postID, reactions := range reactionsMap {
		for _, reaction := range reactions {
			if err := keeper.SaveReaction(ctx, postID, reaction); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
