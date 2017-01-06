package argparse

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Parser object
type Parser struct {
	format     string
	parameter  ParameterMap
	expression func(string) (*Match, error)
}

type expressionType int

const (
	literalType expressionType = iota
	requiredType
	optionalType
)

type expression struct {
	etyp  expressionType
	vtyp  string
	value string
}

// ParameterMap is the map to define parameters
type ParameterMap map[string]*Parameter

var (
	formatRegexp = regexp.MustCompile(`\<(?P<rkey>\w+)(?::(?P<rtype>\w+))?\>|\[(?P<okey>\w+)(?::(?P<otype>\w+))?\]|(?:\\[\[<]|[^\[<])+`)
)

// New create a new Parser object
func New(format string) *Parser {
	p := &Parser{format, ParameterMap{
		"str": NewDefaultParameter(StringMatcher),
		"int": NewDefaultParameter(IntegerMatcher),
	},
		nil}
	p.expression = p.createExpression(format, false)
	return p
}

// NewWithoutWhitespace initializes a new Parser object and sets it to reduce multiple whitespaces to a single space
func NewWithoutWhitespace(format string) *Parser {
	p := &Parser{format, map[string]*Parameter{
		"str": NewDefaultParameter(StringMatcher),
		"int": NewDefaultParameter(IntegerMatcher),
	},
		nil}
	p.expression = p.createExpression(format, true)
	return p
}

// Format returns the format string
func (p *Parser) Format() string {
	return p.format
}

func (p *Parser) createExpression(format string, ignoreWhitespace bool) func(string) (*Match, error) {
	matches := formatRegexp.FindAllStringSubmatch(format, -1)
	if matches == nil || len(matches) == 0 {
		return func(input string) (*Match, error) {
			if len(input) == 0 {
				return NewMatch(), nil
			}
			return nil, errors.New("input didn't match format")
		}
	}

	// convert matches into an easier to use version
	var expr []expression
	rews := regexp.MustCompile(`\s+`)
	for _, match := range matches {
		if len(match[1]) > 0 {
			expr = append(expr, expression{requiredType, match[2], match[1]})
		} else if len(match[3]) > 0 {
			expr = append(expr, expression{optionalType, match[4], match[3]})
		} else {
			literal := regexp.QuoteMeta(match[0])
			if ignoreWhitespace {
				literal = rews.ReplaceAllString(strings.TrimSpace(literal), "\\s+")
			}

			expr = append(expr, expression{literalType, "", literal})
		}
	}

	return func(input string) (*Match, error) {
		var v interface{}
		m := NewMatch()

		for _, e := range expr {
			if ignoreWhitespace {
				input = strings.TrimSpace(input)
			}

			switch e.etyp {
			case literalType:
				re := regexp.MustCompile(`^` + e.value)
				if !re.MatchString(input) {
					return nil, fmt.Errorf("Couldn't match `%v` in `%v`", e.value, input)
				}

				input = re.ReplaceAllString(input, "")
				break

			case optionalType:
				def := p.Parameter(e.vtyp)
				if len(input) > 0 && def.Matches(input) {
					v, input, _ = def.Match(input)
					m.addParameter(e.value, v)
				} else {
					m.addParameter(e.value, nil)
				}
				break

			case requiredType:
				def := p.Parameter(e.vtyp)
				if def.Matches(input) {
					v, input, _ = def.Match(input)
					m.addParameter(e.value, v)
				} else {
					return nil, fmt.Errorf("Couldn't match required `%v` in `%v`", e.vtyp, input)
				}
				break
			}
		}

		return m, nil
	}
}

// Parse creates a match object from input with the given Parser format
func (p *Parser) Parse(input string) (*Match, error) {
	return p.expression(input)
}

// SetParameter sets the parameter description for the given type name
func (p *Parser) SetParameter(typeName string, param *Parameter) {
	p.parameter[typeName] = param
}

// SetParameters sets a map of parameter description
func (p *Parser) SetParameters(params ParameterMap) {
	for k, v := range params {
		p.SetParameter(k, v)
	}
}

// Parameter gets the parameter description for the given type name
func (p *Parser) Parameter(typeName string) *Parameter {
	def, ok := p.parameter[typeName]
	if !ok {
		def = NewParameter()
		p.parameter[typeName] = def
	}

	return def
}
