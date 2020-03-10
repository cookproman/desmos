package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reactions/internal/keeper"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_handleMsgAddPostReaction(t *testing.T) {

	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	require.NoError(t, err)

	tests := []struct {
		name         string
		existingPost *posts.Post
		msg          types.MsgAddPostReaction
		error        error
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgAddPostReaction(posts.PostID(0), "like", user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 0 not found"),
		},
		{
			name:         "Valid message works properly",
			existingPost: &testPost,
			msg:          types.NewMsgAddPostReaction(testPost.PostID, "like", user),
			error:        nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(posts.PostStoreKey(test.existingPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				require.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypeReactionAdded,
					sdk.NewAttribute(posts.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyReactionValue, test.msg.Value),
				))

				var storedPost posts.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(testPost.PostID)), &storedPost)
				require.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.PostReactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				require.Contains(t, storedReactions, types.NewReaction(test.msg.Value, test.msg.User))
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}

func Test_handleMsgRemovePostReaction(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	require.NoError(t, err)

	reaction := types.NewReaction("like", user)
	tests := []struct {
		name             string
		existingPost     *posts.Post
		existingReaction *types.PostReaction
		msg              types.MsgRemovePostReaction
		error            error
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgRemovePostReaction(posts.PostID(0), user, "like"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 0 not found"),
		},
		{
			name:         "PostReaction not found",
			existingPost: &testPost,
			msg:          types.NewMsgRemovePostReaction(testPost.PostID, user, "like"),
			error:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot remove the reaction with value like from user %s as it does not exist", user)),
		},
		{
			name:             "Valid message works properly",
			existingPost:     &testPost,
			existingReaction: &reaction,
			msg:              types.NewMsgRemovePostReaction(testPost.PostID, user, reaction.Value),
			error:            nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			if test.existingReaction != nil {
				store.Set(
					types.PostReactionsStoreKey(test.existingPost.PostID),
					k.Cdc.MustMarshalBinaryBare(&types.PostReactions{*test.existingReaction}),
				)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				require.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(posts.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyReactionValue, test.msg.Reaction),
				))

				var storedPost posts.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(testPost.PostID)), &storedPost)
				require.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.PostReactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				require.NotContains(t, storedReactions, test.existingReaction)
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}
