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
			Expect(param.Matches(input)).To(BeTrue())
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

	Context("With the all matcher and converter", func() {
		param := NewDefaultParameter(AllMatcher)
		input := "this is the test string"

		It("should match the test string", func() {
			Expect(param.Matches(input)).To(BeTrue())
		})
		It("should convert the match to string", func() {
			vMatch, _, _ := param.Match(input)
			_, found := vMatch.(string)
			Expect(found).To(BeTrue())
		})
		It("should return the full test string of type `string` as the match", func() {
			vMatch, _, _ := param.Match(input)
			match, _ := vMatch.(string)
			Expect(match).To(Equal(input))
		})
		It("should return nothing as the remainder", func() {
			_, remain, _ := param.Match(input)
			Expect(remain).To(BeEmpty())
		})
	})

	Context("With the integer matcher and converter", func() {
		param := NewDefaultParameter(IntegerMatcher)
		validInput := "42"
		invalidInput := "haha, fail right?"

		It("should match the valid test string", func() {
			Expect(param.Matches(validInput)).To(BeTrue())
		})
		It("should not match the invalid test string", func() {
			Expect(param.Matches(invalidInput)).To(BeFalse())
		})
		It("should convert the valid match to integer", func() {
			vMatch, _, _ := param.Match(validInput)
			_, found := vMatch.(int)
			Expect(found).To(BeTrue())
		})
		It("should not convert the invalid match to integer", func() {
			vMatch, _, err := param.Match(invalidInput)
			_, found := vMatch.(int)
			Expect(err).To(HaveOccurred())
			Expect(found).To(BeFalse())
		})
		It("should return `42` of type `integer` as the vallid match", func() {
			vMatch, _, _ := param.Match(validInput)
			match, _ := vMatch.(int)
			Expect(match).To(Equal(42))
		})
		It("should return `nil` as the invalid match and error out", func() {
			match, _, err := param.Match(invalidInput)
			Expect(err).To(HaveOccurred())
			Expect(match).To(BeNil())
		})
		It("should return an empty string as the remainder for the valid input", func() {
			_, remain, _ := param.Match(validInput)
			Expect(remain).To(BeEmpty())
		})
		It("should return the original input as the remainder for the invalid input", func() {
			_, remain, _ := param.Match(invalidInput)
			Expect(remain).To(Equal(invalidInput))
		})
	})
})
