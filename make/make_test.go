package make_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/make"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/validate"
)

var _ = Describe("Make & UnMake", func() {
	cases := []map[string]string{
		map[string]string{
			"initial":     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			"final":       "rnbqkbnr/pppppppp/8/8/8/1P6/P1PPPPPP/RNBQKBNR b KQkq - 0 1",
			"move":        "b2b3",
			"description": "regular pawn move, reset HMC on pawn move",
		},
		map[string]string{
			"initial":     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			"final":       "rnbqkbnr/pppppppp/8/8/8/2N5/PPPPPPPP/R1BQKBNR b KQkq - 1 1",
			"move":        "b1c3",
			"description": "regular move",
		},
		map[string]string{
			"initial":     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			"final":       "rnbqkbnr/pppppppp/8/8/1P6/8/P1PPPPPP/RNBQKBNR b KQkq b3 0 1",
			"move":        "b2b4:",
			"description": "sets ep square, resets HMC on pawn move",
		},
		map[string]string{
			"initial":     "rnbqkbnr/pppppp1p/8/6p1/5P2/8/PPPPP1PP/RNBQKBNR b KQkq - 4 10",
			"final":       "rnbqkbnr/pppppp1p/8/8/5p2/8/PPPPP1PP/RNBQKBNR w KQkq - 0 11",
			"move":        "g5xf4",
			"description": "resets HMC on capture",
		},
		map[string]string{
			"initial":     "rnbqkbnr/pppp1ppp/8/4pP2/8/8/PPPPP1PP/RNBQKBNR w KQkq e6 0 1",
			"final":       "rnbqkbnr/pppp1ppp/4P3/8/8/8/PPPPP1PP/RNBQKBNR b KQkq - 0 1",
			"move":        "f5xe6e.p.",
			"description": "EP capture",
		},
		map[string]string{
			"initial":     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKB1R w KQkq - 0 1",
			"final":       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBR1 b Qkq - 1 1",
			"move":        "h1g1",
			"description": "cancels castle",
		},
		map[string]string{
			"initial":     "r1bqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			"final":       "1rbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQk - 1 2",
			"move":        "a8b8",
			"description": "cancels castle",
		},
		map[string]string{
			"initial":     "rnb1kbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			"final":       "rnbk1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 1 2",
			"move":        "e8d8",
			"description": "cancels castle",
		},
		map[string]string{
			"initial":     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQK2R w KQkq - 0 1",
			"final":       "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQ1RK1 b kq - 1 1",
			"move":        "O-O",
			"description": "castle",
		},
		map[string]string{
			"initial":     "r3kbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			"final":       "2kr1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 1 2",
			"move":        "O-O-O",
			"description": "castle",
		},
		map[string]string{
			"initial":     "r3kbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			"final":       "2kr1bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQ - 1 2",
			"move":        "O-O-O",
			"description": "castle",
		},
		map[string]string{
			"initial": "r1bqkbnr/pppppp1p/n5p1/8/8/1P6/PBPPPPPP/RN1QKBNR w KQkq - 0 3",
				"final": "r1bqkbnB/pppppp1p/n5p1/8/8/1P6/P1PPPPPP/RN1QKBNR b KQq - 0 3",
				"move": "b2xh8",
				"description": "check castling rights removed",
		},
	}

	for _, c := range cases {
		ex := c

		desc := fmt.Sprintf("makes moves for: %s", ex["description"])
		It(desc, func() {
			bd, err := board.FromFen(ex["initial"])
			Expect(err).ToNot(HaveOccurred())
			err = validate.Board(bd)
			Expect(err).ToNot(HaveOccurred())

			mv, err := move.Parse(ex["move"])
			Expect(err).ToNot(HaveOccurred())

			Make(mv, bd)
			actual := bd.Fen()
			Expect(actual).To(Equal(ex["final"]))
		})

		desc = fmt.Sprintf("un-makes moves for: %s", ex["description"])
		It(desc, func() {
			bd, err := board.FromFen(ex["initial"])
			Expect(err).ToNot(HaveOccurred())
			err = validate.Board(bd)
			Expect(err).ToNot(HaveOccurred())

			mv, err := move.Parse(ex["move"])
			Expect(err).ToNot(HaveOccurred())

			data := Make(mv, bd)

			bd, err = board.FromFen(ex["final"])
			Expect(err).ToNot(HaveOccurred())

			UnMake(bd, data)
			actual := bd.Fen()
			Expect(actual).To(Equal(ex["initial"]))
		})
	}
})
