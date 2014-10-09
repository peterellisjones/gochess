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
				"  ABCDEFGH" + "\n" +
				"1|XXXXXXXX|1" + "\n" +
				"1|.X......|1" + "\n" +
				"1|.X......|1" + "\n" +
				"1|.X......|1" + "\n" +
				"1|.X......|1" + "\n" +
				"1|.X......|1" + "\n" +
				"1|.X......|1" + "\n" +
				"1|XXXXXXXX|1" + "\n" +
				"  ABCDEFGH" + "\n"
			bb := Row1 | Row8 | FileB
			Expect(bb.String()).To(Equal(expectedOutput))
		})
	})
})
