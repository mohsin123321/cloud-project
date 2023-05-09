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
	// Token contains some configuration parameters for issued JWTs.
	Token TokenConfig
	// ShowDocs use to enable the documentation of the api
	ShowDocs bool

	// RateLimiter contains configuration used to limit the http requests
	// over certain time interval per IP or token.
	RateLimiter RateLimiterConfig
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
	// Name of the database to use
	DbName string
}

// TokenConfig contains token information
type TokenConfig struct {
	Secret string
}

type ServerConfig struct {
	// The port to listen
	Port string
}

type RateLimiterConfig struct {
	// SecondsWindow is the duration of the time window in which the max n of requests is counted (in seconds,
	// for the rate limiter).
	SecondsWindow int

	// MaxReqPerIP is the max number of requests that a single user (by IP) can make in the time interval defined
	// in SecondsWindow.
	MaxReqPerIP int
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
