package eval

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

func (eval *Eval) Board(bd *board.Board) int {
	score := 0

	bd.ForEachPieceOfSide(side.White, func(pc piece.Piece, sq square.Square) {
		score += eval.PieceValue(pc)
		score += eval.PieceSquareValue(pc, sq)
	})

	bd.ForEachPieceOfSide(side.Black, func(pc piece.Piece, sq square.Square) {
		score -= eval.PieceValue(pc)
		score -= eval.PieceSquareValue(pc, sq)
	})

	return score
}

func (eval *Eval) Move(bd *board.Board, mv move.Move) int {
	if mv.IsCastle() {
		return eval.CastleValue(mv)
	}

	pc := bd.At(mv.From())

	score := 0
	score = eval.PieceSquareValue(pc, mv.To()) - eval.PieceSquareValue(pc, mv.From())

	if mv.IsCapture() {
		if mv.IsEpCapture() {
			score += eval.PieceValue(piece.ForSide(piece.Pawn, pc.Side().Other()))
		} else {
			score += eval.PieceValue(bd.At(mv.To()))
		}
	}

	if mv.IsPromotion() {
		score += eval.PieceValue(mv.PromoteTo()) - eval.PieceValue(pc)
	}

	return score
}
