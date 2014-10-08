package square

type Square uint8

const (
  A1 Square = Square(0)
  B1 Square = Square(1)
  C1 Square = Square(2)
  D1 Square = Square(3)
  E1 Square = Square(4)
  F1 Square = Square(5)
  G1 Square = Square(6)
  H1 Square = Square(7)
  A2 Square = Square(8)
  B2 Square = Square(9)
  C2 Square = Square(10)
  D2 Square = Square(11)
  E2 Square = Square(12)
  F2 Square = Square(13)
  G2 Square = Square(14)
  H2 Square = Square(15)
  A3 Square = Square(16)
  B3 Square = Square(17)
  C3 Square = Square(18)
  D3 Square = Square(19)
  E3 Square = Square(20)
  F3 Square = Square(21)
  G3 Square = Square(22)
  H3 Square = Square(23)
  A4 Square = Square(24)
  B4 Square = Square(25)
  C4 Square = Square(26)
  D4 Square = Square(27)
  E4 Square = Square(28)
  F4 Square = Square(29)
  G4 Square = Square(30)
  H4 Square = Square(31)
  A5 Square = Square(32)
  B5 Square = Square(33)
  C5 Square = Square(34)
  D5 Square = Square(35)
  E5 Square = Square(36)
  F5 Square = Square(37)
  G5 Square = Square(38)
  H5 Square = Square(39)
  A6 Square = Square(40)
  B6 Square = Square(41)
  C6 Square = Square(42)
  D6 Square = Square(43)
  E6 Square = Square(44)
  F6 Square = Square(45)
  G6 Square = Square(46)
  H6 Square = Square(47)
  A7 Square = Square(48)
  B7 Square = Square(49)
  C7 Square = Square(50)
  D7 Square = Square(51)
  E7 Square = Square(52)
  F7 Square = Square(53)
  G7 Square = Square(54)
  H7 Square = Square(55)
  A8 Square = Square(56)
  B8 Square = Square(57)
  C8 Square = Square(58)
  D8 Square = Square(59)
  E8 Square = Square(60)
  F8 Square = Square(61)
  G8 Square = Square(62)
  H8 Square = Square(63)
  NULL Square = Square(64)
)

var squareNames[65]string = [65]string{
  "a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
  "a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
  "a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
  "a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
  "a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
  "a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
  "a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
  "a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
  "NULL",
}

func (square Square) String() stirng {
  return squareNames[square]
}