package command

import (
	"fmt"
	"sort"
	"strings"
)

// Registry is a directory of available commands
type Registry struct {
	commands map[string]Command
	name     string
}

// NewRegistry creates a new Registry
func NewRegistry(name string) *Registry {
	return &Registry{
		commands: make(map[string]Command),
		name:     name,
	}
}

// Add registers a command in the Registry
func (r *Registry) Add(cmdName string, cmd Command) {
	r.commands[cmdName] = cmd
}

// List returns a list of all registered commands
func (r *Registry) List() []string {
	result := make([]string, 0, len(r.commands))
	for cmdName := range r.commands {
		result = append(result, cmdName)
	}
	sort.Strings(result)
	return result
}

// Name returns the name of the Registry
func (r *Registry) Name() string {
	return r.name
}

// Run executes a command by name
func (r *Registry) Run(cmdName string, args []string) error {
	cmd, ok := r.commands[cmdName]
	if !ok {
		return fmt.Errorf("command registry %s, missing command: %s", r.name, cmdName)
	}
	return cmd(args)
}

// String returns a string representation of the Registry
func (r *Registry) String() string {
	return fmt.Sprintf("Command registry %s: %s", r.name, strings.Join(r.List(), ", "))
}
