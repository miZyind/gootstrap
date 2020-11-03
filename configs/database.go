package configs

import "github.com/spf13/viper"

// DatabaseConfig ...
type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// InitDatabaseConfig ...
func InitDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetInt("DB_PORT"),
		Name:     viper.GetString("DB_NAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
	}
}
