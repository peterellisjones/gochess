package validate

import (
	"errors"
	"fmt"

	"github.com/peterellisjones/gochess/bitboard"
	bd "github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

// Board returns an error if the board is not valid
func Board(board *bd.Board) error {
	for sq := square.Square(0); sq < square.Square(64); sq++ {
		pc := board.At(sq)

		if pc == piece.Empty {
			for sd := side.White; sd <= side.Black; sd++ {
				if board.BBSide(sd).IsSet(sq) {
					return errors.New("Expected bitboard not to be set (should be empty)")
				}
			}

			for p := piece.WhitePawn; pc <= piece.BlackKing; pc++ {
				if board.BBPiece(p).IsSet(sq) {
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

			for p := piece.WhitePawn; p <= piece.BlackKing; p++ {
				if p != pc && board.BBPiece(p).IsSet(sq) {
					return errors.New("Expected bitboard not to be set")
				}
			}
		}
	}

	// cant be more than 32 pieces
	pieceCount := 0
	for sq := square.Square(0); sq < square.Square(64); sq++ {
		pc := board.At(sq)
		if pc != piece.Empty {
			pieceCount++
		}
	}

	if pieceCount > 32 {
		return errors.New("Can't have more than 16 pieces")
	}

	// must have one king per side
	if board.BBPiece(piece.ForSide(piece.King, side.White)).BitCount() != 1 {
		return errors.New("must have 1 white king")
	}

	if board.BBPiece(piece.ForSide(piece.King, side.Black)).BitCount() != 1 {
		return errors.New("must have 1 black king")
	}

	// can only have 8 pawns per side
	for sd := side.White; sd <= side.Black; sd++ {
		pieceCount = 0
		pawn := piece.ForSide(piece.Pawn, sd)
		for sq := square.Square(0); sq < square.Square(64); sq++ {
			pc := board.At(sq)
			if pc == pawn {
				pieceCount++

			}
		}
		if pieceCount > 8 {
			return fmt.Errorf("%s can't have more than 8 pawns", sd)
		}
	}

	// kings can't be adjacent
	blackKingAttacks := movegeneration.GetKingAttackedSquares(board, side.Black)
	if blackKingAttacks&board.BBPiece(piece.WhiteKing) != bitboard.Empty {
		return errors.New("black king and white king can't be adjacent")
	}

	// king cant be in check if not side to move
	if movegeneration.InCheck(board, board.SideToMove().Other()) {
		return errors.New("king can't be in check if not the moving side")
	}

	// cant have pawns in first or last rows
	endRows := bitboard.Row1 | bitboard.Row8
	if board.BBPiece(piece.WhitePawn)&endRows != bitboard.Empty {
		return errors.New("can't have white pawns on first or last rows")
	}

	if board.BBPiece(piece.BlackPawn)&endRows != bitboard.Empty {
		return errors.New("can't have black pawns on first or last rows")
	}

	// en-passant square must have an accompanying pawn
	if board.EpSquare().Row() == square.Row3 {
		epSquare := board.EpSquare()
		if board.At(epSquare+8) != piece.WhitePawn {
			return errors.New("en-passant square must have a pawn")
		}
	}

	if board.EpSquare().Row() == square.Row6 {
		epSquare := board.EpSquare()
		if board.At(epSquare-8) != piece.BlackPawn {
			return errors.New("en-passant square must have a pawn")
		}
	}

	// en-passant square resets half move clock
	if board.EpSquare() != square.Null {
		if board.HalfMoveClock() != 0 {
			return errors.New("half move clock must be 0 when there is an en-passant square")
		}
	}

	// half move clock cant be more than twice  of full move clock
	if board.HalfMoveClock() > (board.FullMoveNumber()-1)*2+1 {
		return errors.New("half move clock too high")
	}

	// full move number must be >= 1
	if board.FullMoveNumber() < 1 {
		return errors.New("full move number must be greater than zero")
	}

	// can only castle if rook and king not moved
	if board.CastlingRights().WhiteCanCastle() {
		if board.BBPiece(piece.WhiteKing)&bitboard.E1 == bitboard.Empty {
			return errors.New("cant castle if king moved")
		}
		if board.CastlingRights().HasRight(castling.WhiteKSide) {
			if board.BBPiece(piece.WhiteRook)&bitboard.H1 == bitboard.Empty {
				return errors.New("cant castle if rook moved")
			}
		}
		if board.CastlingRights().HasRight(castling.WhiteQSide) {
			if board.BBPiece(piece.WhiteRook)&bitboard.A1 == bitboard.Empty {
				return errors.New("cant castle if rook moved")
			}
		}
	}

	if board.CastlingRights().BlackCanCastle() {
		if board.BBPiece(piece.BlackKing)&bitboard.E8 == bitboard.Empty {
			return errors.New("cant castle if king moved")
		}
		if board.CastlingRights().HasRight(castling.BlackKSide) {
			if board.BBPiece(piece.BlackRook)&bitboard.H8 == bitboard.Empty {
				return errors.New("cant castle if rook moved")
			}
		}

		if board.CastlingRights().HasRight(castling.BlackQSide) {
			if board.BBPiece(piece.BlackRook)&bitboard.A8 == bitboard.Empty {
				return errors.New("cant castle if rook moved")
			}
		}
	}

	return nil
}
