package env_utils

import "os"

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
