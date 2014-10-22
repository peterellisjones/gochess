package eval

import (
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

type Values struct {
	pieceValues       [6]int
	pieceSquareValues [6][64]int
	castleValues      [2]int
}

func (values *Values) PieceValue(pc piece.Piece) int {
	return values.pieceValues[pc.Idx()>>1]
}

func (values *Values) PieceSquareValue(pc piece.Piece, sq square.Square) int {
	if pc.Side() == side.Black {
		sq = sq.Flip()
	}
	return values.pieceSquareValues[pc.Idx()>>1][sq]
}

func (values *Values) CastleValue(mv move.Move) int {
	return values.castleValues[mv.CastleType()]
}
