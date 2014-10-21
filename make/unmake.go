package make

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
)

// UnMake unmakes a move on a board
func UnMake(bd *board.Board, data *MoveData) {

	mv := data.Move

	if mv.IsCastle() {
		unmakeCastle(mv, bd)
	} else {

		// set promotion piece if promotion
		if mv.IsPromotion() {
			bd.Remove(mv.To())
			bd.Add(data.Mover, mv.From())
		} else {
			// otherwise move piece
			bd.Move(mv.To(), mv.From())
		}

		if mv.IsCapture() {
			captureSq := mv.To()
			if mv.IsEpCapture() {
				captureSq = (mv.To() & 7) | (mv.From() & 56)
			}
			bd.Add(data.Captured, captureSq)
		}
	}

	bd.SetIrreversibleData(data.Extra)
}

func unmakeCastle(mv move.Move, bd *board.Board) {
	side := bd.SideToMove() ^ 1

	castle := mv.CastleType()

	to := rookCastleMoves[side][castle][0]
	from := rookCastleMoves[side][castle][1]

	bd.Move(from, to)

	to = kingCastleMoves[side][castle][0]
	from = kingCastleMoves[side][castle][1]

	bd.Move(from, to)
}
