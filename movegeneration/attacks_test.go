package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/side"
)

var _ = Describe("AttackedSquares", func() {
	It("returns the set of attacked squares", func() {
		board, _ := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")

		expectedAttacks := "" +
			"........\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			"XXXXXXXX\n" +
			"XXXXXXXX\n" +
			".XXXXXX.\n"
		attacks := AttackedSquares(board, side.White)
		Expect(attacks.String()).To(Equal(expectedAttacks))

	})
})
