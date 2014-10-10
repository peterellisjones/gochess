package move

import (
	pc "github.com/peterellisjones/gochess/piece"
	sq "github.com/peterellisjones/gochess/square"
)

// Move represents a packed chess move
type Move uint16

const (
	promoFlag   Move = 1 << 12
	captureFlag Move = 2 << 12

	epCaptureFlag      Move = captureFlag | (4 << 12)
	doublePawnPushFlag Move = 8 << 12
)

// Premade moves
const (
	QueenSideCastle Move = Move(4 << 12)
	KingSideCastle  Move = Move(12 << 12)
)

/*
  Moves are packed in order to use as little memory as possible

  pacKing
  bits  0 -  6 : from square
        7 - 12 : to square
        12     : promotion flag
        13     : capture flag
        14     : castle flag
*/

// From returns the square to move from (if any)
func (move Move) From() sq.Square {
	from := Move(move) & Move(63)
	return sq.Square(from)
}

// To returns the square to move to (if any)
func (move Move) To() sq.Square {
	to := (move >> Move(6)) & Move(63)
	return sq.Square(to) // 6 bits to square
}

// PromoteTo returns the piece to promote to (if any)
func (move Move) PromoteTo() pc.Piece {
	return pc.Piece(Move(14)&(move>>Move(13)) + Move(4))
}

// IsQuiet returns true if the move is a quiet move
func (move Move) IsQuiet() bool {
	return (move & (Move(15) << Move(12))) == Move(0)
}

// IsCapture returns true if the move is a capture
func (move Move) IsCapture() bool {
	return (move & captureFlag) != Move(0)
}

// IsEpCapture returns true if the move is an en-passant capture
func (move Move) IsEpCapture() bool {
	return (move & (Move(15) << Move(12))) == epCaptureFlag
}

// IsDoublePawnPush returns true if the move is a doulbe Pawn push
func (move Move) IsDoublePawnPush() bool {
	return (move & (Move(15) << Move(12))) == doublePawnPushFlag
}

// IsCastle returns true if the move is a castle
func (move Move) IsCastle() bool {
	return (move & (Move(7) << Move(12))) == (Move(4) << Move(12))
}

// IsPromotion returns true if the move is a promotion
func (move Move) IsPromotion() bool {
	return (move & promoFlag) != Move(0)
}

// EncodeMove encodes a regular move
func EncodeMove(from sq.Square, to sq.Square) Move {
	return Move(Move(from) | (Move(to) << Move(6)))
}

// EncodeCapture encodes a capture
func EncodeCapture(from sq.Square, to sq.Square) Move {
	return Move(captureFlag | (Move(from) | (Move(to) << Move(6))))
}

// EncodeEpCapture encodes an EP capture
func EncodeEpCapture(from sq.Square, to sq.Square) Move {
	return Move(epCaptureFlag | (Move(from) | (Move(to) << Move(6))))
}

// EncodeDoublePawnPush encodes a double Pawn push
func EncodeDoublePawnPush(from sq.Square, to sq.Square) Move {
	return Move(doublePawnPushFlag | (Move(from) | (Move(to) << Move(6))))
}

// EncodePromotion encodes a promotion
func EncodePromotion(from sq.Square, to sq.Square, piece pc.Piece) Move {
	flags := promoFlag | ((Move(piece) - Move(4)) << Move(13))
	return Move(flags | (Move(from) | (Move(to) << Move(6))))
}

// EncodeCapturePromotion encodes a capturing promotion
func EncodeCapturePromotion(from sq.Square, to sq.Square, piece pc.Piece) Move {
	flags := captureFlag | promoFlag | ((Move(piece) - Move(4)) << Move(13))
	return Move(flags | (Move(from) | (Move(to) << Move(6))))
}

// Flip inverses the move
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
