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
}

// LoadConfig from Viper
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configuration")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s: %s", LoadConfigErrMsg, err)
	}
}

func ApiAuth() *Auth {
	return &Auth{
		apitkn: viper.GetString("soccerama.apitkn"),
	}
}

func (a *Auth) GetAPI() {
	a.URL = fmt.Sprintf("https://api.soccerama.pro/v1.2/competitions?api_token=%s&include=country",
		a.apitkn)
}

func DbAuth() *Auth {
	return &Auth{
		user:     viper.GetString("database.user"),
		password: viper.GetString("database.password"),
		host:     viper.GetString("database.host"),
		port:     viper.GetString("database.port"),
	}
}

func (a *Auth) GetDB() {
	a.URL = fmt.Sprintf("bolt://%s:%s@%s:%s/db/data",
		a.user, a.password, a.host, a.port)
}
