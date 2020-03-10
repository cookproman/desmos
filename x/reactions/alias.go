package reactions

import "github.com/desmos-labs/desmos/x/reactions/internal/types"

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper                = keeper.NewKeeper
	NewHandler               = keeper.NewHandler
	NewQuerier               = keeper.NewQuerier
	ModuleCdc                = types.ModuleCdc
	RegisterCodec            = types.RegisterCodec
	NewMsgAddPostReaction    = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction = types.NewMsgRemovePostReaction
	NewReaction              = types.NewReaction
)

type (
	PostReaction  = types.PostReaction
	PostReactions = types.PostReactions

	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction
)
