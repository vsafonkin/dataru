package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger string
	Server
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var config Config

func LoadConfig(path string) error {
	return yamlConfig(path)
}

func yamlConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open config path error: %w", err)
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	if err := d.Decode(&config); err != nil {
		return fmt.Errorf("decode yaml config error: %w", err)
	}

	return nil
}

func LoadValue(name string) error {
	// ToDo
	return nil
}

func Logger() (string, error) {
	if config.Logger == "" {
		return "", fmt.Errorf("logger has a zero value")
	}
	return config.Logger, nil
}

func Host() (string, error) {
	if config.Host == "" {
		return "", fmt.Errorf("host has a zero value")
	}
	return config.Host, nil
}

func Port() (string, error) {
	if config.Port == "" {
		return "", fmt.Errorf("port has a zero value")
	}
	return config.Port, nil
}
