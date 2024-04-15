package utils

import "os"

func GetSetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	os.Setenv(key, defaultValue)
	return defaultValue
}
