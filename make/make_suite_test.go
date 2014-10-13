package make_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMake(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Make Suite")
}
