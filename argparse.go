package argparse

// Parser object
type Parser struct {
	format    string
	parameter map[string]*Parameter
}

// New create a new Parser object
func New(format string) *Parser {
	return &Parser{format: format}
}

// Format returns the format string
func (p *Parser) Format() string {
	return p.format
}

// Parse creates a match object from input with the given Parser format
func (p *Parser) Parse(input string) {

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
