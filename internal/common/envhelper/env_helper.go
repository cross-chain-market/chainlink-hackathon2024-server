package envhelper

import (
	"github.com/joho/godotenv"
	"log/slog"
)

var myEnv map[string]string

const envLoc = ".env"

func loadEnv() {
	var err error
	if myEnv, err = godotenv.Read(envLoc); err != nil {
		slog.Error("could not load env from %s", envLoc, slog.String("error", err.Error()))
	}
}

func GetEnv(key string) string {
	loadEnv()
	return myEnv[key]
}

func SetEnv(key, value string) {
	loadEnv()
	myEnv[key] = value
}
