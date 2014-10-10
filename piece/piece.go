package piece

import (
	"errors"
	"fmt"
	"github.com/peterellisjones/gochess/side"
)

var pieceChars map[Piece]byte = map[Piece]byte{
	EMPTY:        '.',
	ERROR:        'e',
	WHITE_PAWN:   'P',
	BLACK_PAWN:   'p',
	WHITE_KNIGHT: 'N',
	BLACK_KNIGHT: 'n',
	WHITE_BISHOP: 'B',
	BLACK_BISHOP: 'b',
	WHITE_ROOK:   'R',
	BLACK_ROOK:   'r',
	WHITE_QUEEN:  'Q',
	BLACK_QUEEN:  'q',
	WHITE_KING:   'K',
	BLACK_KING:   'k',
}

var charPieces = map[byte]Piece{
	'P': WHITE_PAWN,
	'p': BLACK_PAWN,
	'N': WHITE_KNIGHT,
	'n': BLACK_KNIGHT,
	'B': WHITE_BISHOP,
	'b': BLACK_BISHOP,
	'R': WHITE_ROOK,
	'r': BLACK_ROOK,
	'Q': WHITE_QUEEN,
	'q': BLACK_QUEEN,
	'K': WHITE_KING,
	'k': BLACK_KING,
}

var pieceNames map[Piece]string = map[Piece]string{
	EMPTY:  "empty",
	KNIGHT: "knight",
	BISHOP: "bishop",
	ROOK:   "rook",
	QUEEN:  "queen",
	KING:   "king",
}

func (piece Piece) Side() side.Side {
	return side.Side(piece & 1)
}

func (piece Piece) Type() Piece {
	return piece & 0xFE
}

func (piece Piece) String() string {
	return piece.Side().String() + " " + pieceNames[piece.Type()]
}

func (piece Piece) Char() byte {
	return pieceChars[piece]
}

func Parse(chr byte) (Piece, error) {
	piece, ok := charPieces[chr]
	if !ok {
		return ERROR, errors.New(fmt.Sprintf("Piece not recognized: '%s'", string(chr)))
	}
	return piece, nil
}
