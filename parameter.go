package argparse

import (
	"errors"
	"regexp"
)

var (
	defaultMatcherRegExp = regexp.MustCompile(`^([-_\w\d]*)`)
)

// Parameter object
type Parameter struct {
	matcher   func(string) (string, string, bool)
	converter func(string) (interface{}, error)
}

func defaultMatcher(input string) (string, string, bool) {
	matches := defaultMatcherRegExp.FindStringSubmatch(input)
	return matches[1], input[len(matches[0]):], true
}

func defaultConverter(input string) (interface{}, error) {
	return input, nil
}

// NewParameter creates a new parameter definition
func NewParameter() *Parameter {
	return &Parameter{defaultMatcher, defaultConverter}
}

// Matches returns a flag if the input matches the expected format
func (p *Parameter) Matches(input string) bool {
	_, _, ok := p.matcher(input)
	return ok
}

// Match splits the input to the expected match and returns the remaining string
func (p *Parameter) Match(input string) (interface{}, string, error) {
	match, remaining, ok := p.matcher(input)
	if !ok {
		return nil, remaining, errors.New("Could not match the parameter")
	}

	obj, err := p.converter(match)
	return obj, remaining, err
}

// SetMatcher sets the matching algorithm for this parameter
func (p *Parameter) SetMatcher(matcher func(string) (string, string, bool)) {
	p.matcher = matcher
}

// SetConverter changes the convert function
func (p *Parameter) SetConverter(converter func(string) (interface{}, error)) {
	p.converter = converter
}
