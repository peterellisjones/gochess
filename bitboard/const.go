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

// represents boards with  X occupied
const (
	A1 Bitboard = Bitboard(1) << iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)
