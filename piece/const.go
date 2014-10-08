package piece

import . "side"

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
  WHITE_PAWN   Piece = PAWN | Piece(WHITE)
  BLACK_PAWN   Piece = PAWN | Piece(BLACK)
  WHITE_KNIGHT Piece = KNIGHT | Piece(WHITE)
  BLACK_KNIGHT Piece = KNIGHT | Piece(BLACK)
  WHITE_BISHOP Piece = BISHOP | Piece(WHITE)
  BLACK_BISHOP Piece = BISHOP | Piece(BLACK)
  WHITE_ROOK   Piece = ROOK | Piece(WHITE)
  BLACK_ROOK   Piece = ROOK | Piece(BLACK)
  WHITE_QUEEN  Piece = QUEEN | Piece(WHITE)
  BLACK_QUEEN  Piece = QUEEN | Piece(BLACK)
  WHITE_KING   Piece = KING | Piece(WHITE)
  BLACK_KING   Piece = KING | Piece(BLACK)
)