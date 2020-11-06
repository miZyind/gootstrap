package configs

import (
	"github.com/mizyind/gootstrap/utils/logger"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")

	if viper.ReadInConfig() == nil {
		logger.Bootstrapper("Env loaded from .env")
	}

	initServerConfig()
	logger.Config(Server)

	initDatabaseConfig()
	logger.Config(Database)
}
