package parser_test

import (
	"strings"
	. "github.com/garslo/go-cmd-repl/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var (
		parser CommandParser
	)

	BeforeEach(func() {
		parser = NewCommandParser()
	})

	It("should parse a simple command", func() {
		commandString := "transmogrify"
		cmd := parser.Parse(commandString)
		Expect(cmd.Name).To(Equal(commandString))
	})

	It("should parse a more complicated command", func() {
		commandParts := []string{"transmogrify", "1", "foo", "1.45"}
		cmd := parser.Parse(strings.Join(commandParts, " "))
		Expect(cmd.Name).To(Equal(commandParts[0]))
		Expect(cmd.Args).To(HaveLen(len(commandParts[1:])))
	})
})
