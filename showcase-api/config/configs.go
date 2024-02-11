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
	configPath := config.getPath()
	yamlFile, err := os.ReadFile(configPath)

	if err != nil {
		panic("file reading failed")
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		panic("config un-marshal failed")
	}
	l.Printf("%v %v", string(yamlFile), config.SConfig)
	return config
}

func (c *Config) Port() string {
	c.l.Printf("port := %s\n", c.SConfig.Port)
	return c.SConfig.Port
}

func (c *Config) getPath() string {
	args := os.Args
	path := fmt.Sprintf("envs/%s.yaml", args[1])
	c.l.Println(path)
	return path
}
