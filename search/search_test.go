package search_test

import (
	"fmt"
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/eval"
	. "github.com/peterellisjones/gochess/search"
)

var _ = Describe("Seach", func() {

	type Case struct {
		fen   string
		move  string
		depth int
	}

	cases := []Case{
		// Case{
		// 	fen:   "rnbqkbnr/pppppppp/8/8/q7/1P6/P1PPPPPP/RNBQKBNR w q",
		// 	move:  "b3xa4",
		// 	depth: 1,
		// },
		// Case{
		// 	fen:   "r3k3/2qb1pp1/p1N1p1r1/1B6/Q2PN3/B3nP2/1R3RPp/7K w q",
		// 	move:  "a4b4",
		// 	depth: 6,
		// },
		// Case{
		// 	fen:   "2r1k2r/p1B2ppp/8/P2b4/1b2N3/3p1P2/6PP/2R1R1K1 w k",
		// 	move:  "e4f6",
		// 	depth: 3,
		// },
		Case{
			fen:   "2r2k1r/p1B2ppp/5N2/P2b4/1b6/3p1P2/6PP/2R1R1K1 w k",
			move:  "c7d6",
			depth: 3,
		},
		// Case{
		// 	fen:   "2r1k2r/p1B2ppp/5N2/P2b4/1b6/3p1P2/6PP/2R1R1K1 b k",
		// 	move:  "e8f8",
		// 	depth: 3,
		// },
		// Case{
		// 	fen:   "2r1k2r/p1B2ppp/5N2/P2b4/1b6/3p1P2/6PP/2R1R1K1 b k",
		// 	move:  "e8f8",
		// 	depth: 2,
		// },
		// Case{
		// 	fen:   "2r1k2r/p1B2ppp/5N2/b2b4/8/3p1P2/6PP/2R1R1K1 w k",
		// 	move:  "e1xe8",
		// 	depth: 1,
		// },
	}

	for _, c := range cases {
		cLocal := c
		Describe(cLocal.fen, func() {

			It(fmt.Sprintf("It chooses %s", cLocal.move), func() {
				bd, err := board.FromFen(cLocal.fen)
				Expect(err).ToNot(HaveOccurred())

				configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
				Expect(err).ToNot(HaveOccurred())

				ev, err := eval.Load(configPath)
				Expect(err).ToNot(HaveOccurred())

				search := New(bd, ev)

				bestMove, score := search.Negamax(cLocal.depth)
				fmt.Println("RESULTS")
				fmt.Println(bestMove)
				fmt.Println(score)

				Expect(bestMove.String()).To(Equal(cLocal.move))
			})
		})
	}

	// It("is correct", func() {
	// 	bd, err := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	// 	Expect(err).ToNot(HaveOccurred())
	//
	// 	configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
	// 	Expect(err).ToNot(HaveOccurred())
	//
	// 	ev, err := eval.Load(configPath)
	// 	Expect(err).ToNot(HaveOccurred())
	//
	// 	search := New(bd, ev)
	// 	score := Score{
	// 		Move:  move.EncodeMove(square.B1, square.C3),
	// 		Score: 50,
	// 	}
	// 	Expect(search.BestMove()).To(Equal(score))
	// })
	//
	// It("is correct", func() {
	// 	bd, err := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
	// 	Expect(err).ToNot(HaveOccurred())
	//
	// 	configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
	// 	Expect(err).ToNot(HaveOccurred())
	//
	// 	ev, err := eval.Load(configPath)
	// 	Expect(err).ToNot(HaveOccurred())
	//
	// 	search := New(bd, ev)
	// 	score := Score{
	// 		Move:  move.EncodeMove(square.B8, square.C6),
	// 		Score: 50,
	// 	}
	// 	Expect(search.BestMove()).To(Equal(score))
	// })
	//
	// Describe("minimax", func() {
	// 	It("works", func() {
	// 		bd, err := board.FromFen("rnbqkbnr/p1pppppp/8/1p6/P7/8/1PPPPPPP/RNBQKBNR w KQkq - 0 1")
	// 		Expect(err).ToNot(HaveOccurred())
	//
	// 		configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
	// 		Expect(err).ToNot(HaveOccurred())
	//
	// 		ev, err := eval.Load(configPath)
	// 		Expect(err).ToNot(HaveOccurred())
	//
	// 		search := New(bd, ev)
	//
	// 		fmt.Println(search.Negamax(2))
	// 	})
	// })
	//
	// Describe("alphabeta", func() {
	// 	It("works", func() {
	// 		bd, err := board.FromFen("rnbqkbnr/p1pppppp/8/1p6/P7/8/1PPPPPPP/RNBQKBNR w KQkq - 0 1")
	// 		Expect(err).ToNot(HaveOccurred())
	//
	// 		configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
	// 		Expect(err).ToNot(HaveOccurred())
	//
	// 		ev, err := eval.Load(configPath)
	// 		Expect(err).ToNot(HaveOccurred())
	//
	// 		search := New(bd, ev)
	//
	// 		fmt.Println(search.AlphaBeta(-1000000, 1000000, 2))
	// 	})
	// })
})
