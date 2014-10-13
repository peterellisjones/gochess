package board

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

// Extra represents the information in a board that cannot be undone
type Extra struct {
	HalfMoveClock  int
	EpSquare       square.Square
	CastlingRights castling.Right
	FullMoveNumber int
	SideToMove     side.Side
}

// Board represents a chess board
type Board struct {
	board     [64]piece.Piece
	bitboards [14]bitboard.Bitboard
	extra     Extra
}

// Add adds a piece to the board on a given square
func (board *Board) Add(pc piece.Piece, sq square.Square) {
	side := pc.Side()
	board.board[sq] = pc
	board.bitboards[pc] = board.bitboards[pc].Set(sq)
	board.bitboards[side] = board.bitboards[side].Set(sq)
}

// Move moves a piece
func (board *Board) Move(from square.Square, to square.Square) {
	pc := board.board[from]
	board.board[from] = piece.Empty
	board.board[to] = pc

	mask := (bitboard.Bitboard(1) << from) | (bitboard.Bitboard(1) << to)
	sd := pc.Side()
	board.bitboards[sd] ^= mask
	board.bitboards[pc] ^= mask
}

// Remove removes a piece from the board
func (board *Board) Remove(sq square.Square) {
	pc := board.board[sq]
	side := pc.Side()

	board.board[sq] = piece.Empty

	mask := bitboard.Bitboard(1) << sq
	board.bitboards[pc] ^= mask
	board.bitboards[side] ^= mask
}

// At returns the piece on a given square (if any)
func (board *Board) At(sq square.Square) piece.Piece {
	return board.board[sq]
}

// BBPiece returns the occupation bitboard for a given piece
func (board *Board) BBPiece(i piece.Piece) bitboard.Bitboard {
	return board.bitboards[i]
}

// BBEmpty returns the bitboard of non occupied squares
func (board *Board) BBEmpty() bitboard.Bitboard {
	return ^(board.bitboards[0] | board.bitboards[1])
}

// BBOccupied returns the bitboard of occupied squares
func (board *Board) BBOccupied() bitboard.Bitboard {
	return board.bitboards[0] | board.bitboards[1]
}

// BBSide returns the occupation bitboard for a given side
func (board *Board) BBSide(i side.Side) bitboard.Bitboard {
	return board.bitboards[i]
}

// SideToMove returns the next side to move
func (board *Board) SideToMove() side.Side {
	return board.extra.SideToMove
}

// EpSquare returns the en-passant square, if any
func (board *Board) EpSquare() square.Square {
	return board.extra.EpSquare
}

// CastlingRights returns a bitmask of possible castling rights for the current position
func (board *Board) CastlingRights() castling.Right {
	return board.extra.CastlingRights
}

// HalfMoveClock returns the half move clock
func (board *Board) HalfMoveClock() int {
	return board.extra.HalfMoveClock
}

// ResetHalfMoveClock sets the half move lcock to zero
func (board *Board) ResetHalfMoveClock() {
	board.extra.HalfMoveClock = 0
}

// FullMoveNumber returns the full move number
func (board *Board) FullMoveNumber() int {
	return board.extra.FullMoveNumber
}

// IrreversibleData returns the set of information that cannot be undone on each move
func (board *Board) IrreversibleData() Extra {
	return board.extra
}

// SetIrreversibleData sets the irreversible data
func (board *Board) SetIrreversibleData(extra Extra) {
	board.extra = extra
}

// EmptyBoard returns an empty board
func EmptyBoard() *Board {
	board := Board{
		extra: Extra{
			SideToMove:     side.White,
			FullMoveNumber: 1,
			EpSquare:       square.Null,
			CastlingRights: castling.NoRights,
			HalfMoveClock:  0,
		},
	}

	return &board
}
