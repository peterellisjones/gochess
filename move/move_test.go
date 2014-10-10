package move_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("Move", func() {
	It("correctly encodes quiet moves", func() {
		move := EncodeMove(square.A3, square.D6)
		Expect(move.From()).To(Equal(square.A3))
		Expect(move.To()).To(Equal(square.D6))
		Expect(move.IsQuiet()).To(Equal(true))
		Expect(move.IsCapture()).To(Equal(false))
		Expect(move.IsEpCapture()).To(Equal(false))
		Expect(move.IsDoublePawnPush()).To(Equal(false))
		Expect(move.IsCastle()).To(Equal(false))
		Expect(move.IsPromotion()).To(Equal(false))
	})

	It("correctly encodes captures", func() {
		move := EncodeCapture(square.H4, square.E3)
		Expect(move.From()).To(Equal(square.H4))
		Expect(move.To()).To(Equal(square.E3))
		Expect(move.IsQuiet()).To(Equal(false))
		Expect(move.IsCapture()).To(Equal(true))
		Expect(move.IsEpCapture()).To(Equal(false))
		Expect(move.IsDoublePawnPush()).To(Equal(false))
		Expect(move.IsCastle()).To(Equal(false))
		Expect(move.IsPromotion()).To(Equal(false))
	})

	It("correctly encodes EP captures", func() {
		move := EncodeEpCapture(square.B4, square.C3)
		Expect(move.From()).To(Equal(square.B4))
		Expect(move.To()).To(Equal(square.C3))
		Expect(move.IsQuiet()).To(Equal(false))
		Expect(move.IsCapture()).To(Equal(true))
		Expect(move.IsEpCapture()).To(Equal(true))
		Expect(move.IsDoublePawnPush()).To(Equal(false))
		Expect(move.IsCastle()).To(Equal(false))
		Expect(move.IsPromotion()).To(Equal(false))
	})

	It("correctly encodes double Pawn pushes", func() {
		move := EncodeDoublePawnPush(square.A2, square.A4)
		Expect(move.From()).To(Equal(square.A2))
		Expect(move.To()).To(Equal(square.A4))
		Expect(move.IsQuiet()).To(Equal(false))
		Expect(move.IsCapture()).To(Equal(false))
		Expect(move.IsEpCapture()).To(Equal(false))
		Expect(move.IsDoublePawnPush()).To(Equal(true))
		Expect(move.IsCastle()).To(Equal(false))
		Expect(move.IsPromotion()).To(Equal(false))
	})

	It("correctly encodes King side castles", func() {
		move := KingSideCastle
		Expect(move.IsQuiet()).To(Equal(false))
		Expect(move.IsCapture()).To(Equal(false))
		Expect(move.IsEpCapture()).To(Equal(false))
		Expect(move.IsDoublePawnPush()).To(Equal(false))
		Expect(move.IsCastle()).To(Equal(true))
		Expect(move.IsPromotion()).To(Equal(false))
	})

	It("correctly encodes Queen side castles", func() {
		move := QueenSideCastle
		Expect(move.IsQuiet()).To(Equal(false))
		Expect(move.IsCapture()).To(Equal(false))
		Expect(move.IsEpCapture()).To(Equal(false))
		Expect(move.IsDoublePawnPush()).To(Equal(false))
		Expect(move.IsCastle()).To(Equal(true))
		Expect(move.IsPromotion()).To(Equal(false))
	})

	It("correctly encodes promotions", func() {
		pieces := []piece.Piece{piece.Knight, piece.Bishop, piece.Rook, piece.Queen}
		for _, piece := range pieces {
			move := EncodePromotion(square.H6, square.H8, piece)
			Expect(move.IsQuiet()).To(Equal(false))
			Expect(move.IsCapture()).To(Equal(false))
			Expect(move.IsEpCapture()).To(Equal(false))
			Expect(move.IsDoublePawnPush()).To(Equal(false))
			Expect(move.IsCastle()).To(Equal(false))
			Expect(move.IsPromotion()).To(Equal(true))
			Expect(move.PromoteTo()).To(Equal(piece))
		}
	})

	It("correctly encodes capture-promotions", func() {
		pieces := []piece.Piece{piece.Knight, piece.Bishop, piece.Rook, piece.Queen}
		for _, piece := range pieces {
			move := EncodeCapturePromotion(square.H6, square.H8, piece)
			Expect(move.IsQuiet()).To(Equal(false))
			Expect(move.IsCapture()).To(Equal(true))
			Expect(move.IsEpCapture()).To(Equal(false))
			Expect(move.IsDoublePawnPush()).To(Equal(false))
			Expect(move.IsCastle()).To(Equal(false))
			Expect(move.IsPromotion()).To(Equal(true))
			Expect(move.PromoteTo()).To(Equal(piece))
		}
	})
})
