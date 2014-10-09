package board

import (
	"github.com/peterellisjones/gochess/fen"
	"github.com/peterellisjones/gochess/square"
)

func FromFen(str string) (*Board, error) {
	parts, err := fen.FenParts(str)
	if err != nil {
		return nil, err
	}

	board := EmptyBoard()
	board.sideToMove = parts.SideToMove
	board.irrev.castlingRights = parts.CastlingRights
	board.irrev.epSquare = parts.EpSquare
	board.irrev.halfMoveClock = parts.HalfMoveClock

	board.fullMoveNumber = parts.FullMoveNumber

	for i, piece := range parts.Board {
		board.Add(piece, square.Square(i))
	}

	return board, nil
}
