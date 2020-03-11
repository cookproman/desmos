package types

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Posts       Posts                  `json:"posts"`
	PollAnswers map[string]UserAnswers `json:"poll_answers_details"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(posts Posts) GenesisState {
	return GenesisState{
		Posts: posts,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Posts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, pollAnswers := range data.PollAnswers {
		for _, pollAnswer := range pollAnswers {
			if err := pollAnswer.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}
