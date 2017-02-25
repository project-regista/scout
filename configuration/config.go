package configuration

import (
	"log"

	"github.com/spf13/viper"
)

const LoadConfigErrMsg = "Error loading config file"

// LoadConfig from Viper
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configuration")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s: %s", LoadConfigErrMsg, err)
	}
}
