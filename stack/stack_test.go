package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/stack"
	"github.com/peterellisjones/gochess/validate"
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

	It("makes and unmakes a series of moves", func() {

		initial := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
		bd, err := board.FromFen(initial)
		Expect(err).ToNot(HaveOccurred())

		st := New(bd)
		Expect(st.Depth()).To(Equal(0))

		moves := []string{
			"e2e4:", "d7d6", "d2d4", "g8f6",
			"b1c3", "g7g6", "c1e3", "f8g7",
		}

		for i := 0; i < len(moves); i++ {
			mv, err := move.Parse(moves[i])
			Expect(err).ToNot(HaveOccurred())

			st.Make(mv)
			Expect(st.Depth()).To(Equal(i + 1))

			err = validate.Board(bd)
			Expect(err).ToNot(HaveOccurred())
		}

		for i := len(moves) - 1; i >= 0; i-- {
			st.UnMake()
			Expect(st.Depth()).To(Equal(i))

			err = validate.Board(bd)
			Expect(err).ToNot(HaveOccurred())
		}

		Expect(bd.Fen()).To(Equal(initial))
	})
})
