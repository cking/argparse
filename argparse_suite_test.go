package argparse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArgparse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Argparse Suite")
}
