package perft_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/perft"
)

var _ = Describe("Perft", func() {

	type Case struct {
		fen     string
		results []Result
	}

	cases := []Case{
		Case{
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			results: []Result{
				Result{Nodes: 20, Captures: 0, Checks: 0, EpCaptures: 0, Castles: 0, Promotions: 0},
				Result{Nodes: 400, Captures: 0, Checks: 0, EpCaptures: 0, Castles: 0, Promotions: 0},
				Result{Nodes: 8902, Captures: 34, Checks: 12, EpCaptures: 0, Castles: 0, Promotions: 0},
				Result{Nodes: 197281, Captures: 1576, Checks: 469, EpCaptures: 0, Castles: 0, Promotions: 0},
			},
		},
		Case{
			fen: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -",
			results: []Result{
				Result{
					Nodes:      48,
					Captures:   8,
					Checks:     0,
					EpCaptures: 0,
					Castles:    2,
					Promotions: 0,
				},
				Result{
					Nodes:      2039,
					Captures:   351,
					Checks:     3,
					EpCaptures: 1,
					Castles:    91,
					Promotions: 0,
				},
				Result{
					Nodes:      97862,
					Captures:   17102,
					Checks:     993,
					EpCaptures: 45,
					Castles:    3162,
					Promotions: 0,
				},
			},
		},
		Case{
			fen: "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w",
			results: []Result{
				Result{
					Nodes:      14,
					Captures:   1,
					EpCaptures: 0,
					Castles:    0,
					Promotions: 0,
					Checks:     2,
				},
				Result{
					Nodes:      191,
					Captures:   14,
					EpCaptures: 0,
					Castles:    0,
					Promotions: 0,
					Checks:     10,
				},
				Result{
					Nodes:      2812,
					Captures:   209,
					EpCaptures: 2,
					Castles:    0,
					Promotions: 0,
					Checks:     267,
				},
				Result{
					Nodes:      43238,
					Captures:   3348,
					EpCaptures: 123,
					Castles:    0,
					Promotions: 0,
					Checks:     1680,
				},
			},
		},
		Case{
			fen: "r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1",
			results: []Result{
				Result{
					Nodes:      6,
					Captures:   0,
					EpCaptures: 0,
					Castles:    0,
					Promotions: 0,
					Checks:     0,
				},
				Result{
					Nodes:      264,
					Captures:   87,
					EpCaptures: 0,
					Castles:    6,
					Promotions: 48,
					Checks:     10,
				},
				Result{
					Nodes:      9467,
					Captures:   1021,
					EpCaptures: 4,
					Castles:    0,
					Promotions: 120,
					Checks:     38,
				},
				Result{
					Nodes:      422333,
					Captures:   131393,
					EpCaptures: 0,
					Castles:    7795,
					Promotions: 60032,
					Checks:     15492,
				},
			},
		},
		Case{
			fen:     "r2q1rk1/pP1p2pp/Q4n2/bb2p3/Npp5/1B3NBn/pPPP1PPP/R3K2R w KQ - 0 2",
			results: []Result{},
		},
	}

	for _, c := range cases {
		cLocal := c
		It(cLocal.fen, func() {
			ex, _ := PerftMoves(cLocal.fen, len(cLocal.results))
			fmt.Println(ex.String())

			actual, err := Perft(cLocal.fen, len(cLocal.results))
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(cLocal.results))
		})
	}
})
