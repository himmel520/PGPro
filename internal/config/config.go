package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Password string
}

// New creates a new Config by reading the config file at the path
// and retrieving the db pwd from the PG_PASSWORD env.
func New(cfgPath string) (*Config, error) {
	file, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %v", err)
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(file, cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration file: %v", err)
	}

	if err := cfg.GetPassword(); err != nil {
		return nil, fmt.Errorf("failed to get password from env: %s", err)
	}

	return cfg, nil
}

// GetPassword retrieves the db pwd from the PG_PASSWORD env.
// Returns an error if not set.
func (cfg *Config) GetPassword() error {
	pwd, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		return errors.New("PG_PASSWORD environment variable is not set")
	}
	cfg.DB.Password = pwd
	return nil
}

// GetDatabaseUrl constructs and returns the database connection string based on the DB configuration.
func (cfg *Config) GetDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, cfg.DB.SSLMode, cfg.DB.Password,
	)
}
