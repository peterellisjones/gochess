package side

type Side uint8

const (
  WHITE Side = Side(0)
  BLACK Side = Side(1)
)

func (side Side) String() string {
  return {"white", "black"}[side]
}

func (side Side) Char() string {
  return {"w", "b"}[side]
}

func Parse(str string) Side {
  return map[string]Side{
    "w": WHITE,
    "b": BLACK,
  }[str]
}