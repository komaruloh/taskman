package config

import (
	"fmt"
	gap "github.com/muesli/go-app-paths"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

var scope = gap.NewScope(gap.User, "taskman")

func TestNewConfig(t *testing.T) {
	// Test that NewConfig returns a non-nil pointer
	t.Run("Returns non-nil Config", func(t *testing.T) {
		cfg, err := NewConfig(scope)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if cfg == nil {
			t.Error("Expected NewConfig() to return a non-nil Config object")
		}
	})

	// Test default values are set correctly
	t.Run("Sets default values", func(t *testing.T) {
		cfg, err := NewConfig(scope)

		// Add assertions for expected default values
		xCfgByte := []byte(`---
app:
  name: "Taskman"
data_dir: ${HOME}/.local/share/taskman
config_path: ${HOME}/.config/taskman/config.yaml
database:
  filename: "taskman.db"
`)
		xCfg := &Config{}
		err = yaml.Unmarshal(xCfgByte, xCfg)
		if err != nil {
			t.Fatalf("error: %v", err)
		}

		xCfg.DataDir = os.ExpandEnv(xCfg.DataDir)
		xCfg.ConfigPath = os.ExpandEnv(xCfg.ConfigPath)
		if err := compareCfgs(cfg, xCfg); err != nil {
			t.Errorf("error: %v", err)
		}
	})

	// Test loading from environment variables (if applicable)
	t.Run("Loads from environment variables", func(t *testing.T) {
		// Set environment variables before calling NewConfig
		datadirValue := os.ExpandEnv("${HOME}/.local/share/taskman")
		t.Setenv("TASKMAN_DATADIR", datadirValue)

		cfg, err := NewConfig(scope)
		if err != nil {
			t.Fatalf("error: %v", err)
		}

		// Assert that environment variables were properly loaded
		if cfg.DataDir != datadirValue {
			t.Errorf("Expected env value %q for DataDir, got %q", datadirValue, cfg.DataDir)
		}
	})
}

// compare two Config structs, expecting them to be equal
func compareCfgs(cfg1, cfg2 *Config) error {
	if cfg1.App.Name != cfg2.App.Name {
		return fmt.Errorf("Expected default value %q for App.Name, got %q", cfg2.App.Name, cfg1.App.Name)
	}

	if cfg1.App.Version != cfg2.App.Version {
		return fmt.Errorf("Expected default value %q for App.Version, got %q", cfg2.App.Version, cfg1.App.Version)
	}

	if cfg1.Initialized != cfg2.Initialized {
		return fmt.Errorf("Expected default value %t for Initialized, got %t", cfg2.Initialized, cfg1.Initialized)
	}

	if cfg1.Database.FileName != cfg2.Database.FileName {
		return fmt.Errorf("Expected default value %q for Database.FileName, got %q", cfg2.Database.FileName, cfg1.Database.FileName)
	}

	if cfg1.DataDir != cfg2.DataDir {
		return fmt.Errorf("Expected default value %s for Datadir, got %s", cfg2.DataDir, cfg1.DataDir)
	}

	if cfg1.Database.FileName != cfg2.Database.FileName {
		return fmt.Errorf("Expected default value %q for Databse.FileName, got %q", cfg2.Database.FileName, cfg1.Database.FileName)
	}

	return nil
}
