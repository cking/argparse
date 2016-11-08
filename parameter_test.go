package argparse_test

import (
	. "github.com/cking/argparse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parameter", func() {
	Context("With the default matcher and converter", func() {
		param := NewParameter()
		input := "this is the test string"

		It("should match the test string", func() {
			Expect(param.Matches(input)).To(Equal(true))
		})
		It("should convert the match to string", func() {
			vMatch, _, _ := param.Match(input)
			_, found := vMatch.(string)
			Expect(found).To(BeTrue())
		})
		It("should return `this` of type `string` as the match", func() {
			vMatch, _, _ := param.Match(input)
			match, _ := vMatch.(string)
			Expect(match).To(Equal("this"))
		})
		It("should return ` is the test string` as the remainder", func() {
			_, remain, _ := param.Match(input)
			Expect(remain).To(Equal(" is the test string"))
		})
	})
})
