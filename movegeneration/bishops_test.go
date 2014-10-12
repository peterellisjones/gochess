package movegeneration_test

import (
  . "github.com/onsi/ginkgo"
  "github.com/peterellisjones/gochess/board"
  . "github.com/peterellisjones/gochess/movegeneration"
  "github.com/peterellisjones/gochess/movelist"
)

var _ = Describe("AddBishopMoves", func() {
  cases := map[string][]string{
    "8/q6/8/2b5/1P6/8/8/8 b -": []string{
      "c5b6", "c5d6", "c5e7", "c5xb4",
      "c5d4", "c5f8", "c5e3", "c5f2",
      "c5g1",
    },
  }

  ItGeneratesMovesFor(cases, func(bd *board.Board, list *movelist.MoveList) {
    generator := New(bd, list)
    generator.AddBishopMoves(bd.SideToMove())
  })
})
