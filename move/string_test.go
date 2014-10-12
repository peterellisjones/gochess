package move_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("Move", func() {
	Describe("String", func() {
		It("renders castles", func() {
			Expect(QueenSideCastle.String()).To(Equal("O-O-O"))
			Expect(KingSideCastle.String()).To(Equal("O-O"))
		})

		It("renders regular moves", func() {
			Expect(EncodeMove(square.F7, square.B4).String()).To(Equal("f7b4"))
		})

		It("renders captures", func() {
			Expect(EncodeCapture(square.F7, square.B4).String()).To(Equal("f7xb4"))
		})

		It("renders en-passant captures", func() {
			Expect(EncodeEpCapture(square.F7, square.B4).String()).To(Equal("f7xb4e.p."))
		})

		It("renders promotions", func() {
			move := EncodePromotion(square.F7, square.B4, piece.Knight)
			Expect(move.String()).To(Equal("f7b4=N"))
		})

		It("renders capture-promotions", func() {
			move := EncodeCapturePromotion(square.F7, square.B4, piece.Knight)
			Expect(move.String()).To(Equal("f7xb4=N"))
		})
	})

	Describe("Parse", func() {
		It("parses castles", func() {
			move, err := Parse("O-O-O")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(QueenSideCastle))

			move, err = Parse("O-O")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(KingSideCastle))
		})

		It("parses quiet moves", func() {
			move, err := Parse("a4b6")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(EncodeMove(square.A4, square.B6)))
		})

		It("parses captures", func() {
			move, err := Parse("a4xb6")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(EncodeCapture(square.A4, square.B6)))
		})

		It("parses en-passant captures", func() {
			move, err := Parse("a4xb8e.p.")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(EncodeEpCapture(square.A4, square.B8)))
		})

		It("parses promotions", func() {
			move, err := Parse("a4b8=Q")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(EncodePromotion(square.A4, square.B8, piece.Queen)))
		})

		It("parses capture promotions", func() {
			move, err := Parse("a4xb8=Q")
			Expect(err).ToNot(HaveOccurred())
			Expect(move).To(Equal(EncodeCapturePromotion(square.A4, square.B8, piece.Queen)))
		})
	})
})
