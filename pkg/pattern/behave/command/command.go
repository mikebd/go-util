package command

// Command is a function that can be executed.
type Command func(args []string) error
