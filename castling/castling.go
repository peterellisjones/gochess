package castling

type CastlingRight uint8

const (
  NO_RIGHTS   CastlingRight = CastlingRight(0)
  BLACK_QSIDE CastlingRight = CastlingRight(1)
  BLACK_KSIDE CastlingRight = CastlingRight(2)
  BLACK_XSIDE CastlingRight = CastlingRight(1 + 2)
  WHITE_QSIDE CastlingRight = CastlingRight(4)
  WHITE_KSIDE CastlingRight = CastlingRight(8)
  WHITE_XSIDE CastlingRight = CastlingRight(4 + 8)
)

func (right CastlingRight) BlackCanCastle() bool {
  return right&BLACK_XSIDE != NO_RIGHTS
}

func (right CastlingRight) WhiteCanCastle() bool {
  return right&WHITE_XSIDE != NO_RIGHTS
}

func (right CastlingRight) RightsForSide(side Side) CastlingRight {
  sideMask := []CastlingRight{WHITE_XSIDE, BLACK_XSIDE}
  return sideMask[right] & right
}

func (rightA CastlingRight) HasRight(rightB CastlingRight) bool {
  return rightA&rightB != NO_RIGHTS
}

func (right CastlingRight) Rights() []CastlingRight {
  rights := []CastlingRight{}
  for i := CastlingRight(0); i < 4; i++ {
    r := CastlingRight(1) << i
    if right&r != NO_RIGHTS {
      rights = append(rights, r)
    }
  }
  return rights
}

var rightNames map[castlingRight]string := map[castlingRight]string{
  BLACK_QSIDE: "q",
  BLACK_KSIDE: "k",
  WHITE_QSIDE: "Q",
  WHITE_KSIDE: "K",
}

func (right CastlingRight) String() string {
  ret := ""
  for _, r := range right.Rights() {
    ret = ret + rightNames[r]
  }
}
