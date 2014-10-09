package bitboard_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBitboard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bitboard Suite")
}
