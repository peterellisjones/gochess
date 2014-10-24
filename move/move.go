package move

import (
	pc "github.com/peterellisjones/gochess/piece"
	sq "github.com/peterellisjones/gochess/square"
)

// Move represents a packed chess move
type Move uint16

const (
	promoFlag   Move = 1 << 15
	captureFlag Move = 1 << 14

	epCaptureFlag      Move = captureFlag | (1 << 12)
	doublePawnPushFlag Move = 1 << 13
)

// Premade moves
const (
	KingSideCastle  Move = Move(1 << 12)
	QueenSideCastle Move = Move(3 << 12)
)

// Null Move
const (
	Null = Move(0)
)

/*
  Moves are packed in order to use as little memory as possible

  pacKing
  bits  0 -  5 : from square
        6 - 11 : to square
        12     : qs castle
        13     : ks castle (with qs castle)

        12     : ep capture flag (with capture flag too)
        13     : double push flag
        14     : capture flag
        15     : promotion flag

        for a promotion bits #12 and #13 store the
        type of piece to promote to

*/

// CastleType returns the type of castle
func (move Move) CastleType() int {
	return int(move>>13) & 1
}

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
	const promoPieceMask = Move(3 << 12)
	return pc.Piece(((move & promoPieceMask) >> 11) + Move(4))
}

// IsQuiet returns true if the move is a quiet move
func (move Move) IsQuiet() bool {
	return (move & (Move(15) << 12)) == Move(0)
}

// IsCapture returns true if the move is a capture
func (move Move) IsCapture() bool {
	return (move & captureFlag) != Move(0)
}

// IsEpCapture returns true if the move is an en-passant capture
func (move Move) IsEpCapture() bool {
	const mask = Move(15 << 12)
	return (move & mask) == epCaptureFlag
}

// IsDoublePawnPush returns true if the move is a doulbe Pawn push
func (move Move) IsDoublePawnPush() bool {
	const mask = Move(15 << 12)
	return (move & mask) == doublePawnPushFlag
}

// IsCastle returns true if the move is a castle
func (move Move) IsCastle() bool {
	const mask = Move(13 << 12)
	return (move & mask) == KingSideCastle
}

// IsPromotion returns true if the move is a promotion
func (move Move) IsPromotion() bool {
	return (move & promoFlag) != Move(0)
}

// EncodeMove encodes a regular move
func EncodeMove(from sq.Square, to sq.Square) Move {
	return Move(Move(from) | (Move(to) << 6))
}

// EncodeCapture encodes a capture
func EncodeCapture(from sq.Square, to sq.Square) Move {
	return Move(captureFlag | (Move(from) | (Move(to) << 6)))
}

// EncodeEpCapture encodes an EP capture
func EncodeEpCapture(from sq.Square, to sq.Square) Move {
	return Move(epCaptureFlag | (Move(from) | (Move(to) << 6)))
}

// EncodeDoublePawnPush encodes a double Pawn push
func EncodeDoublePawnPush(from sq.Square, to sq.Square) Move {
	return Move(doublePawnPushFlag | (Move(from) | (Move(to) << 6)))
}

// EncodePromotion encodes a promotion
func EncodePromotion(from sq.Square, to sq.Square, piece pc.Piece) Move {
	flags := promoFlag | ((Move(piece.Type()) - Move(4)) << 11)
	return Move(flags | (Move(from) | (Move(to) << 6)))
}

// EncodeCapturePromotion encodes a capturing promotion
func EncodeCapturePromotion(from sq.Square, to sq.Square, piece pc.Piece) Move {
	flags := captureFlag | promoFlag | ((Move(piece.Type()) - Move(4)) << 11)
	return Move(flags | (Move(from) | (Move(to) << 6)))
}

// Distance returns the absolute difference between the starting square
//  and ending square
func (move Move) Distance() sq.Square {
	if move.IsCastle() {
		return sq.Square(0)
	} else if move.To() > move.From() {
		return move.To() - move.From()
	}
	return move.From() - move.To()
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
