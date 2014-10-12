package movegeneration

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/square"
)

type getAttacks func(bitboard.Bitboard, square.Square) bitboard.Bitboard

func getPositiveRayAttacks(occupied bitboard.Bitboard, dir int, from square.Square) bitboard.Bitboard {
	attacks := RayAttacks[dir][from]
	blocker := attacks & occupied
	if blocker != 0 {
		sq := blocker.BitScanForward()
		attacks ^= RayAttacks[dir][sq]
	}
	return attacks
}

func getNegativeRayAttacks(occupied bitboard.Bitboard, dir int, from square.Square) bitboard.Bitboard {
	attacks := RayAttacks[dir][from]
	blocker := attacks & occupied
	if blocker != 0 {
		sq := blocker.BitScanReverse()
		attacks ^= RayAttacks[dir][sq]
	}
	return attacks
}

func getRookRayAttacks(occupied bitboard.Bitboard, from square.Square) bitboard.Bitboard {
	var targets bitboard.Bitboard
	targets |= getPositiveRayAttacks(occupied, North, from)
	targets |= getPositiveRayAttacks(occupied, East, from)
	targets |= getNegativeRayAttacks(occupied, South, from)
	targets |= getNegativeRayAttacks(occupied, West, from)
	return targets
}

func getBishopRayAttacks(occupied bitboard.Bitboard, from square.Square) bitboard.Bitboard {
	var targets bitboard.Bitboard
	targets |= getPositiveRayAttacks(occupied, NorthWest, from)
	targets |= getPositiveRayAttacks(occupied, NorthEast, from)
	targets |= getNegativeRayAttacks(occupied, SouthEast, from)
	targets |= getNegativeRayAttacks(occupied, SouthWest, from)
	return targets
}

func getQueenRayAttacks(occupied bitboard.Bitboard, from square.Square) bitboard.Bitboard {
	var targets bitboard.Bitboard
	targets |= getRookRayAttacks(occupied, from)
	targets |= getBishopRayAttacks(occupied, from)
	return targets
}
