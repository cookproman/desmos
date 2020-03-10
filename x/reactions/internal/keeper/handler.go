package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
)

func NewHandler(postsKeeper posts.Keeper, keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgAddPostReaction:
			return handleMsgAddPostReaction(ctx, postsKeeper, keeper, msg)
		case types.MsgRemovePostReaction:
			return handleMsgRemovePostReaction(ctx, postsKeeper, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// -------------
// --- PostReactions
// -------------

// handleMsgAddPostReaction handles the adding of a reaction to a post
func handleMsgAddPostReaction(ctx sdk.Context, postKeeper posts.Keeper, keeper Keeper, msg types.MsgAddPostReaction) (*sdk.Result, error) {

	// Get the post
	post, found := postKeeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s not found", msg.PostID))
	}

	// Create and store the reaction
	reaction := types.NewReaction(msg.Value, msg.User)
	if err := keeper.SaveReaction(ctx, post.PostID, reaction); err != nil {
		return nil, err
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypeReactionAdded,
		sdk.NewAttribute(posts.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Value),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("reaction added properly"),
		Events: sdk.Events{event},
	}
	return &result, nil
}

// handleMsgRemovePostReaction handles the removal of a reaction from a post
func handleMsgRemovePostReaction(ctx sdk.Context, postKeeper posts.Keeper, keeper Keeper, msg types.MsgRemovePostReaction) (*sdk.Result, error) {

	// Get the post
	post, found := postKeeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s not found", msg.PostID))
	}

	// Remove the reaction
	if err := keeper.RemoveReaction(ctx, post.PostID, msg.User, msg.Reaction); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypePostReactionRemoved,
		sdk.NewAttribute(posts.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Reaction),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("reaction removed properly"),
		Events: sdk.Events{event},
	}
	return &result, nil
}
