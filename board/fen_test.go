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

var _ = Describe("Board", func() {
	Describe("Fen", func() {
		It("returns the FEN of the board", func() {
			fen := "rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b Qk d3 5 12"
			bd, err := FromFen(fen)
			Expect(err).ToNot(HaveOccurred())

			Expect(bd.Fen()).To(Equal(fen))
		})

		It("returns the FEN of the board", func() {
			fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
			bd, err := FromFen(fen)
			Expect(err).ToNot(HaveOccurred())

			Expect(bd.Fen()).To(Equal(fen))
		})
	})

	Describe("FromFen", func() {
		Context("When there are no errors", func() {
			var board *Board

			BeforeEach(func() {
				fen := "rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kQ d3 5 12"
				var err error
				board, err = FromFen(fen)
				Expect(err).ToNot(HaveOccurred())
			})

			It("has the correct pieces", func() {
				Expect(board.At(square.A1)).To(Equal(piece.WhiteRook))
				Expect(board.At(square.C5)).To(Equal(piece.BlackPawn))
				Expect(board.At(square.A7)).To(Equal(piece.Empty))

				expectedBoard := "" +
					"rnbqk..r" + "\n" +
					".p...ppp" + "\n" +
					"p......." + "\n" +
					".NpPp..." + "\n" +
					"QPP.P.n." + "\n" +
					"P....N.." + "\n" +
					"....KbPP" + "\n" +
					"R.B..B.R" + "\n"
				Expect(board.String()).To(Equal(expectedBoard))
			})

			It("has the correct bitboards", func() {
				expectedBitboard := "" +
					"........" + "\n" +
					"........" + "\n" +
					"........" + "\n" +
					"...X...." + "\n" +
					".XX.X..." + "\n" +
					"X......." + "\n" +
					"......XX" + "\n" +
					"........" + "\n"
				Expect(board.BBPiece(piece.WhitePawn).String()).To(Equal(expectedBitboard))
			})

			It("has the correct side to move", func() {
				Expect(board.SideToMove()).To(Equal(side.Black))
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
				expectedRights := castling.BlackKSide | castling.WhiteQSide
				Expect(board.CastlingRights()).To(Equal(expectedRights))
			})
		})
	})
})
