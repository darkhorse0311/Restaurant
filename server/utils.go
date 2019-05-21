package server

import (
	"fmt"
	"os"
)

// CheckEnviroment function makes sure all enviroment variables are set befire running
func CheckEnviroment() error {
	variables := []string{
		"PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
		"DB_PORT",
		"YELP_API",
		"ENVIROMENT",
		"JWT_KEY",
		"CACHE_ADDRS",
		"CACHE_DB",
	}

	for _, v := range variables {
		if value := os.Getenv(v); value == "" {
			return fmt.Errorf("enviroment variable %s is required", v)
		}
	}

	return nil
}

// Check checks if error exists
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
