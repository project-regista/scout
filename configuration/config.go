package configuration

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const LoadConfigErrMsg = "Error loading config file"

type Auth struct {
	user     string
	password string
	host     string
	port     string
	apitkn   string
	URL      string
	Key      string
}

// LoadConfig from Viper
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configuration")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s: %s", LoadConfigErrMsg, err)
	}
}

// APIAuth builds the values used for socceramaAPI access
func APIAuth() *Auth {
	return &Auth{
		apitkn: viper.GetString("soccerama.apitkn"),
	}
}

// GetAPI returns the API key
func (a *Auth) GetAPI() {
	a.Key = fmt.Sprintf(a.apitkn)
}

// DbAuth builds values for Db access
func DbAuth() *Auth {
	return &Auth{
		user:     viper.GetString("database.user"),
		password: viper.GetString("database.password"),
		host:     viper.GetString("database.host"),
		port:     viper.GetString("database.port"),
	}
}

// GetDB is used to get Db credentials
func (a *Auth) GetDB() {
	a.URL = fmt.Sprintf("bolt://%s:%s@%s:%s/db/data",
		a.user, a.password, a.host, a.port)
}
