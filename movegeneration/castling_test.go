package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("AddCastles", func() {
	cases := map[string][]string{
		"r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1": []string{
			"O-O-O", "O-O",
		},
		"r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w Qkq - 0 1": []string{
			"O-O-O",
		},
		"r3k1qr/pppppppp/8/8/8/8/PPPPPPPP/R3K1qR b KQkq - 0 1": []string{
			"O-O-O",
		},
		"r3k2r/8/8/8/8/8/8/1R6 b KQkq - 0 1": []string{
			"O-O-O", "O-O",
		},
		"r3k2r/8/8/8/8/8/8/3R4 b KQkq - 0 1": []string{
			"O-O",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
		generator := New(bd, list)
		generator.AddCastles(bd.SideToMove())
	})
})
