package movegeneration

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
)

// InCheck returns true if the side is in check
func InCheck(bd *board.Board, side side.Side) bool {
	king := bd.BBPiece(piece.ForSide(piece.King, side))
	return AreSquaresAttacked(bd, king, side.Other())
}
