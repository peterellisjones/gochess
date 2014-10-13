package movegeneration

import (
	bb "github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/side"
)

// AddCastles generates castles
func (gen *Generator) AddCastles(side side.Side) {
	gen.board.CastlingRights().RightsForSide(side).ForEach(func(castle uint) {

		// check no blocking pieces
		moveMask := [4]bb.Bitboard{
			bb.F1 | bb.G1,
			bb.B1 | bb.C1 | bb.D1,
			bb.F8 | bb.G8,
			bb.B8 | bb.C8 | bb.D8,
		}[castle]

		if (moveMask & gen.board.BBEmpty()) != moveMask {
			return
		}

		// check king doesn't pass through check
		checkMask := [4]bb.Bitboard{
			bb.E1 | bb.F1 | bb.G1,
			bb.C1 | bb.D1 | bb.E1,
			bb.E8 | bb.F8 | bb.G8,
			bb.C8 | bb.D8 | bb.E8,
		}[castle]

		if AreSquaresAttacked(gen.board, checkMask, side.Other()) {
			return
		}

		// OK, add move
		mv := []move.Move{move.KingSideCastle, move.QueenSideCastle}[castle&1]
		gen.list.Add(mv)
	})
}
