package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env      string    `yaml:"env"`
	Name     string    `yaml:"name"`
	Listen   string    `yaml:"listen"`
	Postgres DBConfig  `yaml:"postgres"`
	JWT      JWTConfig `yaml:"jwt"`
}

func New(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

type DBConfig struct {
	Host     string    `yaml:"host"`
	User     string    `yaml:"user"`
	Port     int       `yaml:"port"`
	Password string    `yaml:"password"`
	DBname   string    `yaml:"dbName"`
	Options  DBOptions `yaml:"options"`
}

type DBOptions struct {
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

func (c *DBOptions) UnmarshalYAML(value *yaml.Node) error {
	var raw struct {
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifetime string `yaml:"connMaxLifetime"`
	}

	if err := value.Decode(&raw); err != nil {
		return err
	}

	c.MaxIdleConns = raw.MaxIdleConns
	c.MaxOpenConns = raw.MaxOpenConns

	duration, err := time.ParseDuration(raw.ConnMaxLifetime)
	if err != nil {
		return err
	}
	c.ConnMaxLifetime = duration

	return nil
}

type JWTConfig struct {
	Secret   []byte        `yaml:"secret"`
	Duration time.Duration `yaml:"duration"`
}

func (c *JWTConfig) UnmarshalYAML(value *yaml.Node) error {
	var raw struct {
		Secret   string `yaml:"secret"`
		Duration string `yaml:"duration"`
	}

	if err := value.Decode(&raw); err != nil {
		return err
	}

	duration, err := time.ParseDuration(raw.Duration)
	if err != nil {
		return err
	}

	c.Secret = []byte(raw.Secret)
	c.Duration = duration

	return nil
}
