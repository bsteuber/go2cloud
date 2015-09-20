package env

import (
	"fmt"
	"os"
)

func Get(envVarName string) string {
	envVar := os.Getenv(envVarName)
	if envVar == "" {
		panic(fmt.Sprintf("Environment variable must be set: %v", envVarName))
	}
	return envVar
}
