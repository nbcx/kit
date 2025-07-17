package cli

import (
	"testing"
)

func TestAskForConfirmation(t *testing.T) {
	type args struct {
		prompt string
		def    bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AskForConfirmation(tt.args.prompt, tt.args.def); got != tt.want {
				t.Errorf("AskForConfirmation() = %v, want %v", got, tt.want)
			}
		})
	}
}

// todo: test
func TestAskForConfirmation2(t *testing.T) {
	// ok := AskForConfirmation("you yes?", false)
	// fmt.Println(">>>", ok)
}
