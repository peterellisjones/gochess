package traverse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTraverse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Traverse Suite")
}
