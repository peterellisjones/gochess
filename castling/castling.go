package castling

import (
	"errors"
	"github.com/peterellisjones/gochess/side"
)

// CastlingRight represents a bitmask of possible castling rights
type CastlingRight uint8

// Possible castling rights
const (
	NoRights   CastlingRight = CastlingRight(0)
	BlackQSide CastlingRight = CastlingRight(1)
	BlackKSide CastlingRight = CastlingRight(2)
	blackXSide CastlingRight = CastlingRight(1 + 2)
	WhiteQSide CastlingRight = CastlingRight(4)
	WhiteKSide CastlingRight = CastlingRight(8)
	whiteXSide CastlingRight = CastlingRight(4 + 8)
)

// BlackCanCastle returns true if black can castle
func (right CastlingRight) BlackCanCastle() bool {
	return right&blackXSide != NoRights
}

// WhiteCanCastle returns true if white can castle
func (right CastlingRight) WhiteCanCastle() bool {
	return right&whiteXSide != NoRights
}

// RightsForSide returns the set of rights for a specific side
func (right CastlingRight) RightsForSide(side side.Side) CastlingRight {
	sideMask := []CastlingRight{whiteXSide, blackXSide}[side]
	return sideMask & right
}

// HasRight returns true if a given right exists in the set of rights
func (rightA CastlingRight) HasRight(rightB CastlingRight) bool {
	return rightA&rightB != NoRights
}

// Rights returns an array of individual rights
func (right CastlingRight) Rights() []CastlingRight {
	rights := []CastlingRight{}
	for i := CastlingRight(0); i < 4; i++ {
		r := CastlingRight(1) << i
		if right&r != NoRights {
			rights = append(rights, r)
		}
	}
	return rights
}

var rightNames = map[CastlingRight]string{
	BlackQSide: "q",
	BlackKSide: "k",
	WhiteQSide: "Q",
	WhiteKSide: "K",
}

func (right CastlingRight) String() string {
	ret := ""
	for _, r := range right.Rights() {
		ret = ret + rightNames[r]
	}
	return ret
}

// Parse parses a string representation of a set of castling rights
func Parse(str string) (CastlingRight, error) {
	rights := NoRights

	charToRights := map[byte]CastlingRight{
		'q': BlackQSide,
		'k': BlackKSide,
		'Q': WhiteQSide,
		'K': WhiteKSide,
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
