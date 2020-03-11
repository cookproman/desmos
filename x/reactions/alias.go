package reactions

import (
	"github.com/desmos-labs/desmos/x/reactions/internal/keeper"
	"github.com/desmos-labs/desmos/x/reactions/internal/simulation"
	"github.com/desmos-labs/desmos/x/reactions/internal/types"
)

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
	NewPostReaction          = types.NewPostReaction
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	DecodeStore              = simulation.DecodeStore
	NewMsgAddPostReaction    = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction = types.NewMsgRemovePostReaction
)

type (
	Keeper        = keeper.Keeper
	PostReaction  = types.PostReaction
	PostReactions = types.PostReactions

	GenesisState          = types.GenesisState
	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction
)
