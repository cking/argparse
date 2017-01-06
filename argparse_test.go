package argparse_test

import (
	. "github.com/cking/argparse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Argparse", func() {
	format := "this is required: <required:int>, this is optional: [optional]"
	iResp := " this is required: 42            , this is optional: "
	eResp := "this is required: 42, this is optional: guess-what"
	wResp := "this is required: a, this is optional: wrong format"
	oResp := "this is required 42, this is optional: "
	fResp := "this is required: 42, this is optional: guess-what asdew"

	Context("With an empty format string", func() {
		emptyParser := New("")
		It("should return an empty format string", func() {
			Expect(emptyParser.Format()).To(BeEmpty())
		})

		It("should parse an empty input string", func() {
			_, err := emptyParser.Parse("")
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not parse an input string with content", func() {
			_, err := emptyParser.Parse("xxx")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("With custom parameters", func() {
		parser := New("")
		knownParameter := parser.Parameter("known")

		It("should return the same parameter definition for kown parameters", func() {
			p := parser.Parameter("known")
			Expect(p).To(BeIdenticalTo(knownParameter))
		})

		It("should return a default parameter definition for unkown parameters", func() {
			p := parser.Parameter("unknown")
			Expect(p).ToNot(BeIdenticalTo(knownParameter))
			Expect(p).ToNot(BeNil())
		})

		It("should set a single parameter and return the same instance", func() {
			p := NewParameter()
			parser.SetParameter("single", p)
			Expect(parser.Parameter("single")).To(BeIdenticalTo(p))
		})

		It("should set a multiple parameters and return the same instances", func() {
			p1 := NewParameter()
			p2 := NewParameter()
			parser.SetParameters(map[string]*Parameter{
				"multi1": p1,
				"multi2": p2,
			})
			Expect(parser.Parameter("multi1")).To(BeIdenticalTo(p1))
			Expect(parser.Parameter("multi2")).To(BeIdenticalTo(p2))
		})
	})

	Context("With explicit whitespace", func() {
		explicitParser := New(format)

		It("should return the same format string", func() {
			Expect(explicitParser.Format()).To(Equal(format))
		})

		It("should not parse an empty input string", func() {
			_, err := explicitParser.Parse("")
			Expect(err).To(HaveOccurred())
		})

		It("should not parse an input string with extra junk", func() {
			_, err := explicitParser.Parse(fResp)
			Expect(err).To(HaveOccurred())
		})

		It("should match the input", func() {
			_, err := explicitParser.Parse(eResp)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match the input with extra whitespace", func() {
			_, err := explicitParser.Parse(iResp)
			Expect(err).To(HaveOccurred())
		})

		It("should not match the input with the wrong format", func() {
			_, err := explicitParser.Parse(wResp)
			Expect(err).To(HaveOccurred())
		})

		It("should match the input without the optional part", func() {
			_, err := explicitParser.Parse(oResp)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("With implicit whitespace", func() {
		implicitParser := NewWithoutWhitespace(format)

		It("should return the same format string", func() {
			Expect(implicitParser.Format()).To(Equal(format))
		})

		It("should not parse an empty input string", func() {
			_, err := implicitParser.Parse("")
			Expect(err).To(HaveOccurred())
		})

		It("should not parse an input string with extra junk", func() {
			_, err := implicitParser.Parse(fResp)
			Expect(err).To(HaveOccurred())
		})

		It("should match the input", func() {
			_, err := implicitParser.Parse(eResp)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should match the input with extra whitespace", func() {
			_, err := implicitParser.Parse(iResp)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match the input with the wrong format", func() {
			_, err := implicitParser.Parse(wResp)
			Expect(err).To(HaveOccurred())
		})

		It("should match the input without the optional part", func() {
			_, err := implicitParser.Parse(oResp)
			Expect(err).To(HaveOccurred())
		})
	})
})
