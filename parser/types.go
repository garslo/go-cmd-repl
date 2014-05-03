package parser

import "github.com/garslo/go-cmd-repl"

type Command struct {
	Name string
	Args []cmd_repl.Argument
}

func NewCommand(name string, args ...cmd_repl.Argument) *Command {
	return &Command{
		Name: name,
		Args: args,
	}
}

type CommandParser interface {
	Parse(string) *Command
}
