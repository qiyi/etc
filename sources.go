package etc

// MutableSourceList is a mutable source list...
type MutableSourceList struct {
	sources []Source
}

// Contains test if the source is in this list
func (m *MutableSourceList) Contains(name string) bool {
	s := m.Get(name)
	return s != nil
}

// Get return the config source with the specified name
func (m *MutableSourceList) Get(name string) Source {
	for _, s := range m.sources {
		if name == s.Name() {
			return s
		}
	}
	return nil
}

// List return all config source
func (m *MutableSourceList) List() []Source {
	ss := make([]Source, 0)
	for _, s := range m.sources {
		ss = append(ss, s)
	}
	return ss
}

// AddFirst add the source to the first of the source list
func (m *MutableSourceList) AddFirst(s Source) {
	m.sources = append([]Source{s}, m.sources...)
}

// AddLast add the source to the last of the source list
func (m *MutableSourceList) AddLast(s Source) {
	m.sources = append(m.sources, s)
}

func (m *MutableSourceList) index(name string) int {
	if len(m.sources) != 0 {
		for i, s := range m.sources {
			if name == s.Name() {
				return i
			}
		}
	}
	return -1
}

func (m *MutableSourceList) insert(at int, s Source) {
	if at == len(m.sources) {
		m.sources = append(m.sources, s)
		return
	}
	if at >= 0 && at < len(m.sources) {
		sources := append([]Source{}, m.sources[:at]...)
		sources = append(sources, s)
		sources = append(sources, m.sources[at:]...)
		m.sources = sources
	}
}

// AddBefore add the source to the position before the source associate the name
func (m *MutableSourceList) AddBefore(name string, s Source) {
	idx := m.index(name)
	m.insert(idx, s)
}

// AddAfter add the source to the position after the source associate with the name
func (m *MutableSourceList) AddAfter(name string, s Source) {
	idx := m.index(name)
	m.insert(idx+1, s)
}

// Remove remove the source with the name
func (m *MutableSourceList) Remove(name string) {
	idx := m.index(name)
	if idx >= 0 && idx < len(m.sources) {
		sources := append([]Source{}, m.sources[:idx]...)
		sources = append(sources, m.sources[idx+1:]...)
		m.sources = sources
	}
}

// Replace replcae the source associate the name
func (m *MutableSourceList) Replace(name string, s Source) {
	idx := m.index(name)
	if idx >= 0 && idx < len(m.sources) {
		m.sources[idx] = s
	}
}

// Size return the size of source list
func (m *MutableSourceList) Size() int {
	return len(m.sources)
}
