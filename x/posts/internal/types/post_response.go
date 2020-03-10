package types

import (
	"encoding/json"
	"fmt"
	"github.com/desmos-labs/desmos/x/reactions"
	"strings"
)

// PostQueryResponse represents the data of a post
// that is returned to user upon a query
type PostQueryResponse struct {
	Post
	PollAnswers   []UserAnswer            `json:"poll_answers,omitempty"`
	PostReactions reactions.PostReactions `json:"reactions"`
	Children      PostIDs                 `json:"children"`
}

// String implements fmt.Stringer
func (response PostQueryResponse) String() string {
	out := "ID - [PostReactions] [Children] \n"
	out += fmt.Sprintf("%s - [%s] [%s] \n", response.Post.PostID, response.PostReactions, response.Children)
	return strings.TrimSpace(out)
}

func NewPostResponse(post Post, pollAnswers []UserAnswer, reactions reactions.PostReactions, children PostIDs) PostQueryResponse {
	return PostQueryResponse{
		Post:          post,
		PollAnswers:   pollAnswers,
		PostReactions: reactions,
		Children:      children,
	}
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (response PostQueryResponse) MarshalJSON() ([]byte, error) {
	type temp PostQueryResponse
	return json.Marshal(temp(response))
}

// UnmarshalJSON implements json.Unmarshaler as Amino does
// not respect default json composition
func (response *PostQueryResponse) UnmarshalJSON(data []byte) error {
	type postResponse PostQueryResponse
	var temp postResponse
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*response = PostQueryResponse(temp)
	return nil
}
