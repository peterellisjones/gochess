package board

import (
	"errors"
	"fmt"
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
	side := pc.Side()
	board.board[sq] = pc
	board.bitboards[pc] = board.bitboards[pc].Set(sq)
	board.bitboards[side] = board.bitboards[side].Set(sq)
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

	for i := 0; i < 14; i++ {
		board.bitboards[i] = bitboard.Empty
	}

	for i := 0; i < 64; i++ {
		board.board[i] = piece.Empty
	}

	return &board
}

// Validate returns an error if the board is not valid
func (board *Board) Validate() error {
	for sq := square.Square(0); sq < square.Square(64); sq++ {
		pc := board.board[sq]

		if pc == piece.Empty {
			for i := 0; i < len(board.bitboards); i++ {
				if board.bitboards[i].IsSet(sq) {
					fmt.Println(sq)
					fmt.Println(pc)
					return errors.New("Expected bitboard not to be set (should be empty)")
				}
			}
		} else {
			side := pc.Side()
			if board.BBSide(side.Other()).IsSet(sq) {
				return errors.New("Bitboard was set for opposing side")
			}

			if !board.BBSide(side).IsSet(sq) {
				return errors.New("Bitboard was not set for side")
			}

			if !board.BBPiece(pc).IsSet(sq) {
				return errors.New("Bitboard was not set for piece")
			}

			for i := 2; i < len(board.bitboards); i++ {
				if i != int(pc) && board.bitboards[i].IsSet(sq) {
					return errors.New("Expected bitboard not to be set")
				}
			}
		}
	}
	return nil
}
