package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type IConfig struct {
	l        *log.Logger
	s_Config *Config
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
	return &IConfig{l, config}
}

// public API's
func (c *IConfig) Port() string {
	return c.s_Config.port()
}

func (c *IConfig) Host() string {
	return c.s_Config.host()
}
