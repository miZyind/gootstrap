package configs

import (
	"time"

	"github.com/spf13/viper"
)

// A ServerConfig defines parameters for running an HTTP server
type ServerConfig struct {
	Mode         string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Server represents server config instance
var Server *ServerConfig

func initServerConfig() {
	Server = &ServerConfig{
		Mode:         viper.GetString("SERVER_MODE"),
		Addr:         viper.GetString("SERVER_ADDR"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
	}
}
