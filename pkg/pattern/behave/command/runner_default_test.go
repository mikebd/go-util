package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCommand struct {
	mock.Mock
}

// Execute satisfies the Command function signature
func (m *MockCommand) Execute(args []string) error {
	called := m.Called(args)
	return called.Error(0)
}

func TestRunnerRun(t *testing.T) {
	var mockCommand = func() *MockCommand {
		result := MockCommand{}
		result.On("Execute", []string{}).Return(nil)
		result.On("Execute", []string{"arg1"}).Return(nil)
		return &result
	}()

	var registry = func() *Registry {
		result := NewRegistry("a registry")
		result.Add("command1", mockCommand.Execute)
		return result
	}()

	var runner = DefaultRunner

	tests := []struct {
		name    string
		command string
		args    []string
		wantErr bool
	}{
		{"Valid command - no args", "command1", []string{}, false},
		{"Valid command - 1 arg", "command1", []string{"arg1"}, false},
		{"Invalid command", "invalidCommand", []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCommand.Calls = nil
			err := runner.Run(registry, tt.command, tt.args...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				mockCommand.AssertCalled(t, "Execute", tt.args)
				mockCommand.AssertExpectations(t)
			}
		})
	}
}
