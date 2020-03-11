package types_test

import (
	"errors"
	types2 "github.com/desmos-labs/desmos/x/reactions/internal/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestReaction_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	reaction := types2.NewPostReaction("reaction", user)
	require.Equal(t, `{"owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"reaction"}`, reaction.String())
}

func TestReaction_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name     string
		reaction types2.PostReaction
		error    error
	}{
		{
			name:     "Valid reaction returns no error",
			reaction: types2.NewPostReaction("reaction", user),
			error:    nil,
		},
		{
			name:     "Missing owner returns error",
			reaction: types2.NewPostReaction("reaction", nil),
			error:    errors.New("invalid reaction owner: "),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.error, test.reaction.Validate())
		})
	}
}

func TestReaction_Equals(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name          string
		first         types2.PostReaction
		second        types2.PostReaction
		shouldBeEqual bool
	}{
		{
			name:          "Returns false with different user",
			first:         types2.NewPostReaction("reaction", user),
			second:        types2.NewPostReaction("reaction", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns true with the same data",
			first:         types2.NewPostReaction("reaction", user),
			second:        types2.NewPostReaction("reaction", user),
			shouldBeEqual: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.shouldBeEqual, test.first.Equals(test.second))
		})
	}
}

func TestReactions_AppendIfMissing(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name      string
		reactions types2.PostReactions
		newLike   types2.PostReaction
		expLikes  types2.PostReactions
		expAppend bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: types2.PostReactions{},
			newLike:   types2.NewPostReaction("reaction", user),
			expLikes:  types2.PostReactions{types2.NewPostReaction("reaction", user)},
			expAppend: true,
		},
		{
			name:      "New reaction is appended properly to existing list",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			newLike:   types2.NewPostReaction("reaction", otherLiker),
			expAppend: true,
			expLikes: types2.PostReactions{
				types2.NewPostReaction("reaction", user),
				types2.NewPostReaction("reaction", otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.reactions.AppendIfMissing(test.newLike)
			require.Equal(t, test.expLikes, actual)
			require.Equal(t, test.expAppend, appended)
		})
	}
}

func TestReactions_ContainsOwnerLike(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name        string
		reactions   types2.PostReactions
		owner       sdk.AccAddress
		value       string
		expContains bool
	}{
		{
			name:        "Non-empty list returns true with valid address",
			reactions:   types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:       user,
			value:       "reaction",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   types2.PostReactions{},
			owner:       user,
			value:       "reaction",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found address",
			reactions:   types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:       otherLiker,
			value:       "reaction",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found value",
			reactions:   types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:       user,
			value:       "reaction-2",
			expContains: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expContains, test.reactions.ContainsPostReactionFrom(test.owner, test.value))
		})
	}
}

func TestReactions_IndexOfByUserAndValue(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name      string
		reactions types2.PostReactions
		owner     sdk.AccAddress
		value     string
		expIndex  int
	}{
		{
			name:      "Non-empty list returns proper index with valid value",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:     user,
			value:     "reaction",
			expIndex:  0,
		},
		{
			name:      "Empty list returns -1",
			reactions: types2.PostReactions{},
			owner:     user,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found address",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found value",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expIndex:  -1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expIndex, test.reactions.IndexOfByUserAndValue(test.owner, test.value))
		})
	}
}

func TestReactions_RemoveReaction(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name      string
		reactions types2.PostReactions
		owner     sdk.AccAddress
		value     string
		expResult types2.PostReactions
		expEdited bool
	}{
		{
			name:      "PostReaction is removed from non-empty list",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:     user,
			value:     "reaction",
			expResult: types2.PostReactions{},
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: types2.PostReactions{},
			owner:     user,
			value:     "reaction",
			expResult: types2.PostReactions{},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found address is not edited",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction",
			expResult: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found value is not edited",
			reactions: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expResult: types2.PostReactions{types2.NewPostReaction("reaction", user)},
			expEdited: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, edited := test.reactions.RemovePostReaction(test.owner, test.value)
			require.Equal(t, test.expEdited, edited)
			require.Equal(t, test.expResult, result)
		})
	}
}
