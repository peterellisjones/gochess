package perft_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/perft"
)

var _ = Describe("Perft divide and conquer", func() {

	type Case struct {
		fen     string
		depth   int
		total   int64
		results map[string]int64
	}

	It("can do stuff", func() {
		//bd, _ := board.FromFen("r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1")

	})

	cases := []Case{
		Case{
			fen:   "r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1",
			depth: 3,
			total: 9467,
			results: map[string]int64{
				"g8h8":  1753,
				"f6d5":  1687,
				"f8f7":  1623,
				"b5c4":  1352,
				"c5c4":  1409,
				"d7d5:": 1643,
			},
		},
		Case{ // f8f7 -> 1623
			fen:   "r2q2k1/pP1p1rpp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R w KQ - 1 2",
			depth: 2,
			total: 1623,
			results: map[string]int64{
				"O-O-O":   41,
				"f3g1":    37,
				"f3d4":    39,
				"f3h4":    37,
				"f3xe5":   36,
				"f3g5":    36,
				"a4c3":    38,
				"a4xc5":   36,
				"a4b6":    37,
				"a1b1":    45,
				"a1c1":    41,
				"a1d1":    41,
				"a1xa2":   37,
				"h1f1":    37,
				"h1g1":    37,
				"a6xa5":   37,
				"a6b6":    38,
				"a6c6":    39,
				"a6d6":    35,
				"a6e6":    39,
				"a6xf6":   35,
				"a6xa7":   38,
				"b3xa2":   38,
				"b3c4":    34,
				"b3d5":    37,
				"b3e6":    39,
				"b3xf7":   3,
				"g3f4":    38,
				"g3h4":    37,
				"g3xe5":   36,
				"a6xb5":   31,
				"e1d1":    37,
				"c2c3":    38,
				"d2d3":    35,
				"g2xh3":   33,
				"c2c4:":   36,
				"d2d4:":   39,
				"b7xa8=Q": 33,
				"b7xa8=N": 36,
				"b7xa8=R": 33,
				"b7xa8=B": 36,
				"b7b8=Q":  33,
				"b7b8=N":  36,
				"b7b8=R":  33,
				"b7b8=B":  36,
			},
		},
		Case{ // g8h8 => 1753
			fen:   "r2q1r1k/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R w KQ - 1 2",
			depth: 2,
			total: 1753,
			results: map[string]int64{
				"O-O-O":   42,
				"f3g1":    38,
				"f3d4":    40,
				"f3h4":    38,
				"f3xe5":   37,
				"f3g5":    37,
				"a4c3":    39,
				"a4xc5":   37,
				"a4b6":    38,
				"a1b1":    46,
				"a1c1":    42,
				"a1d1":    42,
				"a1xa2":   38,
				"h1f1":    38,
				"h1g1":    38,
				"a6xa5":   38,
				"a6b6":    39,
				"a6c6":    40,
				"a6d6":    37,
				"a6e6":    40,
				"a6xf6":   34,
				"a6xa7":   39,
				"b3xa2":   39,
				"b3c4":    35,
				"b3d5":    38,
				"b3e6":    40,
				"b3f7":    39,
				"b3g8":    40,
				"g3f4":    39,
				"g3h4":    38,
				"g3xe5":   37,
				"a6xb5":   32,
				"e1d1":    38,
				"c2c3":    39,
				"d2d3":    36,
				"g2xh3":   34,
				"c2c4:":   36,
				"d2d4:":   40,
				"b7xa8=Q": 37,
				"b7xa8=N": 37,
				"b7xa8=R": 37,
				"b7xa8=B": 37,
				"b7b8=Q":  37,
				"b7b8=N":  37,
				"b7b8=R":  37,
				"b7b8=B":  37,
			},
		},
		Case{
			fen:   "r2q1rk1/pP1p2pp/Q4n2/b1p1p3/Npb5/1B3NBn/pPPP1PPP/R3K2R w KQ - 1 2",
			depth: 2,
			total: 1352,
			results: map[string]int64{
				"O-O-O":   38,
				"f3g1":    34,
				"f3d4":    36,
				"f3h4":    34,
				"f3xe5":   32,
				"f3g5":    32,
				"a4c3":    35,
				"a4xc5":   34,
				"a4b6":    34,
				"a1b1":    42,
				"a1c1":    38,
				"a1d1":    38,
				"a1xa2":   34,
				"h1f1":    34,
				"h1g1":    34,
				"a6xa5":   34,
				"a6b6":    35,
				"a6c6":    36,
				"a6d6":    33,
				"a6e6":    4,
				"a6xf6":   32,
				"a6xa7":   35,
				"b3xa2":   36,
				"b3xc4":   4,
				"g3f4":    35,
				"g3h4":    34,
				"g3xe5":   33,
				"a6xc4":   4,
				"a6b5":    35,
				"e1d1":    34,
				"c2c3":    35,
				"d2d3":    34,
				"g2xh3":   30,
				"d2d4:":   36,
				"b7xa8=Q": 33,
				"b7xa8=N": 33,
				"b7xa8=R": 33,
				"b7xa8=B": 33,
				"b7b8=Q":  33,
				"b7b8=N":  33,
				"b7b8=R":  33,
				"b7b8=B":  33,
			},
		},
	}

	for _, c := range cases {
		cLocal := c
		It(cLocal.fen, func() {

			actual, err := PerftMoves(cLocal.fen, cLocal.depth)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(cLocal.results))

			total := int64(0)
			for _, nodes := range actual {
				total += nodes
			}

			Expect(total).To(Equal(cLocal.total))
		})
	}

})
