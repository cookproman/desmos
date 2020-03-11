package keeper

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
)

// -------------
// --- PostReactions
// -------------

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// -------------
// --- PostReactions
// -------------

// SaveReaction allows to save the given reaction inside the store.
// It assumes that the given reaction is valid.
// If another reaction from the same user for the same post and with the same value exists, returns an expError.
// nolint: interfacer
func (k Keeper) SaveReaction(ctx sdk.Context, postID posts.PostID, reaction types.PostReaction) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existent reactions
	var reactions types.PostReactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check for double reactions
	if reactions.ContainsPostReactionFrom(reaction.Owner, reaction.Value) {
		return fmt.Errorf("%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.Value, postID)
	}

	// Save the new reaction
	reactions = append(reactions, reaction)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&reactions))

	return nil
}

// RemovePostReaction removes the reaction from the given user from the post having the
// given postID. If no reaction with the same value was previously added from the given user, an expError
// is returned.
// nolint: interfacer
func (k Keeper) RemoveReaction(ctx sdk.Context, postID posts.PostID, user sdk.AccAddress, value string) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existing reactions
	var reactions types.PostReactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check if the user exists
	if !reactions.ContainsPostReactionFrom(user, value) {
		return fmt.Errorf("cannot remove the reaction with value %s from user %s as it does not exist",
			value, user)
	}

	// Remove and save the reactions list
	if newLikes, edited := reactions.RemovePostReaction(user, value); edited {
		if len(newLikes) == 0 {
			store.Delete(key)
		} else {
			store.Set(key, k.Cdc.MustMarshalBinaryBare(&newLikes))
		}
	}

	return nil
}

// GetPostReactions returns the list of reactions that has been associated to the post having the given id
// nolint: interfacer
func (k Keeper) GetPostReactions(ctx sdk.Context, postID posts.PostID) types.PostReactions {
	store := ctx.KVStore(k.StoreKey)

	var reactions types.PostReactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(postID)), &reactions)
	return reactions
}

// GetReactions allows to returns the list of reactions that have been stored inside the given context
func (k Keeper) GetReactions(ctx sdk.Context) map[posts.PostID]types.PostReactions {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsStorePrefix)

	reactionsData := map[posts.PostID]types.PostReactions{}
	for ; iterator.Valid(); iterator.Next() {
		var postLikes types.PostReactions
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postLikes)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PostReactionsStorePrefix)
		postID, err := posts.ParsePostID(string(idBytes))
		if err != nil {
			// This should never verify
			panic(err)
		}

		reactionsData[postID] = postLikes
	}

	return reactionsData
}
