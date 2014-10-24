package traverse

import (
	bd "github.com/peterellisjones/gochess/board"
	mv "github.com/peterellisjones/gochess/move"
	gen "github.com/peterellisjones/gochess/movegeneration"
	st "github.com/peterellisjones/gochess/stack"
)

type traverser struct {
	board    *bd.Board
	stack    *st.Stack
	maxDepth int
}

func new(board *bd.Board, maxDepth int) *traverser {
	return &traverser{
		stack:    st.New(board),
		board:    board,
		maxDepth: maxDepth,
	}
}

// Traverse traverses every node up to a certain depth, calling a callback
func Traverse(board *bd.Board, depth int, fn func(int, mv.Move)) {

	if depth <= 0 {
		return
	}

	trav := new(board, depth)
	trav.traverse(1, fn)

	return
}

func (trav *traverser) traverse(depth int, fn func(int, mv.Move)) {
	generator := gen.New(trav.board)

	generator.ForEachMove(trav.board.SideToMove(), func(move mv.Move) bool {

		trav.stack.Make(move)
		if gen.InCheck(trav.board, trav.board.SideToMove().Other()) {
			trav.stack.UnMake()
			return false
		}
		stackDepth := trav.stack.Depth()
		fn(stackDepth, move)

		if depth < trav.maxDepth {
			trav.traverse(depth+1, fn)
		}
		trav.stack.UnMake()
		return false
	})
}
