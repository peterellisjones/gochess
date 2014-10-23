package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/movegeneration"
)

var _ = Describe("GenerateKingMoves", func() {
	cases := map[string][]string{
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -": []string{
			"e1f1",
			"e1d1",
		},
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": []string{},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board) []move.Move {
		gen := New(bd)
		return gen.KingMoves(bd.SideToMove())
	})
})
