package eval_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEval(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eval Suite")
}
