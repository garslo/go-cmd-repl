package repl_test

import (
	"os"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	. "github.com/garslo/go-cmd-repl/repl"
)

var _ = Describe("Repl", func() {
	var (
		repl Repl
	)

	BeforeEach(func() {
		repl = NewReplWithIO(os.Stdout, os.Stdin)
	})

	It("should instantiate", func() {})

	It("should allow us to register a command", func() {

	})
})
