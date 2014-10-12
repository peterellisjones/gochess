package movegeneration


import (
  "github.com/peterellisjones/gochess/board"
  "github.com/peterellisjones/gochess/movelist"
)

type Generator struct {
  board *board.Board
  list *movelist.MoveList
}

func New(bd *board.Board, list *movelist.MoveList) Generator {
  return Generator{
    board: bd,
    list: list,
  }
}
