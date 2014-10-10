package piece

import (
	"fmt"
	"github.com/peterellisjones/gochess/side"
)

var pieceChars = map[Piece]byte{
	Empty:       '.',
	Error:       'e',
	WhitePawn:   'P',
	BlackPawn:   'p',
	WhiteKnight: 'N',
	BlackKnight: 'n',
	WhiteBishop: 'B',
	BlackBishop: 'b',
	WhiteRook:   'R',
	BlackRook:   'r',
	WhiteQueen:  'Q',
	BlackQueen:  'q',
	WhiteKing:   'K',
	BlackKing:   'k',
}

var charPieces = map[byte]Piece{
	'P': WhitePawn,
	'p': BlackPawn,
	'N': WhiteKnight,
	'n': BlackKnight,
	'B': WhiteBishop,
	'b': BlackBishop,
	'R': WhiteRook,
	'r': BlackRook,
	'Q': WhiteQueen,
	'q': BlackQueen,
	'K': WhiteKing,
	'k': BlackKing,
}

var pieceNames = map[Piece]string{
	Empty:  "Empty",
	Knight: "Knight",
	Bishop: "Bishop",
	Rook:   "Rook",
	Queen:  "Queen",
	King:   "King",
}

// Side returns the side the piece belongs to
func (piece Piece) Side() side.Side {
	return side.Side(piece & 1)
}

// Type removes the side information
func (piece Piece) Type() Piece {
	return piece & 0xFE
}

func (piece Piece) String() string {
	return piece.Side().String() + " " + pieceNames[piece.Type()]
}

// Char returns a character reprentation of the piece
func (piece Piece) Char() byte {
	return pieceChars[piece]
}

// Parse returns a piece given a character
func Parse(chr byte) (Piece, error) {
	piece, ok := charPieces[chr]
	if !ok {
		return piece, fmt.Errorf("Piece not recognized: '%s'", string(chr))
	}
	return piece, nil
}
