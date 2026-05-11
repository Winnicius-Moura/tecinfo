package config

import (
	"os"
	"strconv"

	"github.com/wnn-dev/contributions-analysis/config/flags"
)

// LoadEnvFile loads configuration from environment variables
func LoadEnvFile() (*flags.Configuration, error) {
	var appConf flags.Configuration

	appConf.Postgres.Database = os.Getenv("DB_NAME")
	appConf.Postgres.Driver = os.Getenv("DB_DRIVER")
	appConf.Postgres.Username = os.Getenv("DB_USERNAME")
	appConf.Postgres.Password = os.Getenv("DB_PASSWORD")
	appConf.Postgres.Host = os.Getenv("DB_HOST")
	appConf.Postgres.Port = os.Getenv("DB_PORT")
	appConf.Address = os.Getenv("ADDRESS")

	appConf.JWT.Secret = os.Getenv("JWT_SECRET")
	appConf.SMTP.Host = os.Getenv("SMTP_HOST")
	appConf.SMTP.Username = os.Getenv("SMTP_USER")
	appConf.SMTP.Password = os.Getenv("SMTP_PASS")
	
	// Port fallback to 587 if not valid
	port := 587
	if p := os.Getenv("SMTP_PORT"); p != "" {
		if val, err := strconv.Atoi(p); err == nil {
			port = val
		}
	}
	appConf.SMTP.Port = port

	return &appConf, nil
}
