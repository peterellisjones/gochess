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
func (gen *Generator) ForEachPawnMove(sd side.Side, fn func(mv move.Move)) {
	pawns := gen.board.BBPiece(piece.ForSide(piece.Pawn, sd))
	diff := [2]square.Square{square.Square(56), square.Square(8)}[sd]

	pushes := pawns.CircularRightShift(diff) & gen.board.BBEmpty()
	promotions := pushes & (bitboard.Row1 | bitboard.Row8)
	gen.forEachPawnPromotion(promotions, diff, fn)

	pushes &= ^promotions
	gen.forEachPawnPush(pushes, diff, fn)

	doublePushes := pushes.CircularRightShift(diff) & gen.board.BBEmpty()
	doublePushes &= [2]bitboard.Bitboard{bitboard.Row4, bitboard.Row5}[sd]
	gen.forEachDoublePawnPush(doublePushes, diff+diff, fn)

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
		gen.forEachPawnCapturePromotion(capturePromotions, diff, fn)

		captures &= ^capturePromotions
		gen.forEachPawnCapture(captures, diff, fn)

		epCaptures := targets & (bitboard.Bitboard(1) << gen.board.EpSquare())
		gen.forEachPawnEpCapture(epCaptures, diff, fn)
	}
}

func (gen *Generator) forEachPawnPush(targets bitboard.Bitboard, diff square.Square, fn func(move.Move)) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeMove(from, to)
		fn(move)
	})
}

func (gen *Generator) forEachDoublePawnPush(targets bitboard.Bitboard, diff square.Square, fn func(move.Move)) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeDoublePawnPush(from, to)
		fn(move)
	})
}

func (gen *Generator) forEachPawnCapture(targets bitboard.Bitboard, diff square.Square, fn func(mv move.Move)) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeCapture(from, to)
		fn(move)
	})
}

func (gen *Generator) forEachPawnEpCapture(targets bitboard.Bitboard, diff square.Square, fn func(mv move.Move)) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeEpCapture(from, to)
		fn(move)
	})
}

func (gen *Generator) forEachPawnPromotion(targets bitboard.Bitboard, diff square.Square, fn func(move.Move)) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		fn(move.EncodePromotion(from, to, piece.Queen))
		fn(move.EncodePromotion(from, to, piece.Rook))
		fn(move.EncodePromotion(from, to, piece.Bishop))
		fn(move.EncodePromotion(from, to, piece.Knight))
	})
}

func (gen *Generator) forEachPawnCapturePromotion(targets bitboard.Bitboard, diff square.Square, fn func(mv move.Move)) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		fn(move.EncodeCapturePromotion(from, to, piece.Queen))
		fn(move.EncodeCapturePromotion(from, to, piece.Rook))
		fn(move.EncodeCapturePromotion(from, to, piece.Bishop))
		fn(move.EncodeCapturePromotion(from, to, piece.Knight))
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
