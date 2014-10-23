package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/movegeneration"
)

var _ = Describe("GenerateRookMoves", func() {
	cases := map[string][]string{
		"1k6/8/8/5K2/8/3n4/5R1n/8 w -": []string{
			"f2f1", "f2g2", "f2xh2", "f2e2",
			"f2d2", "f2c2", "f2b2", "f2a2", "f2f3",
			"f2f4",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board) []move.Move {
		gen := New(bd)
		return gen.RookMoves(bd.SideToMove())
	})
})
