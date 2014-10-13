package board_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/bitboard"
	. "github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
	"github.com/peterellisjones/gochess/validate"
)

var _ = Describe("Board", func() {

	It("is valid", func() {
		fen := "rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b - d3 0 12"
		board, err := FromFen(fen)
		Expect(err).ToNot(HaveOccurred())

		err = validate.Board(board)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Board", func() {
		It("Adds the piece", func() {
			board := EmptyBoard()
			board.Add(piece.BlackBishop, square.C4)
			Expect(board.At(square.C4)).To(Equal(piece.BlackBishop))
			Expect(board.BBSide(side.Black)).To(Equal(bitboard.C4))
			Expect(board.BBSide(side.White)).To(Equal(bitboard.Empty))

			board.Add(piece.WhiteQueen, square.H4)
			Expect(board.At(square.H4)).To(Equal(piece.WhiteQueen))
			Expect(board.BBSide(side.Black)).To(Equal(bitboard.C4))
			Expect(board.BBSide(side.White)).To(Equal(bitboard.H4))
		})
	})
})
