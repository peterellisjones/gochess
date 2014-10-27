package search

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/eval"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/stack"
)

const ScoreMin = -1000000
const ScoreMax = 1000000

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
		Score: ScoreMin,
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
	bestScore := ScoreMin
	bestMove := move.Null

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) bool {
		s := search.eval.Move(search.board, mv)

		if depth > 1 && s < 10000 {
			search.stack.Make(mv)
			_, d := search.Negamax(depth - 1)
			s -= d
			search.stack.UnMake()
		}

		//fmt.Printf("%s, %s => %d\n", search.board.SideToMove(), mv.String(), s)
		//
		// if depth == 3 {
		// 	fmt.Printf("BLACK %s => %d\n", mv.String(), s)
		// }

		if s > bestScore {
			bestScore = s
			bestMove = mv
		}
		return false
	})

	return bestMove, bestScore
}

func (search *Search) AlphaBeta(depth int) (move.Move, int) {
	return search.alphaBeta(ScoreMin, ScoreMax, depth)
}

func (search *Search) alphaBeta(alpha int, beta int, depth int) (move.Move, int) {

	bestMove := move.Null
	bestScore := ScoreMin

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) bool {
		s := search.eval.Move(search.board, mv)

		if depth > 1 {
			search.stack.Make(mv)
			s -= search.alphaBetaScore(-beta, -alpha, depth-1)
			search.stack.UnMake()
		}

		if s >= beta {
			bestScore = s
			bestMove = mv
			return true
		}

		if s > alpha {
			alpha = s
			bestScore = s
			bestMove = mv
		}

		return false
	})

	return bestMove, bestScore
}

func (search *Search) alphaBetaScore(alpha int, beta int, depth int) int {

	ret := ScoreMin

	search.gen.ForEachMove(search.board.SideToMove(), func(mv move.Move) bool {
		s := search.eval.Move(search.board, mv)

		if depth > 1 {
			search.stack.Make(mv)
			s -= search.alphaBetaScore(-beta, -alpha, depth-1)
			search.stack.UnMake()
		}

		if s >= beta {
			ret = beta
			return true
		}

		if s > alpha {
			alpha = s
			ret = s
		}

		return false
	})

	return ret
}
