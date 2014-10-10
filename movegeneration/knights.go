package movegeneration

import (
	"github.com/peterellisjones/gochess/bitboard"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/movelist"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

var knightMoves = [64]bitboard.Bitboard{
	0x0000000000020400, 0x0000000000050800, 0x00000000000A1100, 0x0000000000142200,
	0x0000000000284400, 0x0000000000508800, 0x0000000000A01000, 0x0000000000402000,
	0x0000000002040004, 0x0000000005080008, 0x000000000A110011, 0x0000000014220022,
	0x0000000028440044, 0x0000000050880088, 0x00000000A0100010, 0x0000000040200020,
	0x0000000204000402, 0x0000000508000805, 0x0000000A1100110A, 0x0000001422002214,
	0x0000002844004428, 0x0000005088008850, 0x000000A0100010A0, 0x0000004020002040,
	0x0000020400040200, 0x0000050800080500, 0x00000A1100110A00, 0x0000142200221400,
	0x0000284400442800, 0x0000508800885000, 0x0000A0100010A000, 0x0000402000204000,
	0x0002040004020000, 0x0005080008050000, 0x000A1100110A0000, 0x0014220022140000,
	0x0028440044280000, 0x0050880088500000, 0x00A0100010A00000, 0x0040200020400000,
	0x0204000402000000, 0x0508000805000000, 0x0A1100110A000000, 0x1422002214000000,
	0x2844004428000000, 0x5088008850000000, 0xA0100010A0000000, 0x4020002040000000,
	0x0400040200000000, 0x0800080500000000, 0x1100110A00000000, 0x2200221400000000,
	0x4400442800000000, 0x8800885000000000, 0x100010A000000000, 0x2000204000000000,
	0x0004020000000000, 0x0008050000000000, 0x00110A0000000000, 0x0022140000000000,
	0x0044280000000000, 0x0088500000000000, 0x0010A00000000000, 0x0020400000000000,
}

// GenerateKnightMoves generates knight moves
func GenerateKnightMoves(board *board.Board, list *movelist.MoveList) {
	pc := piece.ForSide(piece.Knight, board.SideToMove())
	addLookupTableMoves(pc, &knightMoves, board, list)
}

func addLookupTableMoves(piece piece.Piece, table *[64]bitboard.Bitboard, board *board.Board, list *movelist.MoveList) {
	side := board.SideToMove()
	enemy := board.BBSide(side.Other())
	empty := board.BBEmpty()

	board.BBPiece(piece).ForEachSetBit(func(from square.Square) {
		targets := table[from]
		captures := targets & enemy
		moves := targets & empty
		addMoves(moves, captures, from, list)
	})
}

func addMoves(moves bitboard.Bitboard, captures bitboard.Bitboard, from square.Square, list *movelist.MoveList) {
	captures.ForEachSetBit(func(to square.Square) {
		list.Add(move.EncodeMove(from, to))
	})

	moves.ForEachSetBit(func(to square.Square) {
		list.Add(move.EncodeMove(from, to))
	})
}
