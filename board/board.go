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
	board           [64]piece.Piece
	bitboards       [14]bitboard.Bitboard
	pieceLists      [12][10]square.Square
	pieceListsIdxs  [64]int
	pieceListsSizes [12]int
	extra           Extra
}

// Add adds a piece to the board on a given square
func (board *Board) Add(pc piece.Piece, sq square.Square) {
	side := pc.Side()
	board.board[sq] = pc
	board.bitboards[pc] = board.bitboards[pc].Set(sq)
	board.bitboards[side] = board.bitboards[side].Set(sq)

	pcIdx := pc - 2
	pieceListIdx := board.pieceListsSizes[pc-2]
	board.pieceListsIdxs[sq] = pieceListIdx
	board.pieceLists[pcIdx][pieceListIdx] = sq
	board.pieceListsSizes[pcIdx]++

	return
}

func (board *Board) ForEachPieceOfSide(sd side.Side, fn func(piece.Piece, square.Square)) {
	for pc := piece.ForSide(piece.Pawn, sd); pc <= piece.BlackKing; pc += 2 {
		pcIdx := pc - 2
		size := board.pieceListsSizes[pcIdx]
		for i := 0; i < size; i++ {
			sq := board.pieceLists[pcIdx][i]
			fn(pc, sq)
		}
	}
}

func (board *Board) ForEachPieceOfSideWithBreak(sd side.Side, fn func(piece.Piece, square.Square) bool) bool {
	for pc := piece.ForSide(piece.Pawn, sd); pc <= piece.BlackKing; pc += 2 {
		pcIdx := pc - 2
		size := board.pieceListsSizes[pcIdx]
		for i := 0; i < size; i++ {
			sq := board.pieceLists[pcIdx][i]
			if fn(pc, sq) {
				return true
			}
		}
	}
	return false
}

func (board *Board) PieceListIndex(sq square.Square) int {
	return board.pieceListsIdxs[sq]
}

func (board *Board) PieceListSize(pc piece.Piece) int {
	return board.pieceListsSizes[pc-2]
}

func (board *Board) PieceList(pc piece.Piece) [10]square.Square {
	return board.pieceLists[pc-2]
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

	pieceListIdx := board.pieceListsIdxs[from]
	board.pieceLists[pc-2][pieceListIdx] = to
	board.pieceListsIdxs[from] = -1
	board.pieceListsIdxs[to] = pieceListIdx
}

// Remove removes a piece from the board
func (board *Board) Remove(sq square.Square) {
	pc := board.board[sq]
	side := pc.Side()

	board.board[sq] = piece.Empty

	mask := bitboard.Bitboard(1) << sq
	board.bitboards[pc] ^= mask
	board.bitboards[side] ^= mask

	pcIdx := pc - 2
	pieceListIdx := board.pieceListsIdxs[sq]
	pieceListLastIdx := board.pieceListsSizes[pcIdx] - 1
	lastPieceSq := board.pieceLists[pcIdx][pieceListLastIdx]
	board.pieceLists[pcIdx][pieceListIdx] = lastPieceSq
	board.pieceListsIdxs[lastPieceSq] = pieceListIdx

	board.pieceListsSizes[pcIdx]--
}

func (board *Board) EachPieceOfType(pc piece.Piece, fn func(square.Square) bool) bool {
	// need to copy piecelist since may be reordered during make and unmake
	pieceList := board.pieceLists[pc-2]
	size := board.pieceListsSizes[pc-2]
	for i := 0; i < size; i++ {
		sq := pieceList[i]
		if fn(sq) {
			return true
		}
	}
	return false
}

func (board *Board) EachPieceOfTypes(fn func(square.Square) bool, pcs ...piece.Piece) bool {
	for _, pc := range pcs {
		// need to copy piecelist since may be reordered during make and unmake
		pieceList := board.pieceLists[pc-2]
		size := board.pieceListsSizes[pc-2]
		for i := 0; i < size; i++ {
			sq := pieceList[i]
			if fn(sq) {
				return true
			}
		}
	}
	return false
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
