package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/posts"
	"math/rand"
)

var (
	reactsValues = []string{"💙", "⬇️", "👎", "like"}
)

// ReactionData contains all the data needed for a reaction to be properly added or removed from a post
type ReactionData struct {
	Value  string
	User   sim.Account
	PostID posts.PostID
}

// RandomReactionData returns a randomly generated reaction data object
func RandomReactionData(r *rand.Rand, accs []sim.Account, postz []posts.Post) ReactionData {
	return ReactionData{
		Value:  RandomReactionValue(r),
		User:   accs[r.Intn(len(accs))],
		PostID: posts.RandomPostID(r, postz),
	}
}

// RandomReactionValue returns a random reaction value
func RandomReactionValue(r *rand.Rand) string {
	return reactsValues[r.Intn(len(reactsValues))]
}

// GetAccount gets the account having the given address from the accs list
func GetAccount(address sdk.Address, accs []sim.Account) *sim.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}
