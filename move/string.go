package move

import (
	"strings"
)

func (move Move) String() string {
	strs := []string{move.From().String(), move.To().String()}
	if move.IsCastle() {
		if move == QueenSideCastle {
			return "O-O-O"
		} else if move == KingSideCastle {
			return "O-O"
		}
	} else if move.IsCapture() {
		return strings.Join(strs, "x")
	}
	return strings.Join(strs, "")
}
