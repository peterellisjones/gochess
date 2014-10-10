package movelist_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMovelist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movelist Suite")
}
