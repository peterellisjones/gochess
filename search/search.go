package search

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/eval"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/stack"
)

type Search struct {
	stack *stack.Stack
	eval  *eval.Eval
	gen   *movegeneration.Generator
	board *board.Board
}

func New(bd *board.Board, ev *eval.Eval) *Search {
	return &Search{
		stack: stack.New(bd),
		eval:  ev,
		gen:   movegeneration.New(bd),
		board: bd,
	}
}

type Score struct {
	Move  move.Move
	Score int
}

func (search *Search) BestMove(depth int) Score {

	best := Score{
		Score: -1000000,
		Move:  move.Null,
	}

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) {
		s := search.eval.Move(search.board, mv)
		if s > best.Score {
			best = Score{
				Score: s,
				Move:  mv,
			}
		}
	})

	return best
}
