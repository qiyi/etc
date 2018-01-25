package etc

const (
	placeholderPrefix = "${"
	placeholderSuffix = "}"
	valueSeparator    = ":"
)

var _ Resolver = &PropertyResolver{}

// PropertyResolver resolve property configuration
type PropertyResolver struct {
	sources *MutableSourceList
}

// Contains test if contains the key
func (p *PropertyResolver) Contains(key string) bool {
	sources := p.sources.List()
	for _, s := range sources {
		if s.Contains(key) {
			return true
		}
	}
	return false
}

// GetRaw get the origin value of configuration
func (p *PropertyResolver) GetRaw(key string) interface{} {
	for _, source := range p.sources.List() {
		value := source.Get(key)
		if value != nil {
			return value
		}
	}
	return nil
}

// Get return the value
func (p *PropertyResolver) Get(key string) interface{} {
	value := p.GetRaw(key)
	if value != nil {
		if v, ok := value.(string); ok {
			return p.Resolve(v)
		}
		return value
	}
	return nil
}

// GetDefault return the value or the default value if the key does not exist
func (s *PropertyResolver) GetDefault(key string, defaultVal interface{}) interface{} {
	value := s.Get(key)
	if value == nil {
		return defaultVal
	}
	return value
}

// Get return the value
func (p *PropertyResolver) GetStr(key string) string {
	value := p.GetRaw(key)
	if value != nil {
		if v, ok := value.(string); ok {
			return p.Resolve(v)
		}
	}
	return ""
}

// Get return the value
func (p *PropertyResolver) GetStrDefault(key string, defaultVal string) string {
	value := p.GetRaw(key)
	if value != nil {
		if v, ok := value.(string); ok {
			return p.Resolve(v)
		}
	}
	return defaultVal
}

func (s *PropertyResolver) Resolve(text string) string {
	//FIXME to support placeholder resolve
	return text;
}

func (s *PropertyResolver) GetSourceList() *MutableSourceList {
	return s.sources
}
