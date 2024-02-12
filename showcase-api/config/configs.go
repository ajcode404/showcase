package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type IConfig struct {
	s_Config *Config
}

type Config struct {
	SConfig *ServerConfig `yaml:"server-config"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func NewConfig(l *log.Logger) *IConfig {
	config := &Config{&ServerConfig{}}
	yamlFile, err := os.ReadFile(config.getPath())
	if err != nil {
		panic("file reading failed")
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		panic("config un-marshal failed")
	}
	return &IConfig{config}
}

func (c *IConfig) Port() string {
	return c.s_Config.port()
}

func (c *Config) port() string {
	return c.SConfig.Port
}

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
