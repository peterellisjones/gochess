package eval

import (
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

type Eval struct {
	pieceValues       [6]int
	pieceSquareValues [6][64]int
	castleValues      [2]int
}

func (eval *Eval) PieceValue(pc piece.Piece) int {
	return eval.pieceValues[pc.Idx()>>1]
}

func (eval *Eval) PieceSquareValue(pc piece.Piece, sq square.Square) int {
	if pc.Side() == side.Black {
		sq = sq.Flip()
	}
	return eval.pieceSquareValues[pc.Idx()>>1][sq]
}

func (eval *Eval) CastleValue(mv move.Move) int {
	return eval.castleValues[mv.CastleType()]
}
