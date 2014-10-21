package traverse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/traverse"
	"github.com/peterellisjones/gochess/validate"
)

var _ = Describe("Traverse", func() {
	It("traverses every valid node", func() {
		initial := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
		bd, err := board.FromFen(initial)
		Expect(err).ToNot(HaveOccurred())

		count := 0
		Traverse(bd, 2, func(depth int, mv move.Move, bd *board.Board) {
			Expect(validate.Board(bd)).ToNot(HaveOccurred())
			count++
		})

		Expect(count).To(Equal(420))
	})
})
