package board_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("FromFen", func() {
	Context("When there are no errors", func() {
		var board *Board

		BeforeEach(func() {
			fen := "rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kQ d3 5 12"
			var err error
			board, err = FromFen(fen)
			Expect(err).ToNot(HaveOccurred())
		})

		It("has the correct pieces", func() {
			Expect(board.At(square.A1)).To(Equal(piece.WHITE_ROOK))
			Expect(board.At(square.C5)).To(Equal(piece.BLACK_PAWN))
			Expect(board.At(square.A7)).To(Equal(piece.EMPTY))
		})

		It("has the correct bitboards", func() {
			Expect(board.BBPiece(piece.WHITE_PAWN)).To(Equal(0xF))
		})

		It("has the correct side to move", func() {
			Expect(board.SideToMove()).To(Equal(side.BLACK))
		})

		It("has the correct half move clock", func() {
			Expect(board.HalfMoveClock()).To(Equal(5))
		})

		It("has the correct full move clock", func() {
			Expect(board.FullMoveNumber()).To(Equal(12))
		})

		It("has the correct EP square", func() {
			Expect(board.EpSquare()).To(Equal(square.D3))
		})

		It("has the correct castling rights", func() {
			expectedRights := castling.BLACK_KSIDE | castling.WHITE_QSIDE
			Expect(board.CastlingRights()).To(Equal(expectedRights))
		})
	})
})
