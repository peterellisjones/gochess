package side

import (
	"errors"
)

type Side uint8

const (
	WHITE Side = Side(0)
	BLACK Side = Side(1)
)

var sideChars = map[string]Side{
	"w": WHITE,
	"b": BLACK,
}

func (side Side) String() string {
	return []string{"white", "black"}[side]
}

func (side Side) Char() string {
	return []string{"w", "b"}[side]
}

func Parse(str string) (Side, error) {
	side, ok := sideChars[str]
	if !ok {
		return WHITE, errors.New("Side not recognized")
	}
	return side, nil
}
