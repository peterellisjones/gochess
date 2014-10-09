package fen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/castling"
	. "github.com/peterellisjones/gochess/fen"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("fenParts", func() {
	It("returns an error if there are the wrong number of rows", func() {
		fen := "rnbqkbnr/pppppppp/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
		_, err := FenParts(fen)
		Expect(err).To(HaveOccurred())
	})

	Context("When there are no errors", func() {
		var parts Parts

		BeforeEach(func() {
			fen := "rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kQ d3 5 12"
			var err error
			parts, err = FenParts(fen)
			Expect(err).ToNot(HaveOccurred())
		})

		It("has the correct side to move", func() {
			Expect(parts.SideToMove).To(Equal(side.BLACK))
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
			expectedRights := castling.BLACK_KSIDE | castling.WHITE_QSIDE
			Expect(parts.CastlingRights).To(Equal(expectedRights))
		})
	})
})
