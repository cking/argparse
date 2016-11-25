package argparse

import (
	"regexp"
	"strings"
)

// Parser object
type Parser struct {
	format     string
	parameter  map[string]*Parameter
	expression func(string)
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

var (
	formatRegexp = regexp.MustCompile(`\<(?P<rkey>\w+)(?::(?P<rtype>\w+))?\>|\[(?P<okey>\w+)(?::(?P<otype>\w+))?\]|(?:\\[\[<]|[^\[<])+`)
)

// New create a new Parser object
func New(format string) *Parser {
	p := &Parser{format, map[string]*Parameter{
		"str": NewDefaultParameter(StringMatcher),
		"int": NewDefaultParameter(IntegerMatcher),
	},
		createExpression(format)}
	return p
}

// Format returns the format string
func (p *Parser) Format() string {
	return p.format
}

func createExpression(format string) func(string) {
	matches := formatRegexp.FindAllStringSubmatch(format, -1)
	if matches == nil || len(matches) == 0 {
		return func(input string) {}
	}

	// convert matches into an easier to use version
	expr := make([]*expression, len(matches))
	for _, match := range matches {
		if len(match[1]) > 0 {
			expr = append(expr, &expression{requiredType, match[2], match[1]})
		} else if len(match[3]) > 0 {
			expr = append(expr, &expression{optionalType, match[4], match[3]})
		} else {
			expr = append(expr, &expression{literalType, "", match[0]})
		}
	}

	return func(input string) {

	}
}

// Parse creates a match object from input with the given Parser format
func (p *Parser) Parse(input string) {
	var pos int
	format := p.format

	pos = strings.Index(p.format, "<")
	if pos > 0 {

	}
}

// SetParameter sets the parameter description for the given type name
func (p *Parser) SetParameter(typeName string, param *Parameter) {
	p.parameter[typeName] = param
}

// SetParameters sets a map of parameter description
func (p *Parser) SetParameters(params map[string]*Parameter) {
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
