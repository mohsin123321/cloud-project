package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Config Configuration

type Configuration struct {
	Database DBconfig
	// Server port of the server
	Server ServerConfig
}

type DBconfig struct {
	// Type of the database to use. Actually, only mysql is supported
	DbType string
	// Address(or hostname) of the database
	DbAddr string
	// Port of the database
	DbPort string
	// User of the database to use
	DbUser string
	// Password of the database to use
	DbPass string
}

type ServerConfig struct {
	// The port to listen
	Port string
}

func ReadConfig() {

	// read the config.json
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config.json")
	}

	if err = json.Unmarshal(raw, &Config); err != nil {
		log.Fatal("Unable to parse configuration file...", err)
	}

}
