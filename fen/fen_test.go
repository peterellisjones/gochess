package fen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/castling"
	. "github.com/peterellisjones/gochess/fen"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("fenParts", func() {
	It("returns an Error if there are the wrong number of rows", func() {
		fen := "rnbqkbnr/pppppppp/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
		_, err := GetParts(fen)
		Expect(err).To(HaveOccurred())
	})

	Context("When there are no Errors", func() {
		var parts Parts

		BeforeEach(func() {
			fen := "rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kQ d3 5 12"
			var err error
			parts, err = GetParts(fen)
			Expect(err).ToNot(HaveOccurred())
		})

		It("has the correct pieces", func() {
			Expect(parts.Board[square.A1]).To(Equal(piece.WhiteRook))
			Expect(parts.Board[square.C5]).To(Equal(piece.BlackPawn))
			Expect(parts.Board[square.A7]).To(Equal(piece.Empty))
		})

		It("has the correct side to move", func() {
			Expect(parts.SideToMove).To(Equal(side.Black))
		})

		It("has the correct half move clock", func() {
			Expect(parts.HalfMoveClock).To(Equal(5))
		})

		It("has the correct full move clock", func() {
			Expect(parts.FullMoveNumber).To(Equal(12))
		})

		It("has the correct EP square", func() {
			Expect(parts.EpSquare).To(Equal(square.D3))
		})

		It("has the correct castling rights", func() {
			expectedRights := castling.BlackKSide | castling.WhiteQSide
			Expect(parts.CastlingRights).To(Equal(expectedRights))
		})
	})
})
