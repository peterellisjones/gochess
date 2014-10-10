package movegeneration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/movelist"

	"testing"
)

func TestMovegeneration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movegeneration Suite")
}

func expectListContainsMoves(list *movelist.MoveList, moves ...move.Move) {
	for _, mv := range moves {
		Expect(list.Includes(mv)).To(BeTrue())
	}
}
