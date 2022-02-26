package env

import (
	"os"
	"strconv"
)

func GetDefault(envName string, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetDefaultInt(envName string, defaultValue int) int {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	valueInt, _ := strconv.ParseInt(value, 10, 64)
	return int(valueInt)
}

func GetDefaultBool(envName string, defaultValue bool) bool {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	valueBool, _ := strconv.ParseBool(value)
	return valueBool
}
