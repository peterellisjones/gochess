package fen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fen Suite")
}
