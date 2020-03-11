package types_test

import (
	types "github.com/desmos-labs/desmos/x/reactions/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateGenesis(t *testing.T) {
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
			name: "Genesis with invalid reaction errors",
			genesis: types.GenesisState{
				Reactions: map[string]types.PostReactions{
					"1": {types.PostReaction{Owner: nil}},
				},
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
