package movegeneration

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

// AddRookMoves generates rook moves
func (gen *Generator) AddRookMoves(sd side.Side) {
	piece := piece.ForSide(piece.Rook, sd)
	gen.addSliderMoves(piece, getRookRayAttacks)
}

// AddBishopMoves generates bishop moves
func (gen *Generator) AddBishopMoves(sd side.Side) {
	piece := piece.ForSide(piece.Bishop, sd)
	gen.addSliderMoves(piece, getBishopRayAttacks)
}

// AddQueenMoves generates queen moves
func (gen *Generator) AddQueenMoves(sd side.Side) {
	piece := piece.ForSide(piece.Queen, sd)
	gen.addSliderMoves(piece, getQueenRayAttacks)
}

func (gen *Generator) addSliderMoves(piece piece.Piece, getRayAttacks getAttacks) {
	enemy := gen.board.BBSide(piece.Side().Other())
	occupied := gen.board.BBOccupied()

	gen.board.EachPieceOfType(piece, func(from square.Square) {
		targets := getRayAttacks(occupied, from)
		captures := targets & enemy
		nonCaptures := targets & (^occupied)
		gen.addMoves(nonCaptures, captures, from)
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
