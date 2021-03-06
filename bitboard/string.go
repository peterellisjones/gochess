package bitboard

import (
	"github.com/peterellisjones/gochess/print"
	"github.com/peterellisjones/gochess/square"
)

func (bitboard Bitboard) String() string {
	return print.Board(func(square square.Square) byte {
		if bitboard.IsSet(square) {
			return 'X'
		}
		return '.'
	})
}
