package perft

import (
	"errors"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/traverse"
)

type Result struct {
	Nodes      int64
	EpCaptures int64
	Promotions int64
	Castles    int64
	Captures   int64
	Checks     int64
}

type SimpleResult struct {
	Nodes int64
}

type Results []Result

func PerftMoves(fen string, depth int) (map[string]int64, error) {
	bd, err := board.FromFen(fen)
	if err != nil {
		return nil, err
	}

	ret := map[string]int64{}

	if depth == 0 {
		return ret, errors.New("Depth cannot be 0")
	}

	traverse.Traverse(bd, 1, func(d int, mv move.Move, bd *board.Board) {
		var results Results

		results, err = Perft(bd.Fen(), depth-1)
		if err != nil {
			return
		}

		ret[mv.String()] = results.LeafNodes()
	})

	return ret, err
}

func (results Results) LeafNodes() int64 {
	return results[len(results)-1].Nodes
}

func Perft(fen string, depth int) (Results, error) {
	bd, err := board.FromFen(fen)
	if err != nil {
		return nil, err
	}

	results := Results(make([]Result, depth))

	traverse.Traverse(bd, depth, func(d int, mv move.Move, bd *board.Board) {
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
		if mv.IsPromotion() {
			results[d-1].Promotions++
		}
		if movegeneration.InCheck(bd, bd.SideToMove()) {
			results[d-1].Checks++
		}
	})

	return results, nil
}

func (results Results) nodes() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Nodes
	}
	return count
}

func (results Results) captures() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Captures
	}
	return count
}

func (results Results) epCaptures() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.EpCaptures
	}
	return count
}

func (results Results) castles() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Castles
	}
	return count
}

func (results Results) promotions() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Promotions
	}
	return count
}

func (results Results) checks() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Checks
	}
	return count
}

// func (results MovesResults) String() string {
// 	str := fmt.Sprintf(
// 		"%10s %10s %10s %10s %10s %10s %10s\n",
// 		"Move",
// 		"Nodes",
// 		"Captures",
// 		"EpCaptures",
// 		"Castles",
// 		"Promotions",
// 		"Checks",
// 	)
// 	for mv, result := range results {
// 		str += fmt.Sprintf(
// 			"%10s %10d %10d %10d %10d %10d %10d\n",
// 			mv.String(),
// 			result.nodes(),
// 			result.captures(),
// 			result.epCaptures(),
// 			result.castles(),
// 			result.promotions(),
// 			result.checks(),
// 		)
// 	}
// 	return str
// }
