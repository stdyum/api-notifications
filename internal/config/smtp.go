package config

import (
	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	From     string `env:"FROM"`
}

func SMTP(config SMTPConfig) (*gomail.Dialer, error) {
	return gomail.NewDialer(config.Host, config.Port, config.User, config.Password), nil
}
