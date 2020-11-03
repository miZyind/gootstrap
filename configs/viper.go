package configs

import "github.com/spf13/viper"

// InitViper ...
func InitViper() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
