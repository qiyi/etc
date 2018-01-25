package etc

// MapSource is the properties source
type MapSource struct {
	name   string
	source map[string]interface{}
}

// NewMapSource create map based source
func NewMapSource(name string, m map[string]interface{}) *MapSource {
	return &MapSource{name: name, source: m}
}

// Name return the MapSource's name
func (s *MapSource) Name() string {
	return s.name
}

// Keys return ths MapSource's configuration keys
func (s *MapSource) Keys() []string {
	keys := make([]string, 0)
	for key := range s.source {
		keys = append(keys, key)
	}
	return keys
}

// Contains test if the key exist in the MapSource
func (s *MapSource) Contains(name string) bool {
	_, ok := s.source[name]
	return ok
}

// Get return the associate value with the key in this MapSource
func (s *MapSource) Get(name string) interface{} {
	v, ok := s.source[name]
	if ok {
		return v
	}
	return nil
}
