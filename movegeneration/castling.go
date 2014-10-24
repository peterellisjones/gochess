package movegeneration

import (
	bb "github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/side"
)

func (gen *Generator) ForEachCastle(side side.Side, fn func(move.Move) bool) bool {
	return gen.board.CastlingRights().RightsForSide(side).ForEach(func(castle uint) bool {

		// check no blocking pieces
		moveMask := [4]bb.Bitboard{
			bb.F1 | bb.G1,
			bb.B1 | bb.C1 | bb.D1,
			bb.F8 | bb.G8,
			bb.B8 | bb.C8 | bb.D8,
		}[castle]

		if (moveMask & gen.board.BBEmpty()) != moveMask {
			return false
		}

		// check king doesn't pass through check
		checkMask := [4]bb.Bitboard{
			bb.E1 | bb.F1 | bb.G1,
			bb.C1 | bb.D1 | bb.E1,
			bb.E8 | bb.F8 | bb.G8,
			bb.C8 | bb.D8 | bb.E8,
		}[castle]

		if AreSquaresAttacked(gen.board, checkMask, side.Other()) {
			return false
		}

		// OK, add move
		mv := []move.Move{move.KingSideCastle, move.QueenSideCastle}[castle&1]
		return fn(mv)
	})
}
