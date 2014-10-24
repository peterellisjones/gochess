package search_test

import (
	"fmt"
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/eval"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/search"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("Seach", func() {
	It("is correct", func() {
		bd, err := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		Expect(err).ToNot(HaveOccurred())

		configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
		Expect(err).ToNot(HaveOccurred())

		ev, err := eval.Load(configPath)
		Expect(err).ToNot(HaveOccurred())

		search := New(bd, ev)
		score := Score{
			Move:  move.EncodeMove(square.B1, square.C3),
			Score: 50,
		}
		Expect(search.BestMove()).To(Equal(score))
	})

	It("is correct", func() {
		bd, err := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
		Expect(err).ToNot(HaveOccurred())

		configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
		Expect(err).ToNot(HaveOccurred())

		ev, err := eval.Load(configPath)
		Expect(err).ToNot(HaveOccurred())

		search := New(bd, ev)
		score := Score{
			Move:  move.EncodeMove(square.B8, square.C6),
			Score: 50,
		}
		Expect(search.BestMove()).To(Equal(score))
	})

	Describe("minimax", func() {
		It("works", func() {
			bd, err := board.FromFen("rnbqkbnr/p1pppppp/8/1p6/P7/8/1PPPPPPP/RNBQKBNR w KQkq - 0 1")
			Expect(err).ToNot(HaveOccurred())

			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())

			ev, err := eval.Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			search := New(bd, ev)

			fmt.Println(search.Negamax(2))
		})
	})

	Describe("alphabeta", func() {
		It("works", func() {
			bd, err := board.FromFen("rnbqkbnr/p1pppppp/8/1p6/P7/8/1PPPPPPP/RNBQKBNR w KQkq - 0 1")
			Expect(err).ToNot(HaveOccurred())

			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())

			ev, err := eval.Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			search := New(bd, ev)

			fmt.Println(search.AlphaBeta(-1000000, 1000000, 2))
		})
	})
})
