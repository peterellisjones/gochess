package movegeneration

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

// GeneratePawnMoves generates pawn moves
func (gen *Generator) ForEachPawnMove(sd side.Side, fn func(mv move.Move) bool) bool {
	pawns := gen.board.BBPiece(piece.ForSide(piece.Pawn, sd))
	diff := [2]square.Square{square.Square(56), square.Square(8)}[sd]

	pushes := pawns.CircularRightShift(diff) & gen.board.BBEmpty()
	promotions := pushes & (bitboard.Row1 | bitboard.Row8)
	if gen.forEachPawnPromotion(promotions, diff, fn) {
		return true
	}

	pushes &= ^promotions
	if gen.forEachPawnPush(pushes, diff, fn) {
		return true
	}

	doublePushes := pushes.CircularRightShift(diff) & gen.board.BBEmpty()
	doublePushes &= [2]bitboard.Bitboard{bitboard.Row4, bitboard.Row5}[sd]
	if gen.forEachDoublePawnPush(doublePushes, diff+diff, fn) {
		return true
	}

	enemy := gen.board.BBSide(sd.Other())
	for i := 0; i < 2; i++ {
		diff := [2][2]square.Square{
			[2]square.Square{
				square.Square(64 - 7),
				square.Square(9),
			},
			[2]square.Square{
				square.Square(64 - 9),
				square.Square(7),
			},
		}[i][sd]

		targets := pawns.CircularRightShift(diff)
		targets &= [2]bitboard.Bitboard{^bitboard.FileH, ^bitboard.FileA}[i]

		captures := targets & enemy
		capturePromotions := captures & (bitboard.Row1 | bitboard.Row8)
		if gen.forEachPawnCapturePromotion(capturePromotions, diff, fn) {
			return true
		}

		captures &= ^capturePromotions
		if gen.forEachPawnCapture(captures, diff, fn) {
			return true
		}

		epCaptures := targets & (bitboard.Bitboard(1) << gen.board.EpSquare())
		if gen.forEachPawnEpCapture(epCaptures, diff, fn) {
			return true
		}
	}
	return false
}

func (gen *Generator) forEachPawnPush(targets bitboard.Bitboard, diff square.Square, fn func(move.Move) bool) bool {
	return targets.ForEachSetBitWithBreak(func(to square.Square) bool {
		from := to.CircularTranslate(diff)
		move := move.EncodeMove(from, to)
		return fn(move)
	})
}

func (gen *Generator) forEachDoublePawnPush(targets bitboard.Bitboard, diff square.Square, fn func(move.Move) bool) bool {
	return targets.ForEachSetBitWithBreak(func(to square.Square) bool {
		from := to.CircularTranslate(diff)
		move := move.EncodeDoublePawnPush(from, to)
		return fn(move)
	})
}

func (gen *Generator) forEachPawnCapture(targets bitboard.Bitboard, diff square.Square, fn func(mv move.Move) bool) bool {
	return targets.ForEachSetBitWithBreak(func(to square.Square) bool {
		from := to.CircularTranslate(diff)
		move := move.EncodeCapture(from, to)
		return fn(move)
	})
}

func (gen *Generator) forEachPawnEpCapture(targets bitboard.Bitboard, diff square.Square, fn func(mv move.Move) bool) bool {
	return targets.ForEachSetBitWithBreak(func(to square.Square) bool {
		from := to.CircularTranslate(diff)
		move := move.EncodeEpCapture(from, to)
		return fn(move)
	})
}

func (gen *Generator) forEachPawnPromotion(targets bitboard.Bitboard, diff square.Square, fn func(move.Move) bool) bool {
	return targets.ForEachSetBitWithBreak(func(to square.Square) bool {
		from := to.CircularTranslate(diff)
		if fn(move.EncodePromotion(from, to, piece.Queen)) {
			return true
		}
		if fn(move.EncodePromotion(from, to, piece.Rook)) {
			return true
		}
		if fn(move.EncodePromotion(from, to, piece.Bishop)) {
			return true
		}
		if fn(move.EncodePromotion(from, to, piece.Knight)) {
			return true
		}
		return false
	})
}

func (gen *Generator) forEachPawnCapturePromotion(targets bitboard.Bitboard, diff square.Square, fn func(mv move.Move) bool) bool {
	return targets.ForEachSetBitWithBreak(func(to square.Square) bool {
		from := to.CircularTranslate(diff)
		if fn(move.EncodeCapturePromotion(from, to, piece.Queen)) {
			return true
		}
		if fn(move.EncodeCapturePromotion(from, to, piece.Rook)) {
			return true
		}
		if fn(move.EncodeCapturePromotion(from, to, piece.Bishop)) {
			return true
		}
		if fn(move.EncodeCapturePromotion(from, to, piece.Knight)) {
			return true
		}
		return false
	})
}

// GetPawnAttackedSquares returns the set of pawn attacks
func GetPawnAttackedSquares(bd *board.Board, attacker side.Side) bitboard.Bitboard {
	attackedSquares := bitboard.Empty

	piece := piece.ForSide(piece.Pawn, attacker)
	pawns := bd.BBPiece(piece)
	for i := 0; i < 2; i++ {
		diff := [2][2]square.Square{
			[2]square.Square{
				square.Square(64 - 7),
				square.Square(9),
			},
			[2]square.Square{
				square.Square(64 - 9),
				square.Square(7),
			},
		}[i][attacker]

		targets := pawns.CircularRightShift(diff)
		targets &= [2]bitboard.Bitboard{^bitboard.FileH, ^bitboard.FileA}[i]

		attackedSquares |= targets
	}

	return attackedSquares
}
