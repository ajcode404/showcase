package config

import (
	"fmt"
	"os"
)

type Config struct {
	SConfig *ServerConfig `yaml:"server-config"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func (c *Config) port() string {
	return c.SConfig.Port
}

func (c *Config) host() string {
	return c.SConfig.Host
}

// private API's
func (c *Config) getPath() string {
	env := getEnv()
	path := fmt.Sprintf("envs/%s.yaml", env)
	return path
}

func getEnv() string {
	args := os.Args
	defaultEnv := "prod"
	var env string
	if len(args) >= 2 {
		env = args[1]
		switch env {
		case "dev":
			return env
		case "prod":
			return env
		case "sandbox":
			return env
		}
	}
	return defaultEnv
}
