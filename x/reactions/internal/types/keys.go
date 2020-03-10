package types

import (
	"github.com/desmos-labs/desmos/x/posts"
	"regexp"
)

const (
	ModuleName = "reactions"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionAddPostReaction    = "add_post_reaction"
	ActionRemovePostReaction = "remove_post_reaction"

	QuerierRoute = ModuleName
)

var (
	ReactionRegEx = regexp.MustCompile("")

	PostReactionsStorePrefix = []byte("reactions")
)

// PostCommentsStoreKey turns an id to a key used to store a post's reactions into the posts store
// nolint: interfacer
func PostReactionsStoreKey(id posts.PostID) []byte {
	return append(PostReactionsStorePrefix, []byte(id.String())...)
}
