package castling

import (
	"errors"

	"github.com/peterellisjones/gochess/side"
)

// Right represents a bitmask of possible castling rights
type Right uint8

// Possible castling rights
const (
	NoRights   Right = Right(0)
	WhiteKSide Right = Right(1)
	WhiteQSide Right = Right(2)
	whiteXSide Right = Right(1 + 2)
	BlackKSide Right = Right(4)
	BlackQSide Right = Right(8)
	blackXSide Right = Right(4 + 8)
)

// BlackCanCastle returns true if black can castle
func (right Right) BlackCanCastle() bool {
	return right&blackXSide != NoRights
}

// WhiteCanCastle returns true if white can castle
func (right Right) WhiteCanCastle() bool {
	return right&whiteXSide != NoRights
}

// RightsForSide returns the set of rights for a specific side
func (right Right) RightsForSide(side side.Side) Right {
	sideMask := []Right{whiteXSide, blackXSide}[side]
	return sideMask & right
}

// HasRight returns true if a given right exists in the set of rights
func (right Right) HasRight(rightB Right) bool {
	return right&rightB != NoRights
}

// Rights returns an array of individual rights
func (right Right) Rights() []Right {
	rights := []Right{}
	for i := Right(0); i < 4; i++ {
		r := Right(1) << i
		if right&r != NoRights {
			rights = append(rights, r)
		}
	}
	return rights
}

var rightNames = map[Right]string{
	WhiteKSide: "K",
	WhiteQSide: "Q",
	BlackKSide: "k",
	BlackQSide: "q",
}

func (right Right) String() string {
	ret := ""
	for _, r := range right.Rights() {
		ret = ret + rightNames[r]
	}
	return ret
}

// ForEach executes a function for each right in the set of rights
func (right Right) ForEach(fn func(uint) bool) bool {
	for i := uint(0); i < 4; i++ {
		r := right & (Right(1) << i)
		if r != NoRights {
			if fn(i) {
				return true
			}
		}
	}
	return false
}

// Parse parses a string representation of a set of castling rights
func Parse(str string) (Right, error) {
	rights := NoRights

	charToRights := map[byte]Right{
		'K': WhiteKSide,
		'Q': WhiteQSide,
		'k': BlackKSide,
		'q': BlackQSide,
	}

	for i := 0; i < len(str); i++ {
		right, ok := charToRights[str[i]]
		if !ok {
			return rights, errors.New("Castling right not recognized")
		}
		rights |= right
	}

	return rights, nil
}
