package posts

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

const (
	EventTypePostCreated            = types.EventTypePostCreated
	EventTypePostEdited             = types.EventTypePostEdited
	EventTypeReactionAdded          = types.EventTypeReactionAdded
	EventTypePostReactionRemoved    = types.EventTypePostReactionRemoved
	EventTypeAnsweredPoll           = types.EventTypeAnsweredPoll
	EventTypeClosePoll              = types.EventTypeClosePoll
	AttributeKeyPostID              = types.AttributeKeyPostID
	AttributeKeyPostParentID        = types.AttributeKeyPostParentID
	AttributeKeyPostOwner           = types.AttributeKeyPostOwner
	AttributeKeyPostEditTime        = types.AttributeKeyPostEditTime
	AttributeKeyPollAnswerer        = types.AttributeKeyPollAnswerer
	AttributeKeyReactionOwner       = types.AttributeKeyReactionOwner
	AttributeKeyReactionValue       = types.AttributeKeyReactionValue
	AttributeKeyCreationTime        = types.AttributeKeyCreationTime
	PostSortByID                    = types.PostSortByID
	PostSortByCreationDate          = types.PostSortByCreationDate
	PostSortOrderAscending          = types.PostSortOrderAscending
	PostSortOrderDescending         = types.PostSortOrderDescending
	ModuleName                      = types.ModuleName
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	MaxPostMessageLength            = types.MaxPostMessageLength
	MaxOptionalDataFieldsNumber     = types.MaxOptionalDataFieldsNumber
	MaxOptionalDataFieldValueLength = types.MaxOptionalDataFieldValueLength
	ActionCreatePost                = types.ActionCreatePost
	ActionEditPost                  = types.ActionEditPost
	ActionAnswerPoll                = types.ActionAnswerPoll
	ActionAddPostReaction           = types.ActionAddPostReaction
	ActionRemovePostReaction        = types.ActionRemovePostReaction
	QuerierRoute                    = types.QuerierRoute
	QueryPost                       = types.QueryPost
	QueryPosts                      = types.QueryPosts
	QueryPollAnswers                = types.QueryPollAnswers
	OpWeightMsgCreatePost           = simulation.OpWeightMsgCreatePost
	OpWeightMsgEditPost             = simulation.OpWeightMsgEditPost
	OpWeightMsgAddReaction          = simulation.OpWeightMsgAddReaction
	OpWeightMsgRemoveReaction       = simulation.OpWeightMsgRemoveReaction
	OpWeightMsgAnswerPoll           = simulation.OpWeightMsgAnswerPoll
	DefaultGasValue                 = simulation.DefaultGasValue
)

var (
	// functions aliases
	DecodeStore               = simulation.DecodeStore
	RandomizedGenState        = simulation.RandomizedGenState
	WeightedOperations        = simulation.WeightedOperations
	SimulateMsgAnswerToPoll   = simulation.SimulateMsgAnswerToPoll
	SimulateMsgCreatePost     = simulation.SimulateMsgCreatePost
	SimulateMsgEditPost       = simulation.SimulateMsgEditPost
	SimulateMsgAddReaction    = simulation.SimulateMsgAddReaction
	SimulateMsgRemoveReaction = simulation.SimulateMsgRemoveReaction
	RandomPost                = simulation.RandomPost
	RandomPostData            = simulation.RandomPostData
	RandomReactionData        = simulation.RandomReactionData
	RandomReactionValue       = simulation.RandomReactionValue
	RandomPostID              = simulation.RandomPostID
	RandomMessage             = simulation.RandomMessage
	RandomSubspace            = simulation.RandomSubspace
	RandomMedias              = simulation.RandomMedias
	RandomPollData            = simulation.RandomPollData
	GetAccount                = simulation.GetAccount
	RegisterCodec             = types.RegisterCodec
	NewGenesisState           = types.NewGenesisState
	DefaultGenesisState       = types.DefaultGenesisState
	ValidateGenesis           = types.ValidateGenesis
	NewPollData               = types.NewPollData
	ArePollDataEquals         = types.ArePollDataEquals
	NewUserAnswer             = types.NewUserAnswer
	NewPostMedia              = types.NewPostMedia
	ValidateURI               = types.ValidateURI
	NewPostMedias             = types.NewPostMedias
	DefaultQueryPostsParams   = types.DefaultQueryPostsParams
	NewMsgAddPostReaction     = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction  = types.NewMsgRemovePostReaction
	ParsePostID               = types.ParsePostID
	NewPost                   = types.NewPost
	NewPostResponse           = types.NewPostResponse
	NewReaction               = types.NewReaction
	PostStoreKey              = types.PostStoreKey
	PostCommentsStoreKey      = types.PostCommentsStoreKey
	PostReactionsStoreKey     = types.PostReactionsStoreKey
	PollAnswersStoreKey       = types.PollAnswersStoreKey
	NewMsgCreatePost          = types.NewMsgCreatePost
	NewMsgEditPost            = types.NewMsgEditPost
	NewMsgAnswerPoll          = types.NewMsgAnswerPoll
	ParseAnswerID             = types.ParseAnswerID
	NewPollAnswer             = types.NewPollAnswer
	NewPollAnswers            = types.NewPollAnswers
	NewHandler                = keeper.NewHandler
	NewKeeper                 = keeper.NewKeeper
	NewQuerier                = keeper.NewQuerier

	// variable aliases
	RandomMimeTypes          = simulation.RandomMimeTypes
	RandomHosts              = simulation.RandomHosts
	ModuleCdc                = types.ModuleCdc
	SubspaceRegEx            = types.SubspaceRegEx
	LastPostIDStoreKey       = types.LastPostIDStoreKey
	PostStorePrefix          = types.PostStorePrefix
	PostCommentsStorePrefix  = types.PostCommentsStorePrefix
	PostReactionsStorePrefix = types.PostReactionsStorePrefix
	PollAnswersStorePrefix   = types.PollAnswersStorePrefix
)

type (
	Keeper                   = keeper.Keeper
	PostData                 = simulation.PostData
	ReactionData             = simulation.ReactionData
	GenesisState             = types.GenesisState
	PollData                 = types.PollData
	UserAnswer               = types.UserAnswer
	UserAnswers              = types.UserAnswers
	PostMedia                = types.PostMedia
	PostMedias               = types.PostMedias
	QueryPostsParams         = types.QueryPostsParams
	MsgAddPostReaction       = types.MsgAddPostReaction
	MsgRemovePostReaction    = types.MsgRemovePostReaction
	PostID                   = types.PostID
	PostIDs                  = types.PostIDs
	Post                     = types.Post
	Posts                    = types.Posts
	OptionalData             = types.OptionalData
	KeyValue                 = types.KeyValue
	PostQueryResponse        = types.PostQueryResponse
	Reaction                 = types.Reaction
	Reactions                = types.Reactions
	MsgCreatePost            = types.MsgCreatePost
	MsgEditPost              = types.MsgEditPost
	MsgAnswerPoll            = types.MsgAnswerPoll
	AnswerID                 = types.AnswerID
	PollAnswer               = types.PollAnswer
	PollAnswers              = types.PollAnswers
	PollAnswersQueryResponse = types.PollAnswersQueryResponse
)
