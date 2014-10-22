package eval_test

import (
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/eval"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("eval", func() {
	Describe("Board", func() {
		It("computes the correct score", func() {
			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())
			values, err := Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			board := board.EmptyBoard()

			board.Add(piece.WhiteBishop, square.C4)
			Expect(board.At(square.C4)).To(Equal(piece.WhiteBishop))

			score := Board(values, board)
			Expect(score).To(Equal(340))
		})

		It("computes the correct score", func() {
			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())
			values, err := Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			board, err := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
			Expect(err).ToNot(HaveOccurred())

			score := Board(values, board)
			Expect(score).To(Equal(0))
		})
	})

	Describe("Move", func() {
		It("computes the correct score", func() {
			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())
			values, err := Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			board, err := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
			Expect(err).ToNot(HaveOccurred())

			mv := move.EncodeDoublePawnPush(square.D2, square.D4)
			score := Move(values, board, mv)
			Expect(score).To(Equal(40))
		})

		It("computes the correct score", func() {
			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())
			values, err := Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			board, err := board.FromFen("r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1")
			Expect(err).ToNot(HaveOccurred())

			mv := move.KingSideCastle
			score := Move(values, board, mv)
			Expect(score).To(Equal(30))

			mv = move.QueenSideCastle
			score = Move(values, board, mv)
			Expect(score).To(Equal(15))
		})
	})
})
