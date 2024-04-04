package config

import "fmt"

type (
	Main struct {
		Auth      Auth      `envPrefix:"AUTH_"`
		WebServer WebServer `envPrefix:"WEB_SERVER_"`
		Database  Database  `envPrefix:"DB_"`
	}

	Auth struct {
		Token string `env:"TOKEN"`
	}

	WebServer struct {
		Port int `env:"PORT"`
	}

	Database struct {
		RetryInterval   int    `env:"RETRY_INTERVAL"`
		MaxIdleConn     int    `env:"MAX_IDLE"`
		MaxConn         int    `env:"MAX_CONN"`
		ConnMaxLifetime string `env:"CONN_MAX_LIFETIME"`
		Host            string `env:"HOST"`
		Port            string `env:"PORT"`
		User            string `env:"USER"`
		Password        string `env:"PASSWORD"`
		Name            string `env:"NAME"`
		SSLMode         string `env:"SSL_MODE"`
	}
)

func (d Database) GetDSN() string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
	)
	return dsn
}
