package piece

var pieceChars map[Piece]byte = map[Piece]byte{
  EMPTY: '.',
  ERROR: 'e',
  WHITE_KNIGHT: 'N',
  BLACK_KNIGHT: 'n',
  WHITE_BISHOP: 'B',
  BLACK_BISHOP: 'b',
  WHITE_ROOK: 'R',
  BLACK_ROOK: 'r',
  WHITE_QUEEN: 'Q',
  BLACK_QUEEN: 'q',
  WHITE_KING: 'K',
  BLACK_KING: 'k',
}

var charPieces map[string]Piece = map[string]Piece{
  "N": WHITE_KNIGHT
  "n": BLACK_KNIGHT
  "B": WHITE_BISHOP
  "b": BLACK_BISHOP
  "R": WHITE_ROOK
  "r": BLACK_ROOK
  "Q": WHITE_QUEEN
  "q": BLACK_QUEEN
  "K": WHITE_KING
  "k": BLACK_KING
}

var pieceNames map[Piece]string = map[Piece]string{
  EMPTY: "empty",
  KNIGHT: "knight",
  BISHOP: "bishop",
  ROOK: "rook",
  QUEEN: "queen",
  KING: "king",
}

func (piece Piece) Side() Side {
  return Side(piece & 1)
}

func (piece Piece) Type() Piece {
  return piece & 0xFE
}

func (piece Piece) String() string {
  return piece.Side().String() + " " + pieceNames[piece.Type()]
}

func (piece Piece) Char() byte {
  return pieceChars[piece]
}

func (char byte) Parse() Piece {
  return charPieces[char]
}