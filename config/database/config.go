package database

import (
	"fmt"
	"strconv"

	"site/api/config/env"
)

// Config generic database configuration.
type Config struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type ConfigOpt func(*Config) error

func ConfigFromValues(host, name, user, password string, port int) ConfigOpt {
	return func(c *Config) error {
		c.Host = host
		c.Database = name
		c.User = user
		c.Password = password
		c.Port = port
		return nil
	}
}

func ConfigFromEnv(config *Config) error {
	config.Host = env.GetDefault("DB_HOST", "localhost")
	config.Database = env.GetDefault("DB_NAME", "site_db")
	config.User = env.GetDefault("DB_USER", "postgres")
	config.Password = env.GetDefault("DB_PASSWORD", "postgres")
	portEnv := env.GetDefault("DB_PORT", "5432")
	config.Port, _ = strconv.Atoi(portEnv)
	if config.Port <= 0 {
		config.Port = 5432
	}
	return nil
}

// DSN used for database connection.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable port=%d",
		c.Host,
		c.Database,
		c.User,
		c.Password,
		c.Port,
	)
}
