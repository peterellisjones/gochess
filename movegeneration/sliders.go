package movegeneration

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

func (gen *Generator) ForEachRookMove(sd side.Side, fn func(move.Move) bool) bool {
	pc := piece.ForSide(piece.Rook, sd)
	return gen.forEachSliderMove(pc, getRookRayAttacks, fn)
}

func (gen *Generator) ForEachBishopMove(sd side.Side, fn func(move.Move) bool) bool {
	pc := piece.ForSide(piece.Bishop, sd)
	return gen.forEachSliderMove(pc, getBishopRayAttacks, fn)
}

func (gen *Generator) ForEachQueenMove(sd side.Side, fn func(move.Move) bool) bool {
	pc := piece.ForSide(piece.Queen, sd)
	return gen.forEachSliderMove(pc, getQueenRayAttacks, fn)
}

func (gen *Generator) ForEachNonDiagonalMove(sd side.Side, fn func(move.Move) bool) bool {
	queen := piece.ForSide(piece.Queen, sd)
	rook := piece.ForSide(piece.Rook, sd)

	enemy := gen.board.BBSide(sd.Other())
	occupied := gen.board.BBOccupied()

	return gen.board.EachPieceOfTypes(func(from square.Square) bool {
		targets := getRookRayAttacks(occupied, from)
		captures := targets & enemy
		nonCaptures := targets & (^occupied)

		if captures.ForEachSetBitWithBreak(func(to square.Square) bool {
			return fn(move.EncodeCapture(from, to))
		}) {
			return true
		}

		return nonCaptures.ForEachSetBitWithBreak(func(to square.Square) bool {
			return fn(move.EncodeMove(from, to))
		})
	}, queen, rook)
}

func (gen *Generator) ForEachDiagonalMove(sd side.Side, fn func(move.Move) bool) bool {
	queen := piece.ForSide(piece.Queen, sd)
	bishop := piece.ForSide(piece.Bishop, sd)

	enemy := gen.board.BBSide(sd.Other())
	occupied := gen.board.BBOccupied()

	return gen.board.EachPieceOfTypes(func(from square.Square) bool {
		targets := getBishopRayAttacks(occupied, from)
		captures := targets & enemy
		nonCaptures := targets & (^occupied)

		if captures.ForEachSetBitWithBreak(func(to square.Square) bool {
			return fn(move.EncodeCapture(from, to))
		}) {
			return true
		}

		return nonCaptures.ForEachSetBitWithBreak(func(to square.Square) bool {
			return fn(move.EncodeMove(from, to))
		})
	}, queen, bishop)
}

func (gen *Generator) forEachSliderMove(pc piece.Piece, getRayAttacks getAttacks, fn func(move.Move) bool) bool {
	enemy := gen.board.BBSide(pc.Side().Other())
	occupied := gen.board.BBOccupied()

	return gen.board.EachPieceOfType(pc, func(from square.Square) bool {

		targets := getRayAttacks(occupied, from)
		captures := targets & enemy
		nonCaptures := targets & (^occupied)

		if captures.ForEachSetBitWithBreak(func(to square.Square) bool {
			return fn(move.EncodeCapture(from, to))
		}) {
			return true
		}

		return nonCaptures.ForEachSetBitWithBreak(func(to square.Square) bool {
			return fn(move.EncodeMove(from, to))
		})
	})
}

func getSliderAttackedSquares(movers bitboard.Bitboard, occupied bitboard.Bitboard, getRayAttacks getAttacks) bitboard.Bitboard {
	attackedSquares := bitboard.Empty
	movers.ForEachSetBit(func(from square.Square) {
		attackedSquares |= getRayAttacks(occupied, from)
	})
	return attackedSquares
}

// GetRookAttackedSquares returns the set of rook attacks
func GetRookAttackedSquares(bd *board.Board, attacker side.Side) bitboard.Bitboard {
	piece := piece.ForSide(piece.Rook, attacker)
	movers := bd.BBPiece(piece)
	return getSliderAttackedSquares(movers, bd.BBOccupied(), getRookRayAttacks)
}

// GetBishopAttackedSquares returns the set of bishop attacks
func GetBishopAttackedSquares(bd *board.Board, attacker side.Side) bitboard.Bitboard {
	piece := piece.ForSide(piece.Bishop, attacker)
	movers := bd.BBPiece(piece)
	return getSliderAttackedSquares(movers, bd.BBOccupied(), getBishopRayAttacks)
}

// GetQueenAttackedSquares returns the set of queen attacks
func GetQueenAttackedSquares(bd *board.Board, attacker side.Side) bitboard.Bitboard {
	piece := piece.ForSide(piece.Queen, attacker)
	movers := bd.BBPiece(piece)
	return getSliderAttackedSquares(movers, bd.BBOccupied(), getQueenRayAttacks)
}
