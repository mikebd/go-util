package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleRegistry() {
	r := NewRegistry("a registry")

	r.Add("Z command", func(_ []string) error { return nil })
	r.Add("A command", func(_ []string) error { return nil })
	r.Add("M command", func(_ []string) error { return nil })

	list := r.List()
	fmt.Println(list)
	fmt.Println(r)

	// Output:
	// [A command M command Z command]
	// Command registry a registry: A command, M command, Z command
}

func TestRegistryRun(t *testing.T) {
	r := NewRegistry("some registry")

	r.Add("OK", func(_ []string) error { return nil })
	r.Add("Fail", func(_ []string) error { return fmt.Errorf("fail") })

	tests := []struct {
		commandName string
		wantErr     error
	}{
		{"OK", nil},
		{"Fail", fmt.Errorf("fail")},
		{"Unknown", fmt.Errorf("command registry some registry, missing command: Unknown")},
	}

	for _, tt := range tests {
		t.Run(tt.commandName, func(t *testing.T) {
			err := r.run(tt.commandName, nil)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
