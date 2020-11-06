package configs

import "github.com/spf13/viper"

// LoadEnv ...
func LoadEnv() bool {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")

	return viper.ReadInConfig() == nil
}
