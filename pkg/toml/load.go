package toml

import (
	"github.com/pelletier/go-toml"
)

// Rule represents a single regex rule from the TOML file.
type Rule struct {
	ID          string   `toml:"id"`
	Description string   `toml:"description"`
	Regex       string   `toml:"regex"`
	Keywords    []string `toml:"keywords"`
}

// Config holds the parsed TOML configuration.
type Config struct {
	Rules []Rule `toml:"rules"`
}

// LoadConfig loads and parses the TOML file.
func LoadConfig(filename string) (*Config, error) {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = tree.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}