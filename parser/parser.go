package parser

import (
	"strings"

	"github.com/garslo/go-cmd-repl"
)

type commandParser struct{}

func NewCommandParser() *commandParser {
	return &commandParser{}
}

func (p *commandParser) Parse(cmd string) *Command {
	parts := strings.Split(cmd, " ")
	if len(parts) == 0 {
		return NewCommand("")
	}
	args := make([]cmd_repl.Argument, len(parts)-1)
	for i, arg := range parts[1:] {
		args[i] = cmd_repl.NewArgument(arg)
	}
	return NewCommand(parts[0], args...)
}
