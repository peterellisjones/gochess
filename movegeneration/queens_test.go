package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("AddQueenMoves", func() {
	cases := map[string][]string{
		"8/2Bq4/PPPPPPPP/8/1N6/6K1/8/kn6 b - - 0 1": []string{
			"d7xc6", "d7xd6", "d7xe6", "d7xc7",
			"d7e7", "d7f7", "d7g7", "d7h7",
			"d7c8", "d7d8", "d7e8",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
		generator := New(bd, list)
		generator.AddQueenMoves(bd.SideToMove())
	})
})
