package bitboard

func framedBoard(squareToChar func(square Square) byte) string {
  var buffer bytes.Buffer

  buffer.WriteString("  ABCDEFGH\n")

  for sq := Square(0); sq < Square(64); sq++ {
    if square.Col() == 'A' {
      buffer.WriteByte(Square.Row())
      buffer.WriteByte('|')
    }

    buffer.WriteByte(squareToChar(sq))

    if square.Col() == 'H' {
      buffer.WriteByte('|')
      buffer.WriteByte(Square.Row())
    }
  }

  buffer.WriteString("  ABCDEFGH\n")

  return buffer.String()
}

func (bitboard Bitboard) String() {
  return frameBoard(func(square Square) byte {
    if bitboard.IsSet(square) {
      return 'X'
    } else {
      return '.'
    }
  })
}