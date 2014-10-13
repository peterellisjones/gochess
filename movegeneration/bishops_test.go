package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("AddBishopMoves", func() {
	cases := map[string][]string{
		"8/k7/8/2b5/1R6/2K5/8/8 b - - 0 1": []string{
			"c5b6", "c5d6", "c5e7", "c5xb4",
			"c5d4", "c5f8", "c5e3", "c5f2",
			"c5g1",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
		generator := New(bd, list)
		generator.AddBishopMoves(bd.SideToMove())
	})
})
