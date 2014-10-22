package board_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("Board", func() {

	Describe("Piece Lists", func() {
		It("Adds piece correctly", func() {
			board := EmptyBoard()

			Expect(board.PieceListIndex(square.C4)).To(Equal(0))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(0))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.Square(0)))

			board.Add(piece.BlackBishop, square.C4)
			Expect(board.At(square.C4)).To(Equal(piece.BlackBishop))

			Expect(board.PieceListIndex(square.C4)).To(Equal(0))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(1))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))

			board.Add(piece.BlackBishop, square.A3)
			Expect(board.At(square.C4)).To(Equal(piece.BlackBishop))

			Expect(board.PieceListIndex(square.A3)).To(Equal(1))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(2))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))
			Expect(board.PieceList(piece.BlackBishop)[1]).To(Equal(square.A3))
		})

		It("Moves a piece correctly", func() {
			board := EmptyBoard()

			board.Add(piece.BlackBishop, square.C4)
			board.Add(piece.BlackBishop, square.B5)
			board.Add(piece.BlackBishop, square.A3)

			Expect(board.PieceListIndex(square.A3)).To(Equal(2))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(3))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))
			Expect(board.PieceList(piece.BlackBishop)[1]).To(Equal(square.B5))
			Expect(board.PieceList(piece.BlackBishop)[2]).To(Equal(square.A3))

			board.Move(square.C4, square.H8)

			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.H8))
			Expect(board.PieceListIndex(square.H8)).To(Equal(0))
		})

		It("Removes pieces correctly", func() {
			board := EmptyBoard()

			board.Add(piece.BlackBishop, square.C4)
			board.Add(piece.BlackBishop, square.B5)
			board.Add(piece.BlackBishop, square.A3)

			Expect(board.PieceListIndex(square.A3)).To(Equal(2))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(3))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))
			Expect(board.PieceList(piece.BlackBishop)[1]).To(Equal(square.B5))
			Expect(board.PieceList(piece.BlackBishop)[2]).To(Equal(square.A3))

			board.Remove(square.C4)

			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(2))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.A3))
			Expect(board.PieceListIndex(square.A3)).To(Equal(0))
			Expect(board.PieceList(piece.BlackBishop)[1]).To(Equal(square.B5))
			Expect(board.PieceListIndex(square.B5)).To(Equal(1))
		})

		It("Removes pieces correctly", func() {
			board := EmptyBoard()

			board.Add(piece.BlackBishop, square.C4)
			board.Add(piece.BlackBishop, square.B5)
			board.Add(piece.BlackBishop, square.A3)

			Expect(board.PieceListIndex(square.A3)).To(Equal(2))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(3))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))
			Expect(board.PieceList(piece.BlackBishop)[2]).To(Equal(square.A3))

			board.Remove(square.B5)

			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(2))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))
			Expect(board.PieceListIndex(square.C4)).To(Equal(0))
			Expect(board.PieceList(piece.BlackBishop)[1]).To(Equal(square.A3))
			Expect(board.PieceListIndex(square.A3)).To(Equal(1))
		})

		It("Removes pieces correctly", func() {
			board := EmptyBoard()

			board.Add(piece.BlackBishop, square.C4)

			Expect(board.PieceListIndex(square.C4)).To(Equal(0))
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(1))
			Expect(board.PieceList(piece.BlackBishop)[0]).To(Equal(square.C4))

			board.Remove(square.C4)
			Expect(board.PieceListSize(piece.BlackBishop)).To(Equal(0))
		})
	})
})
