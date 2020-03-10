package types

// Magpie module event types
const (
	EventTypePostCreated = "post_created"
	EventTypePostEdited  = "post_edited"

	EventTypeAnsweredPoll = "post_poll_answered"
	EventTypeClosePoll    = "post_poll_closed"

	// Post attributes
	AttributeKeyPostID       = "post_id"
	AttributeKeyPostParentID = "post_parent_id"
	AttributeKeyPostOwner    = "post_owner"
	AttributeKeyPostEditTime = "post_edit_time"

	// Poll attributes
	AttributeKeyPollAnswerer = "poll_answerer"

	// Generic attributes
	AttributeKeyCreationTime = "creation_time"
)
