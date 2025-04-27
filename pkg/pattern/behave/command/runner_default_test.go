package command

import "testing"

var registry = func() *Registry {
	result := NewRegistry("a registry")
	result.Add("command1", func(_ []string) error { return nil })
	return result
}()

var runner = DefaultRunner

func TestRunnerRun(t *testing.T) {
	tests := []struct {
		name    string
		command string
		args    []string
		wantErr bool
	}{
		{"Valid command", "command1", []string{}, false},
		{"Invalid command", "invalidCommand", []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runner.Run(registry, tt.command, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("runner.run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
