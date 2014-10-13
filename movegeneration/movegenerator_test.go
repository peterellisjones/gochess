package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("GenerateKnightMoves", func() {
	cases := map[string][]string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": []string{
			"b1a3", "b1c3", "g1f3", "g1h3",
			"a2a3", "b2b3", "c2c3", "d2d3",
			"e2e3", "f2f3", "g2g3", "h2h3",
			"a2a4:", "b2b4:", "c2c4:", "d2d4:",
			"e2e4:", "f2f4:", "g2g4:", "h2h4:",
		},
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -": []string{
			"a1b1", "a1c1", "a1d1",
			"e1d1", "e1f1",
			"h1g1", "h1f1",
			"a2a3", "a2a4:",
			"b2b3",
			"d2e3", "d2f4", "d2g5", "d2h6", "d2c1",
			"e2d3", "e2c4", "e2b5", "e2xa6", "e2d1", "e2f1",
			"g2g3", "g2g4:", "g2xh3",
			"c3b1", "c3d1", "c3a4", "c3b5",
			"f3e3", "f3d3", "f3g3", "f3xh3", "f3f4", "f3f5", "f3xf6", "f3g4", "f3h5",
			"d5d6", "d5xe6",
			"e5d3", "e5c4", "e5c6", "e5xd7", "e5xf7", "e5xg6", "e5g4",
			"O-O", "O-O-O",
		},
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq -": []string{
			"h3xg2",
			"b4b3", "b4xc3",
			"a6b5", "a6c4", "a6d3", "a6xe2", "a6b7", "a6c8",
			"b6a4", "b6c4", "b6xd5", "b6c8",
			"e6xd5",
			"f6xd5", "f6xe4", "f6g4", "f6h5", "f6g8", "f6h7",
			"g6g5",
			"c7c6", "c7c5:",
			"d7d6",
			"e7d6", "e7c5", "e7d8", "e7f8",
			"g7f8", "g7h6",
			"a8b8", "a8c8", "a8d8",
			"e8d8", "e8f8",
			"h8g8", "h8f8", "h8h7", "h8h6", "h8h5", "h8h4",
			"O-O", "O-O-O",
		},
	}

	ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
		generator := New(bd, list)
		generator.AddAllMoves(bd.SideToMove())
	})
})
