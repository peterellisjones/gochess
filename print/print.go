package print

import (
	"bytes"
	"github.com/peterellisjones/gochess/square"
)

// Board returns the object represented as a 8x8 board
func Board(squareToChar func(square.Square) byte) string {
	var buffer bytes.Buffer
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			sq := square.Square((row << 3) | col)
			buffer.WriteByte(squareToChar(sq))
			if sq.Col() == 'H' {
				buffer.WriteByte('\n')
			}
		}
	}

	return buffer.String()
}

// FramedBoard returns the object with a row and column borders
func FramedBoard(squareToChar func(square.Square) byte) string {
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
