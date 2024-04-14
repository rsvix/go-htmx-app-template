package utils

import "os"

func GetSetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	os.Setenv(key, fallback)
	return fallback
}
