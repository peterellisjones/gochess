package board

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

type IrreversibleData struct {
	halfMoveClock  int
	epSquare       square.Square
	castlingRights castling.CastlingRight
	captured       piece.Piece
}

type Board struct {
	board          [64]piece.Piece
	bitboards      [14]bitboard.Bitboard
	irrev          IrreversibleData
	fullMoveNumber int
	sideToMove     side.Side
}

func (board *Board) Add(pc piece.Piece, sq square.Square) {
	board.board[sq] = pc
	board.bitboards[pc] = board.bitboards[pc].Set(sq)
	board.bitboards[pc.Side()] = board.bitboards[pc.Side()].Set(sq)
}

func (board *Board) At(sq square.Square) piece.Piece {
	return board.board[sq]
}

func (board *Board) BBPiece(i piece.Piece) bitboard.Bitboard {
	return board.bitboards[i]
}

func (board *Board) BBSide(i side.Side) bitboard.Bitboard {
	return board.bitboards[i]
}

func (board *Board) SideToMove() side.Side {
	return board.sideToMove
}

func (board *Board) EpSquare() square.Square {
	return board.irrev.epSquare
}

func (board *Board) CastlingRights() castling.CastlingRight {
	return board.irrev.castlingRights
}

func (board *Board) HalfMoveClock() int {
	return board.irrev.halfMoveClock
}

func (board *Board) FullMoveNumber() int {
	return board.fullMoveNumber
}

func (board *Board) Captured() piece.Piece {
	return board.irrev.captured
}

func (board *Board) IrreversibleData() IrreversibleData {
	return board.irrev
}

func EmptyBoard() *Board {
	board := Board{
		fullMoveNumber: 1,
		sideToMove:     side.WHITE,
		irrev: IrreversibleData{
			epSquare:       square.NULL,
			castlingRights: castling.NO_RIGHTS,
			halfMoveClock:  0,
			captured:       piece.EMPTY,
		},
	}

	return &board
}
