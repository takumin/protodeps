package config

type Dependency struct {
	Source   string
	Branch   string
	Commit   string
	Tag      string
	Licenses []string
	Includes []string
	Excludes []string
}

type Config struct {
	CacheDir     string
	OutputDir    string
	Dependencies []Dependency
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o.Apply(c)
	}
	return c
}
