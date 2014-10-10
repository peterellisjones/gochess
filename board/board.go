package board

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

// IrreversibleData represents the information in a board that cannot be undone
type IrreversibleData struct {
	halfMoveClock  int
	epSquare       square.Square
	castlingRights castling.CastlingRight
	captured       piece.Piece
}

// Board represents a chess board
type Board struct {
	board          [64]piece.Piece
	bitboards      [14]bitboard.Bitboard
	irrev          IrreversibleData
	fullMoveNumber int
	sideToMove     side.Side
}

// Add adds a piece to the board on a given square
func (board *Board) Add(pc piece.Piece, sq square.Square) {
	board.board[sq] = pc
	board.bitboards[pc] = board.bitboards[pc].Set(sq)
	board.bitboards[pc.Side()] = board.bitboards[pc.Side()].Set(sq)
}

// At returns the piece on a given square (if any)
func (board *Board) At(sq square.Square) piece.Piece {
	return board.board[sq]
}

// BBPiece returns the occupation bitboard for a given piece
func (board *Board) BBPiece(i piece.Piece) bitboard.Bitboard {
	return board.bitboards[i]
}

// BBSide returns the occupation bitboard for a given side
func (board *Board) BBSide(i side.Side) bitboard.Bitboard {
	return board.bitboards[i]
}

// SideToMove returns the next side to move
func (board *Board) SideToMove() side.Side {
	return board.sideToMove
}

// EpSquare returns the en-passant square, if any
func (board *Board) EpSquare() square.Square {
	return board.irrev.epSquare
}

// CastlingRights returns a bitmask of possible castling rights for the current position
func (board *Board) CastlingRights() castling.CastlingRight {
	return board.irrev.castlingRights
}

// HalfMoveClock returns the half move clock
func (board *Board) HalfMoveClock() int {
	return board.irrev.halfMoveClock
}

// FullMoveNumber returns the full move number
func (board *Board) FullMoveNumber() int {
	return board.fullMoveNumber
}

// Captured returns the piece captured on the previous move, if any
func (board *Board) Captured() piece.Piece {
	return board.irrev.captured
}

// IrreversibleData returns the set of information that cannot be undone on each move
func (board *Board) IrreversibleData() IrreversibleData {
	return board.irrev
}

// EmptyBoard returns
func EmptyBoard() *Board {
	board := Board{
		fullMoveNumber: 1,
		sideToMove:     side.White,
		irrev: IrreversibleData{
			epSquare:       square.Null,
			castlingRights: castling.NoRights,
			halfMoveClock:  0,
			captured:       piece.Empty,
		},
	}

	return &board
}
