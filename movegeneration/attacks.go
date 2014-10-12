package movegeneration

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/side"
)

// AreSquaresAttacked returns true if the target squares are attacked by the specified side
func AreSquaresAttacked(bd *board.Board, targets bitboard.Bitboard, attacker side.Side) bool {
	pawnAttacks := GetPawnAttackedSquares(bd, attacker) & targets
	if pawnAttacks != bitboard.Empty {
		return true
	}

	knightAttacks := GetKnightAttackedSquares(bd, attacker) & targets
	if knightAttacks != bitboard.Empty {
		return true
	}
	kingAttacks := GetKingAttackedSquares(bd, attacker) & targets
	if kingAttacks != bitboard.Empty {
		return true
	}

	rookAttacks := GetRookAttackedSquares(bd, attacker) & targets
	if rookAttacks != bitboard.Empty {
		return true
	}

	bishopAttacks := GetBishopAttackedSquares(bd, attacker) & targets
	if bishopAttacks != bitboard.Empty {
		return true
	}

	queenAttacks := GetQueenAttackedSquares(bd, attacker) & targets
	if queenAttacks != bitboard.Empty {
		return true
	}

	return false
}

// AttackedSquares returns the set of attacked squares
func AttackedSquares(bd *board.Board, attacker side.Side) bitboard.Bitboard {
	attacks := bitboard.Empty
	attacks |= GetPawnAttackedSquares(bd, attacker)
	attacks |= GetKnightAttackedSquares(bd, attacker)
	attacks |= GetKingAttackedSquares(bd, attacker)
	attacks |= GetRookAttackedSquares(bd, attacker)
	attacks |= GetBishopAttackedSquares(bd, attacker)
	attacks |= GetQueenAttackedSquares(bd, attacker)
	return attacks
}
