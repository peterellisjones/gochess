package castling

import (
	"errors"
	"github.com/peterellisjones/gochess/side"
)

type CastlingRight uint8

const (
	NO_RIGHTS   CastlingRight = CastlingRight(0)
	BLACK_QSIDE CastlingRight = CastlingRight(1)
	BLACK_KSIDE CastlingRight = CastlingRight(2)
	BLACK_XSIDE CastlingRight = CastlingRight(1 + 2)
	WHITE_QSIDE CastlingRight = CastlingRight(4)
	WHITE_KSIDE CastlingRight = CastlingRight(8)
	WHITE_XSIDE CastlingRight = CastlingRight(4 + 8)
)

func (right CastlingRight) BlackCanCastle() bool {
	return right&BLACK_XSIDE != NO_RIGHTS
}

func (right CastlingRight) WhiteCanCastle() bool {
	return right&WHITE_XSIDE != NO_RIGHTS
}

func (right CastlingRight) RightsForSide(side side.Side) CastlingRight {
	sideMask := []CastlingRight{WHITE_XSIDE, BLACK_XSIDE}[side]
	return sideMask & right
}

func (rightA CastlingRight) HasRight(rightB CastlingRight) bool {
	return rightA&rightB != NO_RIGHTS
}

func (right CastlingRight) Rights() []CastlingRight {
	rights := []CastlingRight{}
	for i := CastlingRight(0); i < 4; i++ {
		r := CastlingRight(1) << i
		if right&r != NO_RIGHTS {
			rights = append(rights, r)
		}
	}
	return rights
}

var rightNames = map[CastlingRight]string{
	BLACK_QSIDE: "q",
	BLACK_KSIDE: "k",
	WHITE_QSIDE: "Q",
	WHITE_KSIDE: "K",
}

func (right CastlingRight) String() string {
	ret := ""
	for _, r := range right.Rights() {
		ret = ret + rightNames[r]
	}
	return ret
}

func Parse(str string) (CastlingRight, error) {
	rights := NO_RIGHTS

	charToRights := map[byte]CastlingRight{
		'q': BLACK_QSIDE,
		'k': BLACK_KSIDE,
		'Q': WHITE_QSIDE,
		'K': WHITE_KSIDE,
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
