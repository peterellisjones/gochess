package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/movegeneration"
)

var _ = Describe("AddCastles", func() {
	cases := map[string][]string{
		"r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1": []string{
			"O-O-O", "O-O",
		},
		"r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w Qkq - 0 1": []string{
			"O-O-O",
		},
		"r3k1qr/pppppppp/8/8/8/8/PPPPPPPP/R3KBqR b KQkq - 0 1": []string{
			"O-O-O",
		},
		"r3k2r/8/8/8/8/8/8/1RK5 b kq - 0 1": []string{
			"O-O-O", "O-O",
		},
		"r3k2r/8/8/8/8/8/8/3RK3 b kq - 0 1": []string{
			"O-O",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board) []move.Move {
		gen := New(bd)
		return gen.Castles(bd.SideToMove())
	})
})
