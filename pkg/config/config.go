package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name     string   `yaml:"name"`
	Listen   string   `yaml:"listen"`
	Env      string   `yaml:"env"`
	Postgres DBConfig `yaml:"postgres"`
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
