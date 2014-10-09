package piece

import (
	"github.com/peterellisjones/gochess/side"
)

type Piece uint8

const (
	EMPTY Piece = Piece(0)
	ERROR Piece = Piece(1)
)

const (
	PAWN   Piece = Piece(2)
	KNIGHT Piece = Piece(4)
	BISHOP Piece = Piece(6)
	ROOK   Piece = Piece(8)
	QUEEN  Piece = Piece(10)
	KING   Piece = Piece(12)
)

const (
	WHITE_PAWN   Piece = PAWN | Piece(side.WHITE)
	BLACK_PAWN   Piece = PAWN | Piece(side.BLACK)
	WHITE_KNIGHT Piece = KNIGHT | Piece(side.WHITE)
	BLACK_KNIGHT Piece = KNIGHT | Piece(side.BLACK)
	WHITE_BISHOP Piece = BISHOP | Piece(side.WHITE)
	BLACK_BISHOP Piece = BISHOP | Piece(side.BLACK)
	WHITE_ROOK   Piece = ROOK | Piece(side.WHITE)
	BLACK_ROOK   Piece = ROOK | Piece(side.BLACK)
	WHITE_QUEEN  Piece = QUEEN | Piece(side.WHITE)
	BLACK_QUEEN  Piece = QUEEN | Piece(side.BLACK)
	WHITE_KING   Piece = KING | Piece(side.WHITE)
	BLACK_KING   Piece = KING | Piece(side.BLACK)
)
