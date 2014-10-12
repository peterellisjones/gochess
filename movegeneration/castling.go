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
			bb.SquareB8 | bb.SquareC8 | bb.SquareD8,
			bb.SquareF8 | bb.SquareG8,
			bb.SquareB1 | bb.SquareC1 | bb.SquareD1,
			bb.SquareF1 | bb.SquareG1,
		}[castle]

		if (moveMask & gen.board.BBEmpty()) != moveMask {
			return
		}

		// check king doesn't pass through check
		checkMask := [4]bb.Bitboard{
			bb.SquareC8 | bb.SquareD8 | bb.SquareE8,
			bb.SquareE8 | bb.SquareF8 | bb.SquareG8,
			bb.SquareC1 | bb.SquareD1 | bb.SquareE1,
			bb.SquareE1 | bb.SquareF1 | bb.SquareG1,
		}[castle]

		if AreSquaresAttacked(gen.board, checkMask, side.Other()) {
			return
		}

		// OK, add move
		mv := []move.Move{move.QueenSideCastle, move.KingSideCastle}[castle&1]
		gen.list.Add(mv)
	})
}
