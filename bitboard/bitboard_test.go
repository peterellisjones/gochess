package bitboard_test

import (
	. "github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/square"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bitboard", func() {
	Describe("IsSet", func() {
		It("returns true if the bit is set", func() {
			bb := Bitboard(FileB)
			Expect(bb.IsSet(square.A1)).To(BeFalse())
			Expect(bb.IsSet(square.B1)).To(BeTrue())
			Expect(bb.IsSet(square.C1)).To(BeFalse())
			Expect(bb.IsSet(square.D1)).To(BeFalse())
		})
	})
})
