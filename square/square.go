package square

// Square represents a square on a chess board
type Square uint8

var columns = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
var rows = []byte{'1', '2', '3', '4', '5', '6', '7', '8'}

// Col returns the character representing that square's column
func (square Square) Col() byte {
	return columns[square&7]
}

// Row returns the character representing that square's row
func (square Square) Row() byte {
	return rows[(square&7)>>3]
}

var squareNames = [65]string{
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

func (square Square) String() string {
	return squareNames[square]
}
