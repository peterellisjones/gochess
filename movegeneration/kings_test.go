package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("GenerateKingMoves", func() {
	cases := map[string][]string{
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -": []string{
			"e1f1",
			"e1d1",
		},
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": []string{},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
		generator := New(bd, list)
		generator.AddKingMoves(bd.SideToMove())
	})
})
