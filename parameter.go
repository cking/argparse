package argparse

import (
	"regexp"
	"strings"
)

var (
	defaultMatcherRegExp = regexp.MustCompile(`^(\S*)\s*`)
)

// Parameter object
type Parameter struct {
	matcher   func(string) (string, string, bool)
	converter func(string) (interface{}, error)
}

func defaultMatcher(input string) (string, string, bool) {
	if !defaultMatcherRegExp.MatchString(input) {
		return "", input, false
	}

	matches := defaultMatcherRegExp.FindStringSubmatch(input)
	return matches[1], input[len(matches[0]):], true
}

func defaultConverter(input string) (interface{}, error) {
	return input, nil
}

func newParameter() *Parameter {
	return &Parameter{defaultMatcher, defaultConverter}
}

// Matches returns a flag if the input matches the expected format
func (p *Parameter) Matches(input string) bool {
	_, _, ok := p.matcher(input)
	return !ok
}

// Match splits the input to the expected match and returns the remaining string
func (p *Parameter) Match(input string) (string, string) {
	match, remaining, _ := p.matcher(input)
	return match, strings.TrimLeft(remaining, ` `)
}

// SetMatcher sets the matching algorithm for this parameter
func (p *Parameter) SetMatcher(matcher func(string) (string, string, bool)) {
	p.matcher = matcher
}

// Convert the input into the expected object
func (p *Parameter) Convert(input string) (interface{}, error) {
	return p.converter(input)
}

// SetConverter changes the convert function
func (p *Parameter) SetConverter(converter func(string) (interface{}, error)) {
	p.converter = converter
}
