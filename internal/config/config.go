package config

import (
	"log"
	"os"
	"strconv"
)

func IsRequiredEnvironments() bool {
	requiredKeys := []string{
		"HOST",
		"APP_ENV",
		"MYSQL_USERNAME",
		"MYSQL_PASSWORD",
		"MYSQL_DATABASE",
	}

	undefinedKeys := getUndefinedEnvironmentKeys(requiredKeys)

	if len(undefinedKeys) != 0 {
		log.Println("config.IsRequiredEnvironments: undefined required environments:")
		for _, key := range undefinedKeys {
			log.Println(key)
		}
		return false
	}
	return true
}

func getUndefinedEnvironmentKeys(keys []string) []string {
	var undefinedKeys []string
	for _, key := range keys {
		if _, ok := os.LookupEnv(key); !ok {
			undefinedKeys = append(undefinedKeys, key)
		}
	}
	return undefinedKeys
}

func Port() string {
	return os.Getenv("PORT")
}

func Host() string {
	return os.Getenv("HOST")
}

func getAppEnv() string {
	return os.Getenv("APP_ENV")
}

func IsDevelopment() bool {
	return getAppEnv() == "development"
}

func IsProduction() bool {
	return getAppEnv() == "production"
}

func ReadTimeout() int {
	timeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Println("config.ReadTimeout: READ_TIMEOUT is not int.")
		return 5
	}
	return timeout
}

func WriteTimeout() int {
	timeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Println("config.ReadTimeout: 'READ_TIMEOUT' is not int.")
		return 5
	}
	return timeout
}

func ShutdownWaitTime() int {
	waitTime, err := strconv.Atoi(os.Getenv("SHUTDOWN_WAIT_TIME"))
	if err != nil {
		log.Println("config.ShutdownWaitTime: 'SHUTDOWN_WAIT_TIME' is not int.")
		return 30
	}
	return waitTime
}
