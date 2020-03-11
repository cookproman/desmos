package simulation

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding posts type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.HasPrefix(kvA.Key, types.PostReactionsStorePrefix):
		var reactionsA, reactionB types.PostReactions
		cdc.MustUnmarshalBinaryBare(kvA.Value, &reactionsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &reactionB)
		return fmt.Sprintf("ReactionsA: %s\nReactionsB: %s\n", reactionsA, reactionB)
	default:
		panic(fmt.Sprintf("invalid account key %X", kvA.Key))
	}
}
