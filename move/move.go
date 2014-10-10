package move

import (
	"fmt"
	pc "github.com/peterellisjones/gochess/piece"
	sq "github.com/peterellisjones/gochess/square"
	"strings"
)

type Move uint16

const PROMO_FLAG uint16 = 1 << 12
const CAPTURE_FLAG uint16 = 2 << 12

const EP_CAPTURE_FLAG uint16 = CAPTURE_FLAG | (4 << 12)
const DOUBLE_PAWN_PUSH_FLAG uint16 = 8 << 12

const QSIDE_CASTLE Move = Move(4 << 12)
const KSIDE_CASTLE Move = Move(12 << 12)

func (move Move) From() sq.Square {
	from := uint16(move) & uint16(63)
	return sq.Square(from)
}

func (move Move) To() sq.Square {
	to := (uint16(move) >> 6) & uint16(63)
	return sq.Square(to) // 6 bits to square
}

func (move Move) PromoteTo() pc.Piece {
	return pc.Piece(14&(uint16(move)>>13) + 4)
}

func (move Move) IsQuiet() bool {
	return (uint16(move) & (15 << 12)) == 0
}

func (move Move) IsCapture() bool {
	return (uint16(move) & CAPTURE_FLAG) != 0
}

func (move Move) IsEpCapture() bool {
	return (uint16(move) & (15 << 12)) == EP_CAPTURE_FLAG
}

func (move Move) IsDoublePawnPush() bool {
	return (uint16(move) & (15 << 12)) == DOUBLE_PAWN_PUSH_FLAG
}

func (move Move) IsCastle() bool {
	return (uint16(move) & (7 << 12)) == (4 << 12)
}

func (move Move) IsPromotion() bool {
	return (uint16(move) & PROMO_FLAG) != 0
}

func EncodeMove(from sq.Square, to sq.Square) Move {
	return Move(uint16(from) | (uint16(to) << 6))
}

func EncodeCapture(from sq.Square, to sq.Square) Move {
	return Move(CAPTURE_FLAG | (uint16(from) | (uint16(to) << 6)))
}

func EncodeEpCapture(from sq.Square, to sq.Square) Move {
	return Move(EP_CAPTURE_FLAG | (uint16(from) | (uint16(to) << 6)))
}

func EncodeDoublePawnPush(from sq.Square, to sq.Square) Move {
	return Move(DOUBLE_PAWN_PUSH_FLAG | (uint16(from) | (uint16(to) << 6)))
}

func EncodePromotion(from sq.Square, to sq.Square, piece pc.Piece) Move {
	flags := PROMO_FLAG | ((uint16(piece) - 4) << 13)
	return Move(flags | (uint16(from) | (uint16(to) << 6)))
}

func EncodeCapturePromotion(from sq.Square, to sq.Square, piece pc.Piece) Move {
	flags := CAPTURE_FLAG | PROMO_FLAG | ((uint16(piece) - 4) << 13)
	return Move(flags | (uint16(from) | (uint16(to) << 6)))
}

func (move Move) ToString() string {
	strs := []string{move.From().String(), move.To().String()}
	if move.IsCastle() {
		if move == QSIDE_CASTLE {
			return "O-O-O"
		} else if move == KSIDE_CASTLE {
			return "O-O"
		}
	} else if move.IsCapture() {
		return strings.Join(strs, "x")
	}
	return strings.Join(strs, "")
}

func (move Move) Debug() {
	fmt.Println(move.ToString())
}

func (move Move) Flip() Move {

	if move.IsCastle() {
		return move
	}

	// flip from square
	from := move.From()
	from = from.Flip()
	move &= ^Move(63)
	move |= Move(from)

	// flip to square
	to := move.To()
	to = to.Flip()
	move &= ^Move(63 << 6)
	move |= Move(to) << 6
	return move
}
