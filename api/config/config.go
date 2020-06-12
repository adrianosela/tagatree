package config

import (
	"os"
)

// Conf carries the api configuration
type Conf struct {
	Version        string
	MongoDBConnStr string
	MongoDBName    string
}

// FromEnv loads configuration
// from the environment
func FromEnv(version string) Conf {
	c := Conf{
		Version:        version,
		MongoDBConnStr: os.Getenv("MONGO_CONN_STR"),
		MongoDBName:    os.Getenv("MONGO_DB_NAME"),
	}
	return c
}
