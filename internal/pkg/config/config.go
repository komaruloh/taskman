package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"

	gap "github.com/muesli/go-app-paths"
	"github.com/spf13/viper"
)

type Config struct {
	App         App `json:"app" yaml:"app"`
	Initialized bool
	ConfigPath  string   `json:"config_path" yaml:"config_path"`
	DataDir     string   `json:"data_dir" yaml:"data_dir"`
	Database    Database `json:"database" yaml:"database"`
}

type App struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
}

type Database struct {
	FileName string `json:"filename" yaml:"filename"`
}

//go:embed config.yaml
var defaultConfig []byte

// NewConfig setuo viper and returns a new Config struct
// App may access the config file using both Config struct of viper.
// Choose one of them to access the config file.
func NewConfig(scope *gap.Scope) (*Config, error) {
	cfg := &Config{}

	if err := cfg.getDefaultValues(); err != nil {
		// Log the error but don't return it
		log.Printf("unable to load default config: %v", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("unable to unmarshal default config: %v", err)
	}

	if err := cfg.setupViper(scope); err != nil {
		log.Fatalf("error setting up viper: %v", err)
	}

	if err := cfg.readConfig(); err != nil {
		// Log the error and set Initialized to false
		log.Printf("error reading config file: %v", err)
		cfg.Initialized = false
	}

	// attempt to write config file
	if !cfg.IsInitialized() {
		// Log the error but don't return it
		log.Printf("config is not initialized")

		dataDirs, err := scope.DataDirs()
		if err != nil {
			log.Fatalf("error getting data dirs: %v", err)
		}
		log.Printf("cool %s", dataDirs[0])

		// dump default config to config.yaml
		if err := viper.WriteConfigAs(cfg.ConfigPath); err != nil {
			log.Fatalf("error writing config file: %v", err)
		}
	}

	return cfg, nil
}

func (cfg *Config) getDefaultValues() error {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return fmt.Errorf("error reading default config: %v", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return fmt.Errorf("error unmarshaling default config: %v", err)
	}

	cfg.ConfigPath = os.ExpandEnv(viper.GetString("config_path"))
	cfg.DataDir = os.ExpandEnv(viper.GetString("data_dir"))
	log.Printf("viper data_dir: %s", viper.GetString("data_dir"))
	log.Printf("defaultConfig: %s", defaultConfig)

	return nil
}

// IsInitialized check if config values are from default config.
// If so, it means the config file is not initialized.
func (cfg *Config) IsInitialized() bool {
	return cfg.Initialized
}

func (cfg *Config) setupViper(scope *gap.Scope) error {
	if scope == nil {
		return fmt.Errorf("scope is nil")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	configDirs, err := scope.ConfigDirs()
	if err != nil {
		return fmt.Errorf("error getting config dirs: %v", err)
	}

	for _, configDir := range configDirs {
		viper.AddConfigPath(configDir)
	}
	cfg.ConfigPath = configDirs[0]

	viper.AutomaticEnv()

	return nil
}

// readConfig reads the config file and unmarshals it into the Config struct
func (cfg *Config) readConfig() error {
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return fmt.Errorf("error unmarshaling config: %v", err)
	}

	return nil
}
