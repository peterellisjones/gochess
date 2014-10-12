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
func (gen *Generator) AddPawnMoves(sd side.Side) {
	pawns := gen.board.BBPiece(piece.ForSide(piece.Pawn, sd))
	diff := [2]square.Square{square.Square(56), square.Square(8)}[sd]

	pushes := pawns.CircularRightShift(diff) & gen.board.BBEmpty()
	promotions := pushes & (bitboard.Row1 | bitboard.Row8)
	gen.addPawnPromotions(promotions, diff)

	pushes &= ^promotions
	gen.addPawnPushes(pushes, diff)

	doublePushes := pushes.CircularRightShift(diff) & gen.board.BBEmpty()
	doublePushes &= [2]bitboard.Bitboard{bitboard.Row4, bitboard.Row5}[sd]
	gen.addPawnDoublePushes(doublePushes, diff+diff)

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
		gen.addPawnCapturePromotions(capturePromotions, diff)

		captures &= ^capturePromotions
		gen.addPawnCaptures(captures, diff)

		epCaptures := targets & (bitboard.Bitboard(1) << gen.board.EpSquare())
		gen.addPawnEpCaptures(epCaptures, diff)
	}
}

func (gen *Generator) addPawnPushes(targets bitboard.Bitboard, diff square.Square) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeMove(from, to)
		gen.list.Add(move)
	})
}

func (gen *Generator) addPawnDoublePushes(targets bitboard.Bitboard, diff square.Square) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeDoublePawnPush(from, to)
		gen.list.Add(move)
	})
}

func (gen *Generator) addPawnCaptures(targets bitboard.Bitboard, diff square.Square) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeCapture(from, to)
		gen.list.Add(move)
	})
}

func (gen *Generator) addPawnEpCaptures(targets bitboard.Bitboard, diff square.Square) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		move := move.EncodeEpCapture(from, to)
		gen.list.Add(move)
	})
}

func (gen *Generator) addPawnPromotions(targets bitboard.Bitboard, diff square.Square) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		gen.list.Add(move.EncodePromotion(from, to, piece.Queen))
		gen.list.Add(move.EncodePromotion(from, to, piece.Rook))
		gen.list.Add(move.EncodePromotion(from, to, piece.Bishop))
		gen.list.Add(move.EncodePromotion(from, to, piece.Knight))
	})
}

func (gen *Generator) addPawnCapturePromotions(targets bitboard.Bitboard, diff square.Square) {
	targets.ForEachSetBit(func(to square.Square) {
		from := to.CircularTranslate(diff)
		gen.list.Add(move.EncodeCapturePromotion(from, to, piece.Queen))
		gen.list.Add(move.EncodeCapturePromotion(from, to, piece.Rook))
		gen.list.Add(move.EncodeCapturePromotion(from, to, piece.Bishop))
		gen.list.Add(move.EncodeCapturePromotion(from, to, piece.Knight))
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
