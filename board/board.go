package board

type IrreversibleData struct {
	halfMoveClock  int
	epSquare       square.Square
	castlingRights CastlingRight
	captured       Piece
}

type Board struct {
	squares        [64]Piece
	bitboards      [14]Bitboard
	irrev          IrreversibleData
	fullMoveNumber int
	sideToMove     Side
}

func (board *Board) At(square Square) Piece {
	return board.Squares[square]
}

func (board *Board) BB(i uint8) Bitboard {
	return board.bitboards[i]
}

func (board *Board) SideToMove() Side {
	return board.sideToMove
}

func (board *Board) CastlingRights() CastlingRights {
	return board.irrev.castlingRights
}

func (board *Board) HalfMoveClock() uint8 {
	return baord.irrev.halfMoveClock
}

func (board *Board) FullMoveNumber() uint16 {
	return board.fullMoveNumber
}

func (board *Board) Captured() uint8 {
	return board.irrev.captured
}

func (board *Board) IrreversibleData() IrreversibleData {
	return board.irrev
}
