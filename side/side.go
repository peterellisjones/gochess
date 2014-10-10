package side

import (
	"errors"
)

// Side represents the side to play
type Side uint8

// The possible sides
const (
	White Side = Side(0)
	Black Side = Side(1)
)

var sideChars = map[byte]Side{
	'w': White,
	'b': Black,
}

func (side Side) String() string {
	return []string{"white", "black"}[side]
}

// Char returns the char representation of a side
func (side Side) Char() string {
	return []string{"w", "b"}[side]
}

// Parse returns the side given a char
func Parse(str byte) (Side, error) {
	side, ok := sideChars[str]
	if !ok {
		return White, errors.New("Side not recognized")
	}
	return side, nil
}
