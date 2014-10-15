package make

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

// MoveData records information about the move to make
// undo easier
type MoveData struct {
	Captured piece.Piece
	Move     move.Move
	Extra    board.Extra
}

var rookCastleMoves = [2][2][2]square.Square{
	[2][2]square.Square{
		[2]square.Square{square.H1, square.F1},
		[2]square.Square{square.A1, square.D1},
	},
	[2][2]square.Square{
		[2]square.Square{square.H8, square.F8},
		[2]square.Square{square.A8, square.D8},
	},
}

var kingCastleMoves = [2][2][2]square.Square{
	[2][2]square.Square{
		[2]square.Square{square.E1, square.G1},
		[2]square.Square{square.E1, square.C1},
	},
	[2][2]square.Square{
		[2]square.Square{square.E8, square.G8},
		[2]square.Square{square.E8, square.C8},
	},
}

// Make makes a move on a board
func Make(mv move.Move, bd *board.Board) *MoveData {

	side := bd.SideToMove()
	moveData := MoveData{
		Move:  mv,
		Extra: bd.IrreversibleData(),
	}

	newIrreversibleData := board.Extra{
		SideToMove:     side ^ 1,
		EpSquare:       square.Null,
		FullMoveNumber: bd.FullMoveNumber() + int(side),
		HalfMoveClock:  bd.HalfMoveClock() + 1,
	}

	if mv.IsCastle() {
		newIrreversibleData.CastlingRights = makeCastle(mv, bd)
	} else {
		newIrreversibleData.CastlingRights = updateRights(mv, bd)
		isCapture := mv.IsCapture()

		// reset HMC if pawn move or capture
		if bd.At(mv.From()).Type() == piece.Pawn || isCapture {
			newIrreversibleData.HalfMoveClock = 0
		}

		// set EP square if double pawn push
		if mv.IsDoublePawnPush() {
			epSquare := (mv.From() + mv.To()) >> 1
			newIrreversibleData.EpSquare = epSquare
		}

		// remove and record captured piece if capture
		if isCapture {
			var captureSq square.Square
			if mv.IsEpCapture() {
				captureSq = (mv.To() & 7) | (mv.From() & 56)
			} else {
				captureSq = mv.To()
			}
			captured := bd.At(captureSq)
			bd.Remove(captureSq)
			moveData.Captured = captured
		}

		// move piece
		bd.Move(mv.From(), mv.To())
	}

	bd.SetIrreversibleData(newIrreversibleData)

	return &moveData
}

func makeCastle(mv move.Move, bd *board.Board) castling.Right {
	side := bd.SideToMove()

	castle := mv.CastleType()

	from := rookCastleMoves[side][castle][0]
	to := rookCastleMoves[side][castle][1]

	bd.Move(from, to)

	from = kingCastleMoves[side][castle][0]
	to = kingCastleMoves[side][castle][1]

	bd.Move(from, to)

	rights := [2]castling.Right{
		castling.WhiteKSide | castling.WhiteQSide,
		castling.BlackKSide | castling.BlackQSide,
	}

	return bd.CastlingRights() & (^rights[side])
}

func updateRights(mv move.Move, bd *board.Board) castling.Right {
	original := bd.CastlingRights()
	side := bd.SideToMove()
	newRights := original

	for i := uint(0); i < 4; i++ {
		rook := rookCastleMoves[i>>1][i&1][0]
		if mv.From() == rook || mv.To() == rook {
			right := castling.WhiteKSide << i
			newRights &= ^right
		}
	}

	king := kingCastleMoves[side][0][0]
	if mv.From() == king {
		right := (castling.WhiteKSide | castling.WhiteQSide) << (side << 1)
		newRights &= ^right
	}

	return newRights
}
