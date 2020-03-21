package env

import (
	"fmt"
	"os"
	"strconv"
)

func RequiredString(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("required env var %s not found", name)
	}

	return value, nil
}

func DefaultString(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}

func RequiredInt(name string) (int, error) {
	value, err := RequiredString(name)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("required env var %s could not be converted to int", name)
	}

	return intValue, nil
}

func DefaultInt(name string, defaultValue int) (int, error) {
	value, err := RequiredString(name)
	if err != nil {
		return defaultValue, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("env var %s could not be converted to int", name)
	}

	return intValue, nil
}
