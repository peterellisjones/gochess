package traverse

import (
	bd "github.com/peterellisjones/gochess/board"
	mv "github.com/peterellisjones/gochess/move"
	gen "github.com/peterellisjones/gochess/movegeneration"
	ml "github.com/peterellisjones/gochess/movelist"
	st "github.com/peterellisjones/gochess/stack"
)

type traverser struct {
	board *bd.Board
	stack *st.Stack
}

func new(board *bd.Board) *traverser {
	return &traverser{
		stack: st.New(board),
		board: board,
	}
}

// Traverse traverses every node up to a certain depth, calling a callback
func Traverse(board *bd.Board, depth int, fn func(int, mv.Move, *bd.Board)) {

	if depth <= 0 {
		return
	}

	trav := new(board)
	trav.traverse(depth, fn)

	return
}

func (trav *traverser) traverse(depth int, fn func(int, mv.Move, *bd.Board)) {
	if depth <= 0 {
		return
	}

	list := ml.New()
	generator := gen.New(trav.board, list)
	generator.AddAllMoves(trav.board.SideToMove())

	//fen := trav.board.Fen()

	list.ForEach(func(move mv.Move) {
		trav.stack.Make(move)
		if gen.InCheck(trav.board, trav.board.SideToMove().Other()) {
			trav.stack.UnMake()
			return
		}
		fn(trav.stack.Depth(), move, trav.board)
		trav.traverse(depth-1, fn)
		trav.stack.UnMake()
	})

	// if trav.board.Fen() != fen {
	// 	panic("FEN CHANGED!")
	// }
}
