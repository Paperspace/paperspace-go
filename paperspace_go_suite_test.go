package paperspace_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPaperspaceGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PaperspaceGo Suite")
}
