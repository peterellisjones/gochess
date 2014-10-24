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

func (search *Search) BestMove() Score {

	best := Score{
		Score: -1000000,
		Move:  move.Null,
	}

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) bool {
		s := search.eval.Move(search.board, mv)

		if s > best.Score {
			best = Score{
				Score: s,
				Move:  mv,
			}
		}
		return false
	})

	return best
}

func (search *Search) Negamax(depth int) (move.Move, int) {
	bestScore := -100000
	bestMove := move.Null

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) bool {
		s := search.eval.Move(search.board, mv)

		if depth > 1 {
			search.stack.Make(mv)
			_, d := search.Negamax(depth - 1)
			s -= d
			search.stack.UnMake()
		}

		if s > bestScore {
			bestScore = s
			bestMove = mv
		}
		return false
	})

	return bestMove, bestScore
}

func (search *Search) AlphaBeta(alpha int, beta int, depth int) (move.Move, int) {

	bestMove := move.Null

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) bool {
		s := search.eval.Move(search.board, mv)

		if depth > 1 {
			search.stack.Make(mv)
			_, d := search.AlphaBeta(-beta, -alpha, depth-1)
			s -= d
			search.stack.UnMake()
		}

		if s >= beta {
			alpha = beta
			bestMove = mv
			return true
		}

		if s > alpha {
			bestMove = mv
			alpha = s
		}

		return false
	})

	return bestMove, alpha
}
