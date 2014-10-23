package movegeneration_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/movelist"
	"github.com/peterellisjones/gochess/validate"

	"testing"
)

func TestMovegeneration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movegeneration Suite")
}

func expectListContainsMoves(list *movelist.MoveList, moves ...string) {
	for _, mvString := range moves {
		mv, err := move.Parse(mvString)
		Expect(err).ToNot(HaveOccurred())
		Expect(list.Includes(mv)).To(BeTrue())
	}
}

type MoveGenerator func(bd *board.Board) []move.Move

func ItGeneratesMovesFor(cases map[string][]string, fn MoveGenerator) {
	for fen, expectedMoves := range cases {
		desc := fmt.Sprintf("Moves for %s", fen)
		exMoves := expectedMoves
		fenStr := fen
		Describe(desc, func() {
			for _, expectedMove := range exMoves {
				mvStr := expectedMove
				It(fmt.Sprintf("generates %s", mvStr), func() {
					bd, err := board.FromFen(fenStr)
					Expect(err).ToNot(HaveOccurred())
					err = validate.Board(bd)
					Expect(err).ToNot(HaveOccurred())

					list := fn(bd)
					mv, err := move.Parse(mvStr)
					Expect(err).ToNot(HaveOccurred())
					Expect(list).To(ContainElement(mv))
				})
			}

			It(fmt.Sprintf("generates %d moves", len(exMoves)), func() {
				bd, err := board.FromFen(fenStr)
				Expect(err).ToNot(HaveOccurred())

				list := fn(bd)
				for _, mv := range list {
					Expect(exMoves).To(ContainElement(mv.String()))
				}

				Expect(len(list)).To(Equal(len(exMoves)))
			})
		})
	}
}
