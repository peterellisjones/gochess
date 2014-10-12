package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("GenerateKnightMoves", func() {
	cases := map[string][]string{
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -": []string{
			"c3b1", "c3d1", "c3a4", "c3b5",
			"e5xg6", "e5xd7", "e5xf7", "e5d3",
			"e5c4", "e5g4", "e5c6",
		},
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": []string{
			"b1a3", "b1c3", "g1f3", "g1h3",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
		generator := New(bd, list)
		generator.AddKnightMoves(bd.SideToMove())
	})
})
