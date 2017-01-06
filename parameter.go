package argparse

import (
	"errors"
	"regexp"
	"strconv"
)

// ParameterMatcher defines the default matching algorithms
type ParameterMatcher int

const (
	// StringMatcher returns the default implementation
	StringMatcher ParameterMatcher = 0
	// IntegerMatcher converts the input to an integer
	IntegerMatcher = iota
	// AllMatcher matches everything leftover
	AllMatcher = iota
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

// NewDefaultParameter creates a new parameter definition using one of the default sets
func NewDefaultParameter(matcher ParameterMatcher) *Parameter {
	param := NewParameter()

	switch matcher {
	case IntegerMatcher:
		param.SetMatcherRegexp(regexp.MustCompile(`^(\d+)`))
		param.SetConverter(func(input string) (interface{}, error) {
			return strconv.Atoi(input)
		})
	case AllMatcher:
		param.SetMatcherRegexp(regexp.MustCompile(`(.*)`))
	}

	return param
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

// SetMatcherRegexp sets the matching algorithm using a regexp
func (p *Parameter) SetMatcherRegexp(re *regexp.Regexp) {
	p.SetMatcher(func(input string) (string, string, bool) {
		if !re.MatchString(input) {
			return "", input, false
		}

		matches := re.FindStringSubmatch(input)
		return matches[1], input[len(matches[0]):], true
	})
}

// SetConverter changes the convert function
func (p *Parameter) SetConverter(converter func(string) (interface{}, error)) {
	p.converter = converter
}
