package etc

var _ Environment = &StandardEnvironment{}

type StandardEnvironment struct {
	PropertyResolver
}

// NewEnv create an environment
func NewEnv(p map[string]interface{}) *StandardEnvironment {
	// TODO add command line and os env source
	env := &StandardEnvironment{PropertyResolver{&MutableSourceList{}}}
	env.sources.AddLast(NewMapSource("defaultProperties", p))
	return env
}

func (env *StandardEnvironment) GetActiveProfiles() []string {
	// TODO support profiles
	return []string{"default"}
}

func (env *StandardEnvironment) SetActiveProfiles(profiles []string) {
	// TODO support profiles
}
func (env *StandardEnvironment) AddActiveProfile(profile string) {
	// TODO suuport profiles
}
func (env *StandardEnvironment) Accept(profile string) bool {
	// TODO support profiles
	return true
}
func (env *StandardEnvironment) GetSourceList() SourceList {
	return env.sources
}
