package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/pelletier/go-toml"
)

func (c *Config) LoadIfExist(files ...string) error {
	var f string

	for _, v := range files {
		if v == "" {
			continue
		}
		if _, err := os.Stat(v); err == nil {
			f = v
			break
		}
	}

	if f == "" {
		return errors.New("Required Config File: .protodeps.[json|ya?ml|toml]")
	}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	switch filepath.Ext(f) {
	case ".json":
		if err := json.Unmarshal(b, c); err != nil {
			return err
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(b, c); err != nil {
			return err
		}
	case ".toml":
		if err := toml.Unmarshal(b, c); err != nil {
			return err
		}
	default:
		return errors.New("Required Config File Type: [json|ya?ml|toml]")
	}

	return nil
}
