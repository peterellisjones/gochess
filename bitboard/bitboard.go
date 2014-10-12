package bitboard

import (
	sq "github.com/peterellisjones/gochess/square"
)

// Set sets that bit in the bitboard
func (bitboard Bitboard) Set(square sq.Square) Bitboard {
	mask := Bitboard(1) << square
	return Bitboard(bitboard | mask)
}

// IsSet returns true if bit is set
func (bitboard Bitboard) IsSet(square sq.Square) bool {
	mask := Bitboard(1) << square
	return mask&bitboard != Empty
}

// CircularRightShift does a circular right shift (ie bits are not truncated)
func (bitboard Bitboard) CircularRightShift(square sq.Square) Bitboard {
	right := bitboard >> square
	left := bitboard << (64 - square)
	return right | left
}

// ForEach iterates over each square
// call func with (square, isSet) where isSet is true
// if square is occupied
func (bitboard Bitboard) ForEach(fn func(sq.Square, bool)) {
	for i := sq.Square(0); i < sq.Square(64); i++ {
		fn(i, bitboard.IsSet(i))
	}
}

// ForEachSetBit iterates over each square
// call func with (square, isSet) where isSet is true
// if square is occupied
func (bitboard Bitboard) ForEachSetBit(fn func(sq.Square)) {
	for ; bitboard != 0; bitboard &= bitboard - 1 {
		fn(bitboard.BitScanForward())
	}
}

// BitCount returns the number of set bits
func (bitboard Bitboard) BitCount() int {
	count := 0
	for ; bitboard != 0; bitboard &= bitboard - 1 {
		bitboard.BitScanForward()
	}
	return count
}
