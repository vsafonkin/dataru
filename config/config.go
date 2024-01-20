package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger string   `yaml:"logger"`
	Names  []string `yaml:"names"`
	Server
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var config Config

var defaultConfig = Config{
	Logger: "console",
	Server: Server{
		Host: "localhost",
		Port: 8000,
	},
}

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

func Logger() string {
	if config.Logger == "" {
		fmt.Println("Logger is not set, default value:", defaultConfig.Logger)
		return defaultConfig.Logger
	}
	return config.Logger
}

func Host() string {
	if config.Host == "" {
		fmt.Println("Host is not set, default value:", defaultConfig.Server.Host)
		return defaultConfig.Server.Host
	}
	return config.Host
}

func Port() int {
	if config.Port == 0 {
		fmt.Println("Port is not set, default value:", defaultConfig.Server.Port)
		return defaultConfig.Server.Port
	}
	return config.Port
}

func Names() []string {
	return config.Names
}
