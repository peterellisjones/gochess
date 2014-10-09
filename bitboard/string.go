package bitboard

import (
	"bytes"
	"github.com/peterellisjones/gochess/square"
)

func framedBoard(squareToChar func(square.Square) byte) string {
	var buffer bytes.Buffer

	buffer.WriteString("  ABCDEFGH\n")

	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			sq := square.Square((row << 3) | col)
			if sq.Col() == 'A' {
				buffer.WriteByte(sq.Row())
				buffer.WriteByte('|')
			}

			buffer.WriteByte(squareToChar(sq))

			if sq.Col() == 'H' {
				buffer.WriteByte('|')
				buffer.WriteByte(sq.Row())
				buffer.WriteByte('\n')
			}
		}
	}

	buffer.WriteString("  ABCDEFGH\n")

	return buffer.String()
}

func (bitboard Bitboard) String() string {
	return framedBoard(func(square square.Square) byte {
		if bitboard.IsSet(square) {
			return 'X'
		}
		return '.'
	})
}
