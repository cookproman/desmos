package types_test

import (
	types2 "github.com/desmos-labs/desmos/x/reactions/internal/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name        string
		genesis     types.GenesisState
		shouldError bool
	}{
		{
			name:        "DefaultGenesis does not error",
			genesis:     types.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Genesis with invalid post errors",
			genesis: types.GenesisState{
				Posts:     types.Posts{types.Post{PostID: types.PostID(0)}},
				Reactions: map[string]types2.Reactions{},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid reaction errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				Reactions: map[string]types2.Reactions{
					"1": {types2.PostReaction{Owner: nil}},
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid poll answers errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				PollAnswers: map[string]types.UserAnswers{
					"1": {
						types.NewUserAnswer([]types.AnswerID{}, user),
					},
				},
				Reactions: map[string]types2.Reactions{},
			},
			shouldError: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.shouldError {
				require.Error(t, types.ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, types.ValidateGenesis(test.genesis))
			}
		})
	}
}
