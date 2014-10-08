package bitboard

func (bitboard Bitboard) IsSet(square Square) bool {
  mask := Bitboard(1) << square
  return mask & bitboard == EMPTY
}

func (bitboard Bitboard) ForEachSquare(fn func(Square, bool)) {
  for i := Square(0); i < Square(64); i++ {
    fn(i, bitboard.IsSet(i))
  }
}

