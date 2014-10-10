package bitboard_test

import (
	. "github.com/peterellisjones/gochess/bitboard"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bitboard", func() {
	Describe("String", func() {
		It("returns a nice representation of the bitboard", func() {
			expectedOutput := "" +
				"XXXXXXXX\n" +
				".X......\n" +
				".X......\n" +
				".X......\n" +
				".X......\n" +
				".X......\n" +
				".X......\n" +
				"XXXXXXXX\n"
			bb := Row1 | Row8 | FileB
			Expect(bb.String()).To(Equal(expectedOutput))
		})
	})
})
