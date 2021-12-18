package main

import (
	"os"
)

// Get enviroment variables from enviroment
func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}
