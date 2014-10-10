package bitboard

// Bitboard represents a board using 64 bits
type Bitboard uint64

// represents a full board
const Full Bitboard = 0xFFFFFFFFFFFFFFFF

// represents an Empty board
const Empty Bitboard = 0

// represents boards with row X occupied
const (
	Row1 Bitboard = Bitboard(0xFF) << (iota * 8)
	Row2
	Row3
	Row4
	Row5
	Row6
	Row7
	Row8
)

// represents boards with col X occupied
const (
	FileA Bitboard = Bitboard(0x0101010101010101) << iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// represents boards with square X occupied
const (
	SquareA1 Bitboard = Bitboard(1) << iota
	SquareB1
	SquareC1
	SquareD1
	SquareE1
	SquareF1
	SquareG1
	SquareH1
	SquareA2
	SquareB2
	SquareC2
	SquareD2
	SquareE2
	SquareF2
	SquareG2
	SquareH2
	SquareA3
	SquareB3
	SquareC3
	SquareD3
	SquareE3
	SquareF3
	SquareG3
	SquareH3
	SquareA4
	SquareB4
	SquareC4
	SquareD4
	SquareE4
	SquareF4
	SquareG4
	SquareH4
	SquareA5
	SquareB5
	SquareC5
	SquareD5
	SquareE5
	SquareF5
	SquareG5
	SquareH5
	SquareA6
	SquareB6
	SquareC6
	SquareD6
	SquareE6
	SquareF6
	SquareG6
	SquareH6
	SquareA7
	SquareB7
	SquareC7
	SquareD7
	SquareE7
	SquareF7
	SquareG7
	SquareH7
	SquareA8
	SquareB8
	SquareC8
	SquareD8
	SquareE8
	SquareF8
	SquareG8
	SquareH8
)
