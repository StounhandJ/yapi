package env

import (
	"os"
)

type Env struct {
	HttpHost  string
	GlagolUrl string
}

var Config Env

func init() {
	Config.HttpHost = os.Getenv("HTTP_HOST")
	if Config.HttpHost == "" {
		Config.HttpHost = ":8001"
	}
	Config.GlagolUrl = os.Getenv("GLAGOL_URL")
}
