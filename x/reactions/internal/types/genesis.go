package types

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Reactions map[string]PostReactions `json:"reactions"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(reactions map[string]PostReactions) GenesisState {
	return GenesisState{
		Reactions: reactions,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, reactions := range data.Reactions {
		for _, record := range reactions {
			if err := record.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
