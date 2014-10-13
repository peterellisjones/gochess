package stack

import (
	bd "github.com/peterellisjones/gochess/board"
	mk "github.com/peterellisjones/gochess/make"
	mv "github.com/peterellisjones/gochess/move"
)

// Stack represents a board
// upon which moves can be made and undone
type Stack struct {
	board    *bd.Board
	moveData [127]*mk.MoveData
	depth    int
}

// New creates a new stack
func New(board *bd.Board) *Stack {
	return &Stack{
		board: board,
		depth: 0,
	}
}

// Make makes a move and increments the depth
func (stack *Stack) Make(move mv.Move) {
	stack.moveData[stack.depth] = mk.Make(move, stack.board)
	stack.depth++
}

// UnMake unmakes a move and decrements the depth
func (stack *Stack) UnMake() {
	stack.depth--
	moveData := stack.moveData[stack.depth]
	mk.UnMake(stack.board, moveData)
}

// Depth returns the height of the stack
func (stack *Stack) Depth() int {
	return stack.depth
}
