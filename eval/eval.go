package eval

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

func Board(values *Values, bd *board.Board) int {
	score := 0

	bd.ForEachPieceOfSide(side.White, func(pc piece.Piece, sq square.Square) {
		score += values.PieceValue(pc)
		score += values.PieceSquareValue(pc, sq)
	})

	bd.ForEachPieceOfSide(side.Black, func(pc piece.Piece, sq square.Square) {
		score -= values.PieceValue(pc)
		score -= values.PieceSquareValue(pc, sq)
	})

	return score
}

func Move(values *Values, bd *board.Board, mv move.Move) int {
	if mv.IsCastle() {
		return values.CastleValue(mv)
	}

	pc := bd.At(mv.From())

	score := 0
	score = values.PieceSquareValue(pc, mv.To()) - values.PieceSquareValue(pc, mv.From())

	if mv.IsCapture() {
		if mv.IsEpCapture() {
			score += values.PieceValue(piece.ForSide(piece.Pawn, pc.Side().Other()))
		} else {
			score += values.PieceValue(bd.At(mv.To()))
		}
	}

	if mv.IsPromotion() {
		score += values.PieceValue(mv.PromoteTo()) - values.PieceValue(pc)
	}

	return score
}
