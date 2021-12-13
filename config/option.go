package config

type Option interface {
	Apply(*Config)
}

type OutputDir string

func (o OutputDir) Apply(c *Config) {
	c.OutputDir = string(o)
}

type CacheDir string

func (o CacheDir) Apply(c *Config) {
	c.CacheDir = string(o)
}
