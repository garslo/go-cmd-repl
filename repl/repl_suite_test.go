package repl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRepl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repl Suite")
}
