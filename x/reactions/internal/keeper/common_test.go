package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
	"time"

	"github.com/desmos-labs/desmos/x/reactions/internal/keeper"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
)

//TODO add post keeper generation
func SetupTestInput() (sdk.Context, keeper.Keeper) {

	// define store store keys
	magpieKey := sdk.NewKVStoreKey("magpie")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(magpieKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	// create a Cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	return ctx, keeper.NewKeeper(cdc, magpieKey)
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
var timeZone, _ = time.LoadLocation("UTC")
var testPostCreationDate = time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
var testPostEndPollDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)
var testPostEndPollDateExpired = time.Date(2019, 1, 1, 1, 15, 00, 000, timeZone)

var testPost = posts.NewPost(
	posts.PostID(3257),
	posts.PostID(0),
	"Post message",
	false,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	map[string]string{},
	testPostCreationDate,
	testPostOwner,
)
