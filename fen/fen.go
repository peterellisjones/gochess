package fen

import (
	"errors"
	"github.com/peterellisjones/gochess/castling"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/side"
	"github.com/peterellisjones/gochess/square"
	"strconv"
	"strings"
)

type Parts struct {
	Board          [64]piece.Piece
	SideToMove     side.Side
	CastlingRights castling.CastlingRight
	EpSquare       square.Square
	HalfMoveClock  int
	FullMoveNumber int
}

func FenParts(fen string) (Parts, error) {
	parts := Parts{
		SideToMove:     side.WHITE,
		CastlingRights: castling.NO_RIGHTS,
		EpSquare:       square.NULL,
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
		rights, err := parseCastlingRights(arr[2])
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
	return strconv.Atoi(str)
}

func parseFullMoveNumber(str string) (int, error) {
	return strconv.Atoi(str)
}

func parseEpSquare(str string) (square.Square, error) {
	return square.Parse(str)
}

func parseSideToMove(str string) (side.Side, error) {
	return side.Parse(str)
}

func parseCastlingRights(str string) (castling.CastlingRight, error) {
	return castling.Parse(str)
}
