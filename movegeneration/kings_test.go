package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	. "github.com/peterellisjones/gochess/movegeneration"
	"github.com/peterellisjones/gochess/movelist"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("GenerateKingMoves", func() {
	It("generates the correct moves", func() {
		bd, err := board.FromFen("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -")
		Expect(err).ToNot(HaveOccurred())
		list := movelist.New()
		GenerateKingMoves(bd, list)
		Expect(list.Length()).To(Equal(2))
		expectedMoves := []move.Move{
			move.EncodeMove(square.E1, square.F1),
			move.EncodeMove(square.E1, square.D1),
		}
		expectListContainsMoves(list, expectedMoves...)

		bd, err = board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		Expect(err).ToNot(HaveOccurred())
		list = movelist.New()
		GenerateKingMoves(bd, list)
		Expect(list.Length()).To(Equal(0))
	})
})
