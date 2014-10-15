package perft_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPerft(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Perft Suite")
}
