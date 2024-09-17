package configs

import (
	"github.com/ggorockee/toolbox"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func New() *Config {
	return &Config{
		DBHost:     toolbox.Getenv("MYSQL_HOST", "localhost"),
		DBPort:     toolbox.Getenv("MYSQL_PORT", "3306"),
		DBUser:     toolbox.Getenv("MYSQL_USER", "ggorockee"),
		DBPassword: toolbox.Getenv("MYSQL_PASSWORD", "ggorockee"),
		DBName:     toolbox.Getenv("MYSQL_DATABASE", "backend"),
		JWTSecret:  toolbox.Getenv("JWT_SECRET", "ggorockee"),
	}
}
