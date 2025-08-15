package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		PidFile string `yaml:"pid_file"`
		LogFile string `yaml:"log_file"`
	} `yaml:"server"`

	Parser struct {
		BangPrefix     string `yaml:"bang_prefix"`
		ShortcutPrefix string `yaml:"shortcut_prefix"`
		SessionPrefix  string `yaml:"session_prefix"`
	} `yaml:"parser"`

	Store struct {
		DBPath string `yaml:"db_path"`
	} `yaml:"store"`

	Service struct {
		KeepTrack       bool   `yaml:"keep_track"`
		DefaultProvider string `yaml:"default_provider"`
	} `yaml:"service"`
}

func NewDefaultAppConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Warning: Could not get user home directory: %v. Using /tmp paths.", err)
		homeDir = "/tmp"
	}

	return &Config{
		Server: struct {
			Port    string `yaml:"port"`
			PidFile string `yaml:"pid_file"`
			LogFile string `yaml:"log_file"`
		}{
			Port:    ":5980",
			PidFile: filepath.Join(homeDir, ".michi", "michi.pid"),
			LogFile: filepath.Join(homeDir, ".michi", "michi.log"),
		},
		Parser: struct {
			BangPrefix     string `yaml:"bang_prefix"`
			ShortcutPrefix string `yaml:"shortcut_prefix"`
			SessionPrefix  string `yaml:"session_prefix"`
		}{
			BangPrefix:     "!",
			ShortcutPrefix: "@",
			SessionPrefix:  "#",
		},
		Store: struct {
			DBPath string `yaml:"db_path"`
		}{
			DBPath: filepath.Join(homeDir, ".michi", "index.db"),
		},
		Service: struct {
			KeepTrack       bool   `yaml:"keep_track"`
			DefaultProvider string `yaml:"default_provider"`
		}{
			KeepTrack:       true,
			DefaultProvider: "g",
		},
	}
}

func LoadConfig(configFilePath string) (*Config, error) {
	cfg := NewDefaultAppConfig()

	if configFilePath == "" {
		log.Println("No specific config file path provided or found in default locations. Using default configuration.")
		return cfg, nil
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Config file not found at '%s'. Using default configuration.", configFilePath)
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to read config file '%s': %w", configFilePath, err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file '%s': %w", configFilePath, err)
	}

	return cfg, nil
}

func ExpandTilde(path string) string {
	if len(path) > 0 && path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Printf("Warning: Could not get user home directory to expand '%s': %v", path, err)
			return path
		}
		return filepath.Join(homeDir, path[1:])
	}
	return path
}

func EnsureConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDirPath := filepath.Join(homeDir, ".michi")

	if err := os.MkdirAll(configDirPath, 0o755); err != nil {
		return "", err
	}
	return configDirPath, nil
}
