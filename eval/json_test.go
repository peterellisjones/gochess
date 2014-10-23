package eval_test

import (
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/gochess/eval"
	"github.com/peterellisjones/gochess/piece"
	"github.com/peterellisjones/gochess/square"
)

var _ = Describe("eval", func() {
	Describe("Load", func() {
		It("loads the correct values", func() {
			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())
			eval, err := Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			Expect(eval.PieceValue(piece.WhitePawn)).To(Equal(100))
			Expect(eval.PieceSquareValue(piece.WhitePawn, square.D5)).To(Equal(25))
			Expect(eval.PieceSquareValue(piece.WhiteKing, square.A3)).To(Equal(-10))
			Expect(eval.PieceSquareValue(piece.BlackPawn, square.D2)).To(Equal(50))
		})
	})

	Describe("ToJSON", func() {
		It("create JSON", func() {
			configPath, err := filepath.Abs(path.Join("..", "assets", "default_eval_conf.json"))
			Expect(err).ToNot(HaveOccurred())
			eval, err := Load(configPath)
			Expect(err).ToNot(HaveOccurred())

			Expect(eval.ToJSON()).To(ContainSubstring("\"name\":\"knight\",\"value\":320"))
		})
	})
})
