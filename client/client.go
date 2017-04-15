// Package client implements an HTTP client to get data from Soccerama API
// It starts by loading the necessary configuration and then initializes
// the Scout client, making it available for Regista's components to use it.
package client

import (
	"fmt"

	"github.com/spf13/viper"
)

// Soccerama wraps configuration values for Soccerama client
type Soccerama struct {
	APIToken   string
	APIVersion string
	APIHost    string
}

// Client wraps Scout client
type Client struct{ *Soccerama }

// New returns a Client value(!)
func New(filepath string) (Client, error) {

	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		return Client{}, fmt.Errorf("Error loading config file: %s", err)
	}

	// Load soccerama config values
	soccerama, err := loadSoccerama()
	if err != nil {
		return Client{}, fmt.Errorf("Failed to load configurations: %s", err)
	}

	// We don't need Viper anymore since the config values are already loaded
	viper.Reset()

	return Client{Soccerama: soccerama}, nil
}

// loadSoccerama checks if configuration values are available
// and initialize soccerama struct with these values
func loadSoccerama() (*Soccerama, error) {

	keys := []string{
		"api_token",
		"api_version",
		"api_host",
	}

	for _, v := range keys {
		if !viper.IsSet(fmt.Sprintf("soccerama.%s", v)) {
			return nil, fmt.Errorf("Could not load: %s", v)
		}
	}

	return &Soccerama{
		APIToken:   viper.GetString("soccerama.api_token"),
		APIVersion: viper.GetString("soccerama.api_version"),
		APIHost:    viper.GetString("soccerama.api_host"),
	}, nil
}

// Implement the Stringer interface
func (c Client) String() string {
	return fmt.Sprintf("api token: %s\napi_version: %s\napi_host: %s\n",
		c.APIToken, c.APIVersion, c.APIHost)
}
