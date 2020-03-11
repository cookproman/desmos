package simulation

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	types "github.com/desmos-labs/desmos/x/reactions/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"
	"testing"
	"time"
)

var (
	privKey         = ed25519.GenPrivKey().PubKey()
	postCreatorAddr = sdk.AccAddress(privKey.Address())

	postID = posts.PostID(3257)

	timeZone, _ = time.LoadLocation("UTC")
)

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	return
}

func TestDecodeStore(t *testing.T) {
	cdc := makeTestCodec()

	reactions := types.PostReactions{
		types.NewPostReaction("like", postCreatorAddr),
		types.NewPostReaction("💙", postCreatorAddr),
	}

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.PostReactionsStoreKey(postID), Value: cdc.MustMarshalBinaryBare(&reactions)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"PostReactions", fmt.Sprintf("ReactionsA: %s\nReactionsB: %s\n", reactions, reactions)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
