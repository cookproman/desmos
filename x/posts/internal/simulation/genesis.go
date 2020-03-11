package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

var (
	RandomMimeTypes = []string{"audio/aac", "application/x-bzip2", "audio/ogg", "image/webp", "image/png"}
	RandomHosts     = []string{"https://example.com/", "https://ipfs.ink/"}
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simState *module.SimulationState) {
	posts := randomPosts(simState)
	postsGenesis := types.NewGenesisState(posts)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(postsGenesis)
}

// randomPosts returns randomly generated genesis accounts
func randomPosts(simState *module.SimulationState) (posts types.Posts) {
	postsNumber := simState.Rand.Intn(100)

	posts = make(types.Posts, postsNumber)
	for index := 0; index < postsNumber; index++ {
		id := index + 1

		postData := RandomPostData(simState.Rand, simState.Accounts)

		posts[index] = types.NewPost(
			types.PostID(id),
			types.PostID(simState.Rand.Intn(id)),
			postData.Message,
			postData.AllowsComments,
			postData.Subspace,
			postData.OptionalData,
			postData.CreationDate,
			postData.Creator.Address,
		).WithMedias(postData.Medias)

		if postData.PollData != nil {
			posts[index] = posts[index].WithPollData(*postData.PollData)
		}
	}

	return posts
}
