package movelist

import (
	"github.com/peterellisjones/gochess/move"
)

// MoveList represents a list of moves
// in the current implemnetation, may produce an array out o fbounds exception
type MoveList struct {
	moves [128]move.Move
	size  int
}

// New returns an empty MoveList
func New() *MoveList {
	return &MoveList{
		size: 0,
	}
}

// Length returns the length of the list
func (list *MoveList) Length() int {
	return list.size
}

// Clear empties the list
func (list *MoveList) Clear() {
	list.size = 0
}


// Add adds a move to a list
func (list *MoveList) Add(move move.Move) {
	list.moves[list.size] = move
	list.size++
}

// ForEach executes a function on each move in the list
func (list *MoveList) ForEach(fn func(move.Move)) {
	for i := 0; i < list.size; i++ {
		fn(list.moves[i])
	}
}

// Includes returns true if the list contains a given move
func (list *MoveList) Includes(move move.Move) bool {
	for i := 0; i < list.size; i++ {
		if list.moves[i] == move {
			return true
		}
	}
	return false
}

func (list *MoveList) String() string {
	str := ""
	count := 0
	list.ForEach(func(mv move.Move) {
		str += mv.String()
		count++
		if count%8 == 0 {
			str += "\n"
		} else {
			str += "\t"
		}
	})
	return str
}
