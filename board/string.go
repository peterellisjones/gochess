package board

import (
	"github.com/peterellisjones/gochess/print"
	"github.com/peterellisjones/gochess/square"
)

func (board *Board) String() string {
	return print.Board(func(square square.Square) byte {
		return board.At(square).Char()
	})
}
