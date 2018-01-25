package etc

const (
	activeProfilesName = "etc.profiles.active"
)

// Resolver is the config resolver
type Resolver interface {
	Contains(key string) bool
	Get(key string) interface{}
	GetDefault(key string, defaultVal interface{}) interface{}
	GetStr(key string) string
	GetStrDefault(key string, defaultVal string) string
	Resolve(text string) string
}

// Environment is the resolver with profile support
type Environment interface {
	Resolver
	GetActiveProfiles() []string
	SetActiveProfiles(profiles []string)
	AddActiveProfile(profile string)
	Accept(profile string) bool
	GetSourceList() SourceList
}

// Source is the config source
type Source interface {
	Name() string
	Keys() []string
	Contains(name string) bool
	Get(name string) interface{}
}

// SourceList manage the config sources
type SourceList interface {
	Contains(name string) bool
	Get(name string) Source
	List() []Source
}
