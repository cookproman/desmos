package posts

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func convertPostPollAnswersMap(answers map[PostID]UserAnswers) map[string]UserAnswers {
	answersMap := make(map[string]UserAnswers, len(answers))
	for key, value := range answers {
		answersMap[key.String()] = value
	}
	return answersMap
}

func convertGenesisPostPollAnswers(pollAnswers map[string]UserAnswers) map[PostID]UserAnswers {
	answersMap := make(map[PostID]UserAnswers, len(pollAnswers))
	for key, value := range pollAnswers {
		postID, err := ParsePostID(key)
		if err != nil {
			panic(err)
		}
		answersMap[postID] = value
	}
	return answersMap
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Posts:       k.GetPosts(ctx),
		PollAnswers: convertPostPollAnswersMap(k.GetPollAnswersMap(ctx)),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	// Sort the posts so that they are inserted based on their IDs
	sort.Sort(data.Posts)
	for _, post := range data.Posts {
		keeper.SavePost(ctx, post)
	}

	pollAnswersMap := convertGenesisPostPollAnswers(data.PollAnswers)
	for postID, usersAnswersDetails := range pollAnswersMap {
		for _, userAnswersDetails := range usersAnswersDetails {
			keeper.SavePollAnswers(ctx, postID, userAnswersDetails)
		}
	}
	return []abci.ValidatorUpdate{}
}
