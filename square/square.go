package square

import (
	"errors"
)

// Square represents a square on a chess board
type Square uint8

// represents rows
const (
	Row1 = byte(iota)
	Row2
	Row3
	Row4
	Row5
	Row6
	Row7
	Row8
)

var columns = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
var rows = []byte{'1', '2', '3', '4', '5', '6', '7', '8'}

// Col returns the character representing that square's column
func (square Square) Col() byte {
	return columns[square&7]
}

// Row returns the character representing that square's row
func (square Square) Row() byte {
	return rows[(square&7)>>3]
}

var squareNames = [65]string{
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"-",
}

// Flip inverts the square
func (square Square) Flip() Square {
	return square ^ Square(56)
}

// ForEach iterates over all squares
func ForEach(fn func(Square)) {
	for sq := A1; sq <= H8; sq++ {
		fn(sq)
	}
}

// CircularTranslate transaltes the square mod 64
func (square Square) CircularTranslate(diff Square) Square {
	return (square + diff) & 63
}

func (square Square) String() string {
	return squareNames[square]
}

// Parse returns a square given a string representation (upper or lower case)
func Parse(str string) (Square, error) {

	if str == "-" {
		return Null, nil
	}

	if len(str) != 2 {
		return Null, errors.New("Square must be of format '<col><row>' eg 'A4'")
	}

	colStr := str[0]
	rowStr := str[1]
	square := Square(0)

	if colStr >= 'A' && colStr <= 'H' {
		square = Square(colStr - 'A')
	} else if colStr >= 'a' && colStr <= 'h' {
		square = Square(colStr - 'a')
	} else {
		return Null, errors.New("Column must be in range A..H or a..h")
	}

	if rowStr >= '1' && rowStr <= '8' {
		square |= Square((rowStr - '1') << 3)
	} else {
		return Null, errors.New("Row must be in range 1..8")
	}

	return square, nil
}
