package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/garslo/go-cmd-repl/parser"
)

type repl struct {
	out      io.Writer
	in       io.Reader
	parser   parser.CommandParser
	commands map[string]ActionFunc
}

func New() *repl {
	return NewReplWithIOAndParser(os.Stdout, os.Stdin, parser.NewCommandParser())
}

func NewReplWithIOAndParser(out io.Writer, in io.Reader, parser parser.CommandParser) *repl {
	return &repl{
		out:      out,
		in:       in,
		parser:   parser,
		commands: make(map[string]ActionFunc),
	}
}

func (r *repl) RegisterCommand(name string, action ActionFunc) {
	r.commands[name] = action
}

func (r *repl) Start() error {
	reader := bufio.NewReader(r.in)
	var err error
	for {
		fmt.Fprintf(r.out, "> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		command := r.parser.Parse(line)
		r.runCommand(command)
	}
	return err
}

func (r *repl) runCommand(command *parser.Command) {
	action, ok := r.commands[command.Name]
	if !ok {
		fmt.Fprintf(r.out, "Error: command '%s' not found\n", command.Name)
		return
	}
	fmt.Fprintf(r.out, "%v\n", action(command.Args...))
}
