package eval

import (
	"encoding/json"
	"io/ioutil"

	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

type pieceValuesJSON struct {
	Name         string    `json:"name"`
	Value        int       `json:"value"`
	SquareValues [8][8]int `json:"square values"`
}

type decodedJSON struct {
	PieceValues []pieceValuesJSON `json:"pieces"`
}

func (eval *Eval) ToJSON() string {
	d := decodedJSON{
		PieceValues: []pieceValuesJSON{},
	}

	for i := 0; i < 6; i++ {
		pc := piece.Piece(i*2 + 2)
		pieceValues := pieceValuesJSON{}
		pieceValues.Name = pc.TypeName()
		pieceValues.Value = eval.PieceValue(pc)

		squareValues := [8][8]int{}

		square.ForEach(func(sq square.Square) {
			squareValues[sq.Flip()>>3][sq&7] = eval.PieceSquareValue(pc, sq)
		})

		pieceValues.SquareValues = squareValues
		d.PieceValues = append(d.PieceValues, pieceValues)
	}

	bytes, _ := json.Marshal(d)
	return string(bytes)
}

func Load(path string) (*Eval, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	d := decodedJSON{
		PieceValues: []pieceValuesJSON{},
	}

	err = json.Unmarshal(bytes, &d)
	if err != nil {
		return nil, err
	}

	eval := Eval{
		pieceValues:       pieceValuesFromJSON(d),
		pieceSquareValues: pieceSquareValuesFromJSON(d),
	}

	eval.castleValues[0] = 0 +
		eval.PieceSquareValue(piece.WhiteKing, square.G1) -
		eval.PieceSquareValue(piece.WhiteKing, square.E1) +
		eval.PieceSquareValue(piece.WhiteRook, square.F1) -
		eval.PieceSquareValue(piece.WhiteRook, square.H1)

	eval.castleValues[1] = 0 +
		eval.PieceSquareValue(piece.WhiteKing, square.C1) -
		eval.PieceSquareValue(piece.WhiteKing, square.E1) +
		eval.PieceSquareValue(piece.WhiteRook, square.D1) -
		eval.PieceSquareValue(piece.WhiteRook, square.A1)

	return &eval, nil
}

func pieceValuesFromJSON(d decodedJSON) [6]int {
	arr := [6]int{}

	for _, info := range d.PieceValues {
		pc := piece.ParseType(info.Name)
		arr[pc.Idx()>>1] = info.Value
	}

	return arr
}

func pieceSquareValuesFromJSON(d decodedJSON) [6][64]int {
	arr := [6][64]int{}

	for _, info := range d.PieceValues {
		pc := piece.ParseType(info.Name)
		for rIdx, row := range info.SquareValues {
			for colIdx, val := range row {
				square := 8*(7-rIdx) + colIdx
				if square >= 0 && square < 64 {
					arr[pc.Idx()>>1][square] = val
				}
			}
		}
	}

	return arr
}
