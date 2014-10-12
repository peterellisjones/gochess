package fen

import (
	"errors"
	"strconv"
	"strings"

	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
)

// Parts represent the individual parts of a FEN string
type Parts struct {
	Board          [64]piece.Piece
	SideToMove     side.Side
	CastlingRights castling.Right
	EpSquare       square.Square
	HalfMoveClock  int
	FullMoveNumber int
}

// GetParts returns the parts of a FEN string
func GetParts(fen string) (Parts, error) {
	parts := Parts{
		SideToMove:     side.White,
		CastlingRights: castling.NoRights,
		EpSquare:       square.Null,
		HalfMoveClock:  0,
		FullMoveNumber: 1,
	}

	arr := strings.Split(fen, " ")

	if len(arr) > 0 {
		board, err := parseBoard(arr[0])
		if err != nil {
			return parts, err
		}
		parts.Board = board
	}

	if len(arr) > 1 {
		side, err := parseSideToMove(arr[1])
		if err != nil {
			return parts, err
		}
		parts.SideToMove = side
	}

	if len(arr) > 2 {
		rights, err := parseRights(arr[2])
		if err != nil {
			return parts, err
		}
		parts.CastlingRights = rights
	}

	if len(arr) > 3 {
		ep, err := parseEpSquare(arr[3])
		if err != nil {
			return parts, err
		}
		parts.EpSquare = ep
	}

	if len(arr) > 4 {
		clock, err := parseHalfMoveClock(arr[4])
		if err != nil {
			return parts, err
		}
		parts.HalfMoveClock = clock
	}

	if len(arr) > 5 {
		num, err := parseFullMoveNumber(arr[5])
		if err != nil {
			return parts, err
		}
		parts.FullMoveNumber = num
	}

	return parts, nil
}

func parseBoard(str string) ([64]piece.Piece, error) {
	board := [64]piece.Piece{}
	rows := strings.Split(str, "/")
	if len(rows) != 8 {
		return board, errors.New("Couldn't find 8 rows in board")
	}

	for r := 0; r < 8; r++ {
		row := rows[7-r]
		col := 0
		for pos := 0; pos < len(row); pos++ {
			c := row[pos]
			square := (r << 3) | col
			if c == '/' {
				break
			} else if c >= '1' && c <= '8' {
				col += int(c - '0')
			} else {
				piece, err := piece.Parse(c)
				if err != nil {
					return board, err
				}
				board[square] = piece
				col++
			}
		}
	}

	return board, nil
}

func parseHalfMoveClock(str string) (int, error) {
	if str == "-" {
		return 0, nil
	}
	return strconv.Atoi(str)
}

func parseFullMoveNumber(str string) (int, error) {
	if str == "-" {
		return 1, nil
	}
	return strconv.Atoi(str)
}

func parseEpSquare(str string) (square.Square, error) {
	if str == "-" {
		return square.Null, nil
	}
	return square.Parse(str)
}

func parseSideToMove(str string) (side.Side, error) {
	if str == "-" {
		return side.White, nil
	}
	return side.Parse(str[0])
}

func parseRights(str string) (castling.Right, error) {
	if str == "-" {
		return castling.NoRights, nil
	}
	return castling.Parse(str)
}
