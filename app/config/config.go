package config

import (
	"fmt"
	"os"
)

// Service configuration
type Config struct {
	Mongo        Mongo
	Server       Server
	STATUS_CODES map[string]int
}

type Mongo struct {
	Hostname     string `default:"localhost" envconfig:"MONGODB_HOSTNAME"`
	Port         string `default:"27017" envconfig:"MONGODB_PORT"`
	DatabaseName string `default:"redirect" envconfig:"REDIRECT_DATABASE"`
}

type Server struct {
	Port string `default:":8080" envconfig:"SERVER_PORT"`
}

func (m *Mongo) Url() string {
	if os.Getenv("MONGO_URI") != "" {
		return os.Getenv("MONGO_URI")
	}
	return fmt.Sprintf("mongodb://%s:%s", m.Hostname, m.Port)
}

// configuration constructor
func New() *Config {

	STATUS_CODES := map[string]int{
		"REDIRECT":              302,
		"MOVED_PERMANENTLY":     301,
		"BAD_REQUEST":           400,
		"NOT_FOUND":             404,
		"INTERNAL_SERVER_ERROR": 500,
		"OK":                    200,
	}

	return &Config{
		Mongo:        Mongo{Hostname: "localhost", Port: "27017", DatabaseName: "redirect"},
		Server:       Server{Port: ":8080"},
		STATUS_CODES: STATUS_CODES,
	}
}
