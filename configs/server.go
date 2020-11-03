package configs

import (
	"time"

	"github.com/spf13/viper"
)

// ServerConfig ...
type ServerConfig struct {
	Mode         string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// InitServerConfig ...
func InitServerConfig() *ServerConfig {
	return &ServerConfig{
		Mode:         viper.GetString("SERVER_MODE"),
		Addr:         viper.GetString("SERVER_ADDR"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
	}
}
