package move

import (
	"errors"

	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

func (move Move) String() string {
	if move.IsCastle() {
		if move == QueenSideCastle {
			return "O-O-O"
		} else if move == KingSideCastle {
			return "O-O"
		}
	}

	str := move.From().String()
	if move.IsCapture() {
		str += "x"
	}
	str += move.To().String()
	if move.IsEpCapture() {
		str += "e.p."
	}
	if move.IsPromotion() {
		str += "=" + string(move.PromoteTo().Type().Char())
	}
	if move.IsDoublePawnPush() {
		str += ":"
	}
	return str
}

// Binary returns a string explaining the packed move bits
func (move Move) Binary() string {
	str := ""
	for i := Move(0); i < Move(16); i++ {
		if move&(Move(1)<<i) != Move(0) {
			str += "X"
		} else {
			str += "."
		}
	}
	return str + "\n" + "FFFFFFTTTTTT1234" + "\n"
}

// Parse parses a string representation of a move
func Parse(str string) (Move, error) {
	// castles
	if str == "O-O-O" {
		return QueenSideCastle, nil
	}
	if str == "O-O" {
		return KingSideCastle, nil
	}

	if len(str) < 4 {
		return Move(0), errors.New("Invalid move")
	}

	fromStr := str[0:2]
	from, err := square.Parse(fromStr)
	if err != nil {
		return Move(0), err
	}

	var toStr string
	if str[2] == 'x' {
		toStr = str[3:5]
	} else {
		toStr = str[2:4]
	}
	to, err := square.Parse(toStr)
	if err != nil {
		return Move(0), err
	}

	// quiet moves
	if len(str) == 4 {
		return EncodeMove(from, to), nil
	}

	// double pawn push
	if len(str) == 5 && str[4] == ':' {
		return EncodeDoublePawnPush(from, to), nil
	}

	// captures
	if len(str) == 5 && str[2] == 'x' {
		return EncodeCapture(from, to), nil
	}

	// en-passant capture
	if len(str) == 9 && str[2] == 'x' && str[5:9] == "e.p." {
		return EncodeEpCapture(from, to), nil
	}

	// promotions
	if len(str) == 6 && str[4] == '=' {
		pc, err := piece.Parse(str[5])
		if err != nil {
			return Move(0), err
		}
		return EncodePromotion(from, to, pc), nil
	}

	// capture promotions
	if len(str) == 7 && str[2] == 'x' && str[5] == '=' {
		pc, err := piece.Parse(str[6])
		if err != nil {
			return Move(0), err
		}

		return EncodeCapturePromotion(from, to, pc), nil
	}
	return Move(0), nil
}
