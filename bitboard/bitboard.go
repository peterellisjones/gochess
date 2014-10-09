package bitboard

import (
	sq "github.com/peterellisjones/gochess/square"
)

// IsSet returns true if bit is set
func (bitboard Bitboard) IsSet(square sq.Square) bool {
	mask := Bitboard(1) << square
	return mask&bitboard != Empty
}

// ForEachSquare iterates over each square
// call func with (square, isSet) where isSet is true
// if square is occupied
func (bitboard Bitboard) ForEachSquare(fn func(sq.Square, bool)) {
	for i := sq.Square(0); i < sq.Square(64); i++ {
		fn(i, bitboard.IsSet(i))
	}
}
