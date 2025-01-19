package risks

type State string

const (
	Open          State = "open"
	Closed        State = "closed"
	Accepted      State = "accepted"
	Investigating State = "investigating"
)

type Risk struct {
	ID          string `json:"id"`
	State       State  `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Validate if state has a valid input
func isValidState(state State) bool {
	switch state {
	case Open, Closed, Accepted, Investigating:
		return true
	default:
		return false
	}
}
