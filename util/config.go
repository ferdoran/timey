package util

import (
	"os"
)

func GetEnvWithDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}
