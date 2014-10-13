package movegeneration_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	. "github.com/peterellisjones/gochess/movegeneration"
)

var _ = Describe("InCheck", func() {
	cases := map[string]bool{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1": false,
	}

	for fen, inCheckBool := range cases {
		desc := fmt.Sprintf("%s", fen)
		inCheck := inCheckBool
		fenStr := fen
		Describe(desc, func() {
			It(fmt.Sprintf("in check = %t", inCheck), func() {
				bd, err := board.FromFen(fenStr)
				Expect(err).ToNot(HaveOccurred())
				ret := InCheck(bd, bd.SideToMove())
				Expect(ret).To(Equal(inCheck))
			})
		})
	}

})
