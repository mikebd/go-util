package command

import (
	"errors"
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

const mockMethodName = "Execute"

func createRegistryWithCommand(commandName string, mock *MockCommand) *Registry {
	r := NewRegistry("test registry")
	r.Add(commandName, mock.Execute)
	return r
}

func TestDefaultRunnerRun(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockCommand)
		setupReg       func(*MockCommand) *Registry
		command        string
		args           []string
		wantErrMessage string
		assertMock     func(*testing.T, *MockCommand)
	}{
		{
			name: "Valid command - no args",
			setupMock: func(m *MockCommand) {
				m.On(mockMethodName, []string{}).Return(nil)
			},
			setupReg: func(m *MockCommand) *Registry {
				return createRegistryWithCommand("command1", m)
			},
			command: "command1",
			args:    []string{},
			assertMock: func(t *testing.T, m *MockCommand) {
				m.AssertExpectations(t)
			},
		},
		{
			name: "Valid command - with args",
			setupMock: func(m *MockCommand) {
				m.On(mockMethodName, []string{"arg1", "arg2"}).Return(nil)
			},
			setupReg: func(m *MockCommand) *Registry {
				return createRegistryWithCommand("command1", m)
			},
			command: "command1",
			args:    []string{"arg1", "arg2"},
			assertMock: func(t *testing.T, m *MockCommand) {
				m.AssertExpectations(t)
			},
		},
		{
			name: "Command execution error",
			setupMock: func(m *MockCommand) {
				err := errors.New("some error")
				m.On(mockMethodName, []string{}).Return(err)
			},
			setupReg: func(m *MockCommand) *Registry {
				return createRegistryWithCommand("command1", m)
			},
			command:        "command1",
			args:           []string{},
			wantErrMessage: "some error",
			assertMock: func(t *testing.T, m *MockCommand) {
				m.AssertExpectations(t)
			},
		},
		{
			name: "Invalid command on empty registry",
			setupReg: func(_ *MockCommand) *Registry {
				return NewRegistry("test registry")
			},
			command:        "invalidCommand",
			args:           []string{},
			wantErrMessage: "command registry test registry, missing command: invalidCommand",
			assertMock: func(t *testing.T, m *MockCommand) {
				m.AssertNotCalled(t, mockMethodName)
			},
		},
		{
			name: "Invalid command on non-empty registry",
			setupReg: func(m *MockCommand) *Registry {
				return createRegistryWithCommand("command1", m)
			},
			command:        "invalidCommand",
			args:           []string{},
			wantErrMessage: "command registry test registry, missing command: invalidCommand",
			assertMock: func(t *testing.T, m *MockCommand) {
				m.AssertNotCalled(t, mockMethodName)
			},
		},
		{
			name:           "Nil registry",
			command:        "command1",
			args:           []string{},
			wantErrMessage: "nil registry",
			assertMock: func(t *testing.T, m *MockCommand) {
				m.AssertNotCalled(t, mockMethodName)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCmd := new(MockCommand)
			if tt.setupMock != nil {
				tt.setupMock(mockCmd)
			}

			var reg *Registry
			if tt.setupReg != nil {
				reg = tt.setupReg(mockCmd)
			}

			err := DefaultRunner.Run(reg, tt.command, tt.args...)

			if tt.wantErrMessage != "" {
				assert.EqualError(t, err, tt.wantErrMessage)
			} else {
				assert.NoError(t, err)
			}

			tt.assertMock(t, mockCmd)
		})
	}
}
