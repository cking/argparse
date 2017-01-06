package argparse

type Match struct {
	matches map[string]interface{}
}

func NewMatch() *Match {
	m := &Match{}
	m.matches = make(map[string]interface{})
	return m
}

func (m *Match) addParameter(key string, value interface{}) {
	m.matches[key] = value
}
