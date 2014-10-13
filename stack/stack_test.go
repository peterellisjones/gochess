package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/stack"
)

var _ = Describe("Make & UnMake", func() {
	It("makes and unmakes moves", func() {
		initial := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
		bd, err := board.FromFen(initial)
		Expect(err).ToNot(HaveOccurred())

		st := New(bd)
		Expect(st.Depth()).To(Equal(0))

		mv, err := move.Parse("c2c4:")
		Expect(err).ToNot(HaveOccurred())

		st.Make(mv)
		Expect(st.Depth()).To(Equal(1))

		expectedFen := "rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b KQkq c3 0 1"
		Expect(bd.Fen()).To(Equal(expectedFen))

		st.UnMake()
		Expect(st.Depth()).To(Equal(0))

		Expect(bd.Fen()).To(Equal(initial))
	})
})
