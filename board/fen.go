package board

import (
	"github.com/peterellisjones/gochess/fen"
	"github.com/peterellisjones/gochess/square"
)

// FromFen returns a board given a FEN string
func FromFen(str string) (*Board, error) {
	parts, err := fen.GetParts(str)
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
