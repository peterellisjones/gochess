package board

import (
	"github.com/peterellisjones/gochess/fen"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

// FromFen returns a board given a FEN string
func FromFen(str string) (*Board, error) {
	parts, err := fen.PartsFromString(str)
	if err != nil {
		return nil, err
	}

	board := EmptyBoard()
	board.extra.SideToMove = parts.SideToMove
	board.extra.CastlingRights = parts.CastlingRights
	board.extra.EpSquare = parts.EpSquare
	board.extra.HalfMoveClock = parts.HalfMoveClock
	board.extra.FullMoveNumber = parts.FullMoveNumber

	for i, pc := range parts.Board {
		if pc != piece.Empty {
			board.Add(pc, square.Square(i))
		}
	}

	return board, nil
}

// Fen returns the FEN representation of the board
func (board *Board) Fen() string {
	parts := fen.New()
	parts.SideToMove = board.SideToMove()
	parts.CastlingRights = board.CastlingRights()
	parts.EpSquare = board.EpSquare()
	parts.HalfMoveClock = board.HalfMoveClock()
	parts.FullMoveNumber = board.FullMoveNumber()

	for sq := square.Square(0); sq < square.Square(64); sq++ {
		parts.Board[sq] = board.At(sq)
	}

	return parts.String()
}
