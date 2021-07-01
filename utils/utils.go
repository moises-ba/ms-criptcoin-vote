package utils

import "os"

func GetEnv(envName, defaultValue string) string {
	if val := os.Getenv(envName); len(val) > 0 {
		return val
	} else {
		return defaultValue
	}

}
