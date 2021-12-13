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
	CacheDir     string `json:"cache_dir" yaml:"cache_dir" toml:"cache_dir"`
	OutputDir    string `json:"output_dir" yaml:"output_dir" toml:"output_dir"`
	Dependencies map[string]Dependency
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o.Apply(c)
	}
	return c
}
