package environment

import (
	"os"
	"strconv"
)

type OsEnv = string

const (
	PORT        OsEnv = "PORT"
	APP_ENV     OsEnv = "APP_ENV"
	LOG_LEVEL   OsEnv = "LOG_LEVEL"
	DB_HOST     OsEnv = "DB_HOST"
	DB_PORT     OsEnv = "DB_PORT"
	DB_DATABASE OsEnv = "DB_DATABASE"
	DB_USERNAME OsEnv = "DB_USERNAME"
	DB_PASSWORD OsEnv = "DB_PASSWORD"
	DB_SCHEMA   OsEnv = "DB_SCHEMA"
)

type DbEnv struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Schema   string
}

type Env struct {
	Port     int
	AppEnv   AppEnv
	LogLevel string
	Db       *DbEnv
}

var envInstance *Env

func New() *Env {
	if envInstance == nil {
		Refresh()
	}

	return envInstance
}

func Refresh() {
	envInstance = &Env{
		Port:     getIntValue(PORT),
		AppEnv:   determineAppEnv(),
		LogLevel: getStringValue(LOG_LEVEL),
		Db: &DbEnv{
			Host:     getStringValue(DB_HOST),
			Port:     getStringValue(DB_PORT),
			Database: getStringValue(DB_DATABASE),
			Username: getStringValue(DB_USERNAME),
			Password: getStringValue(DB_PASSWORD),
			Schema:   getStringValue(DB_SCHEMA),
		},
	}
}

func getStringValue(key OsEnv) string {
	val := os.Getenv(key)

	return val
}

func getIntValue(key OsEnv) int {
	e := os.Getenv(key)
	value, err := strconv.Atoi(e)
	if err != nil {
		panic(err)
	}

	return value
}
