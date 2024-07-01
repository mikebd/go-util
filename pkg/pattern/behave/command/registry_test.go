package command

import (
	"fmt"
	"testing"
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

func TestRegistry_Run(t *testing.T) {
	r := NewRegistry("a registry")

	r.Add("OK", func(_ []string) error { return nil })
	r.Add("Fail", func(_ []string) error { return fmt.Errorf("fail") })

	tests := []struct {
		commandName string
		wantErr     error
	}{
		{"OK", nil},
		{"Fail", fmt.Errorf("fail")},
		{"Unknown", fmt.Errorf("command registry a registry, missing command: Unknown")},
	}

	for _, tt := range tests {
		t.Run(tt.commandName, func(t *testing.T) {
			err := r.Run(tt.commandName, nil)
			if err == nil && tt.wantErr == nil {
				return
			}
			if err.Error() != tt.wantErr.Error() {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}
