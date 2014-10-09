package bitboard_test

import (
	. "github.com/peterellisjones/gochess/bitboard"
	sq "github.com/peterellisjones/gochess/square"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bitboard", func() {
	Describe("IsSet", func() {
		It("returns true if the bit is set", func() {
			bb := Bitboard(FileB)
			Expect(bb.IsSet(sq.A1)).To(BeFalse())
			Expect(bb.IsSet(sq.B1)).To(BeTrue())
			Expect(bb.IsSet(sq.C1)).To(BeFalse())
			Expect(bb.IsSet(sq.D1)).To(BeFalse())
		})
	})
})
