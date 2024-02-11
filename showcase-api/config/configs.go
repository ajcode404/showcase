package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	l       *log.Logger
	SConfig *ServerConfig `yaml:"server-config"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func NewConfig(l *log.Logger) *Config {
	config := &Config{l, &ServerConfig{}}
	yamlFile, err := os.ReadFile(config.getPath())
	if err != nil {
		panic("file reading failed")
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		panic("config un-marshal failed")
	}
	return config
}

func (c *Config) Port() string {
	return c.SConfig.Port
}

func (c *Config) getPath() string {
	env := getEnv()
	path := fmt.Sprintf("envs/%s.yaml", env)
	return path
}

func getEnv() string {
	args := os.Args
	var env string
	if len(args) < 2 {
		env = "prod"
	} else {
		env = args[1]
	}
	return env
}
