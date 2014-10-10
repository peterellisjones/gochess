package piece

import (
	"github.com/peterellisjones/gochess/side"
)

// Piece represents a chess piece
type Piece uint8

// possible pieces
const (
	Empty  Piece = Piece(0)
	Error  Piece = Piece(1)
	Pawn   Piece = Piece(2)
	Knight Piece = Piece(4)
	Bishop Piece = Piece(6)
	Rook   Piece = Piece(8)
	Queen  Piece = Piece(10)
	King   Piece = Piece(12)
)

// possible pieces including side information
const (
	WhitePawn   Piece = Pawn | Piece(side.White)
	BlackPawn   Piece = Pawn | Piece(side.Black)
	WhiteKnight Piece = Knight | Piece(side.White)
	BlackKnight Piece = Knight | Piece(side.Black)
	WhiteBishop Piece = Bishop | Piece(side.White)
	BlackBishop Piece = Bishop | Piece(side.Black)
	WhiteRook   Piece = Rook | Piece(side.White)
	BlackRook   Piece = Rook | Piece(side.Black)
	WhiteQueen  Piece = Queen | Piece(side.White)
	BlackQueen  Piece = Queen | Piece(side.Black)
	WhiteKing   Piece = King | Piece(side.White)
	BlackKing   Piece = King | Piece(side.Black)
)
