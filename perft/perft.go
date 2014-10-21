package perft

import (
	"fmt"
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

type Results []Result

type MovesResults map[move.Move]Results

func PerftMoves(fen string, depth int) (MovesResults, error) {
	bd, err := board.FromFen(fen)
	if err != nil {
		return nil, err
	}

	ret := MovesResults{}

	traverse.Traverse(bd, 1, func(d int, mv move.Move, bd *board.Board) {
		var results []Result
		results, err = Perft(bd.Fen(), depth-1)
		if err != nil {
			return
		}

		// fmt.Println("***************************")
		// fmt.Println(mv)
		// fmt.Println(bd.Fen())

		ret[mv] = results
	})

	return ret, err
}

func Perft(fen string, depth int) ([]Result, error) {
	bd, err := board.FromFen(fen)
	if err != nil {
		return nil, err
	}

	results := make([]Result, depth)

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

func (results Results) Nodes() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Nodes
	}
	return count
}

func (results Results) Captures() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Captures
	}
	return count
}

func (results Results) EpCaptures() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.EpCaptures
	}
	return count
}

func (results Results) Castles() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Castles
	}
	return count
}

func (results Results) Promotions() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Promotions
	}
	return count
}

func (results Results) Checks() int64 {
	count := int64(0)
	for _, result := range results {
		count += result.Checks
	}
	return count
}

func (results MovesResults) String() string {
	str := "Move\tNodes\nCaptures\nEpCaptures\nCastles\nPromotions\nChecks"
	for mv, result := range results {
		str += fmt.Sprintf(
			"%s\t%d\t%d\n",
			mv.String(),
			result.Nodes(),
			result.Captures(),
			result.EpCaptures(),
			result.Castles(),
			result.Promotions(),
			result.Checks(),
		)
	}
	return str
}
