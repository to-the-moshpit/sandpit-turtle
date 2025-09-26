package environment

import (
	"log"
	"os"
)

type AppEnv = string

var AppEnvEnum = &struct {
	Local AppEnv
	Prod  AppEnv
}{
	Local: "local",
	Prod:  "prod",
}

func determineAppEnv() AppEnv {
	val := os.Getenv(APP_ENV)
	def := AppEnvEnum.Local

	switch val {
	case string(AppEnvEnum.Local):
		return AppEnvEnum.Local
	case string(AppEnvEnum.Prod):
		return AppEnvEnum.Prod
	default:
		log.Printf("No %v specified defaulting to %v", APP_ENV, string(def))
		return def
	}
}
