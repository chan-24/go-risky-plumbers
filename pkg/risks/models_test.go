package risks

import "testing"

func Test_isValidState(t *testing.T) {
	tests := []struct {
		state State
		want  bool
	}{
		{Open, true},
		{Closed, true},
		{Accepted, true},
		{Investigating, true},
		{"invalid", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(string(tt.state), func(t *testing.T) {
			if got := isValidState(tt.state); got != tt.want {
				t.Errorf("isValidState() = %v, want %v", got, tt.want)
			}
		})
	}
}
