package movegeneration

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/movelist"
	"github.com/peterellisjones/gochess/side"
)

// Generator generates moves
type Generator struct {
	board *board.Board
	list  *movelist.MoveList
}

// New returns a new move generator
func New(bd *board.Board, list *movelist.MoveList) Generator {
	return Generator{
		board: bd,
		list:  list,
	}
}

// AddAllMoves generates all moves
func (gen *Generator) AddAllMoves(side side.Side) {
	gen.AddCastles(side)
	gen.AddKingMoves(side)
	gen.AddQueenMoves(side)
	gen.AddRookMoves(side)
	gen.AddBishopMoves(side)
	gen.AddKnightMoves(side)
	gen.AddPawnMoves(side)
}
