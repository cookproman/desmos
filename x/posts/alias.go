package posts

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper           = keeper.NewKeeper
	NewHandler          = keeper.NewHandler
	NewQuerier          = keeper.NewQuerier
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	ParsePostID         = types.ParsePostID
	NewPost             = types.NewPost
	NewPostMedia        = types.NewPostMedia
	NewPostMedias       = types.NewPostMedias
	NewPollData         = types.NewPollData
	NewPollAnswer       = types.NewPollAnswer
	NewPollAnswers      = types.NewPollAnswers
	NewUserAnswer       = types.NewUserAnswer
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	WeightedOperations  = simulation.WeightedOperations
	NewMsgCreatePost    = types.NewMsgCreatePost
	NewMsgEditPost      = types.NewMsgEditPost

	DecodeStore        = simulation.DecodeStore
	RandomizedGenState = simulation.RandomizedGenState
	RandomPost         = simulation.RandomPost
	RandomPostData     = simulation.RandomPostData
	RandomPostID       = simulation.RandomPostID
	RandomMessage      = simulation.RandomMessage
	RandomSubspace     = simulation.RandomSubspace
	RandomMedias       = simulation.RandomMedias
	RandomPollData     = simulation.RandomPollData

	AttributeKeyPostID = types.AttributeKeyPostID
	PostStoreKey       = types.PostStoreKey
)

type (
	Keeper      = keeper.Keeper
	PostID      = types.PostID
	PostIDs     = types.PostIDs
	Post        = types.Post
	Posts       = types.Posts
	PostMedia   = types.PostMedia
	PostMedias  = types.PostMedias
	PollData    = types.PollData
	AnswerID    = types.AnswerID
	UserAnswer  = types.UserAnswer
	UserAnswers = types.UserAnswers

	PostQueryResponse = types.PostQueryResponse
	GenesisState      = types.GenesisState
	MsgCreatePost     = types.MsgCreatePost
	MsgEditPost       = types.MsgEditPost
	MsgAnswerPoll     = types.MsgAnswerPoll
	QueryPostsParams  = types.QueryPostsParams

	PostData = simulation.PostData
)
