package configs

import "github.com/spf13/viper"

// A DatabaseConfig defines parameters for connecting to database
type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// Database represents server config instance
var Database *DatabaseConfig

func initDatabaseConfig() {
	Database = &DatabaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetInt("DB_PORT"),
		Name:     viper.GetString("DB_NAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
	}
}
