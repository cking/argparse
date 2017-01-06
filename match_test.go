package argparse_test

import (
	. "github.com/cking/argparse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Match", func() {
	format := "int: <i:int>, string: <s>, optional: [o]"
	resp := "int: 42, string: str, optional: "
	parser := New(format)

	Context("with a valid input", func() {
		matches, err := parser.Parse(resp)

		It("should not error out", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(matches).ToNot(BeNil())
		})

		It("should have a match `i` and `s`", func() {
			Expect(matches.HasMatch("i")).To(BeTrue())
			Expect(matches.HasMatch("s")).To(BeTrue())
		})

		It("should not have a match `o`", func() {
			Expect(matches.HasMatch("o")).To(BeFalse())
		})

		It("should have a match `i` of type int", func() {
			Expect(matches.GetInteger("i")).To(BeNumerically("==", 42))
			Expect(matches.GetString("i")).To(BeNil())
		})

		It("should have a match `s` of type string", func() {
			Expect(matches.GetString("s")).To(BeEquivalentTo("str"))
			Expect(matches.GetInteger("s")).To(BeNil())
		})

		It("should have a match `o` of value nil", func() {
			Expect(matches.GetMatch("o")).To(BeNil())
		})
	})

	Context("with an invalid input", func() {
		_, err := parser.Parse("")

		It("should error out", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})
