package perft

import (
  "github.com/peterellisjones/gochess/board"
  "github.com/peterellisjones/gochess/movegeneration"
  "github.com/peterellisjones/gochess/move"
    "github.com/peterellisjones/gochess/traverse"
  )

type Result struct{
  Nodes int64
  EpCaptures int64
  Promotions int64
  Castles int64
  Captures int64
  Checks int64
}

func Perft(fen string, depth int) ([]Result, error) {
  bd, err := board.FromFen(fen)
  if err != nil {
    return nil, err
  }

  results := make([]Result, depth)

  traverse.Traverse(bd, depth, func(d int, mv move.Move,  bd *board.Board) {
    results[d-1].Nodes++
    if mv.IsCapture() {
      results[d-1].Captures++
    }
    if mv.IsCastle() {
      results[d-1].Castles++
    }
    if mv.IsEpCapture() {
      results[d-1].EpCaptures++
    }
    if mv.IsPromotion(){
      results[d-1].Promotions++
    }
    if movegeneration.InCheck(bd, bd.SideToMove()) {
      results[d-1].Checks++
    }
  })

  return results, nil
}
