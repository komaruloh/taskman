package config

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	BaseDir  string   `json:"base_dir" yaml:"base_dir"`
	Database Database `json:"database" yaml:"database"`
}

type Database struct {
	FileNmae string `json:"file_name" yaml:"file_name"`
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	// go:embed config.yaml
	var value []byte
	cfg := &Config{}
	err := yaml.Unmarshal(value, cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return cfg
}
