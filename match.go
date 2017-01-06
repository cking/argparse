package argparse

// Match holds the parameter results of a string parsing operation
type Match struct {
	matches map[string]interface{}
}

// NewMatch creates a new empty Match struct
func NewMatch() *Match {
	m := &Match{}
	m.matches = make(map[string]interface{})
	return m
}

func (m *Match) addParameter(key string, value interface{}) {
	m.matches[key] = value
}

// HasMatch checks if a requested parameter has a value
func (m *Match) HasMatch(key string) bool {
	return m.GetMatch(key) != nil
}

// GetMatch returns the result of a parameter
func (m *Match) GetMatch(key string) interface{} {
	return m.matches[key]
}

// GetInteger returns the result of a parameter as an integer
func (m *Match) GetInteger(key string) int {
	v, ok := m.GetMatch(key).(int)
	if ok {
		return v
	}
	return -1
}

// GetString returns the result of a parameter as a string
func (m *Match) GetString(key string) string {
	v, ok := m.GetMatch(key).(string)
	if ok {
		return v
	}
	return ""
}
