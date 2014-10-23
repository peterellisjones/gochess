package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/movegeneration"
)

var _ = Describe("GeneratePawnMoves", func() {
	cases := map[string][]string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": []string{
			"a2a3", "b2b3", "c2c3", "d2d3",
			"e2e3", "f2f3", "g2g3", "h2h3",
			"a2a4:", "b2b4:", "c2c4:", "d2d4:",
			"e2e4:", "f2f4:", "g2g4:", "h2h4:",
		},
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1": []string{
			"a7a6", "b7b6", "c7c6", "d7d6",
			"e7e6", "f7f6", "g7g6", "h7h6",
			"a7a5:", "b7b5:", "c7c5:", "d7d5:",
			"e7e5:", "f7f5:", "g7g5:", "h7h5:",
		},
		"K7/8/8/8/8/2p5/3p4/4k3 b -": []string{
			"c3c2",
			"d2d1=Q", "d2d1=R", "d2d1=B", "d2d1=N",
		},
		"8/8/5pK1/1p4R1/Q7/8/7p/R3k2B b - - 0 1": []string{
			"b5xa4", "b5b4", "f6xg5", "f6f5",
		},
		"k1K5/8/8/8/8/8/2P5/8 w - b3": []string{
			"c2c3", "c2c4:", "c2xb3e.p.",
		},
		"KN1nN1k1/4P3/8/8/8/8/8/8 w - b3 0 1": []string{
			"e7xd8=Q", "e7xd8=R", "e7xd8=B", "e7xd8=N",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board) []move.Move {
		gen := New(bd)
		return gen.PawnMoves(bd.SideToMove())
	})
})
