package config

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration struct have all fields from configuration JSON file
// Configuration struct have all fields from configuration JSON file
type Configuration struct {
	Server Server `json:"server"`
}

// Server defines http server configuration (address and port)
type Server struct {
	IPAddress string `json:"ip_address"`
	Port      string `json:"port"`
}

// Load loads configuration from path
func Load(path string) (Configuration, error) {
	var c Configuration
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(configJSON, &c)
	return c, err
}
