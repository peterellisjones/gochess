package piece

import (
	"fmt"

	"github.com/peterellisjones/gochess/side"
)

type pieceInfo struct {
	Name     string
	TypeName string
	Char     string
}

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
	Empty:  "empty",
	Knight: "knight",
	Bishop: "bishop",
	Rook:   "rook",
	Queen:  "queen",
	King:   "king",
}

// ForEach executes a callback for each piece
func ForEach(fn func(Piece)) {
	for p := WhitePawn; p <= BlackKing; p++ {
		fn(p)
	}
}

// Side returns the side the piece belongs to
func (piece Piece) Side() side.Side {
	return side.Side(piece & 1)
}

func (piece Piece) TypeName() string {
	return pieceNames[piece.Type()]
}

// Type removes the side information
func (piece Piece) Type() Piece {
	return piece & 0xFE
}

// Idx returns the index of the piece in the array of valid pieces
func (piece Piece) Idx() int {
	return int(piece) - 2
}

// ForSide returns a piece for a given side
func ForSide(piece Piece, side side.Side) Piece {
	return piece.Type() | Piece(side)
}

func (piece Piece) String() string {
	return piece.Side().String() + " " + pieceNames[piece.Type()]
}

// ShortString returns the piece as a one-character string
func (piece Piece) ShortString() string {
	return fmt.Sprintf("%s", piece.Char())
}

// ParseName returns the piece given a lowercase name
func ParseType(name string) Piece {
	return map[string]Piece{
		"pawn":   WhitePawn,
		"knight": WhiteKnight,
		"bishop": WhiteBishop,
		"rook":   WhiteRook,
		"queen":  WhiteQueen,
		"king":   WhiteKing,
	}[name]
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
