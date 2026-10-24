package commands

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestBenchPromptFlagWithoutValue(t *testing.T) {
	server := reflectingProvider(t)

	tests := []struct {
		name string
		args []string
		want string
	}{
		{"flag with value", []string{"leaky", "--prompt", "Say hi"}, "Say hi"},
		{"trailing flag keeps the default prompt", []string{"leaky", "--prompt"}, benchDefaultPrompt},
		{"trailing short flag keeps the default prompt", []string{"leaky", "-p"}, benchDefaultPrompt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, stdout, _ := leakContext(t, server.URL)
			if _, err := runBench(context.Background(), c, tt.args); err != nil {
				t.Fatalf("runBench: %v", err)
			}
			want := fmt.Sprintf("prompt: %q", tt.want)
			if !strings.Contains(stdout.String(), want) {
				t.Fatalf("output does not contain %s:\n%s", want, stdout.String())
			}
		})
	}
}
