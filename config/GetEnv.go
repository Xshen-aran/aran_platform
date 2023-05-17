package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Env = map[string]string{}
)

const (
	dev  string = ".env.development"
	prod string = ".env.production"
	env  string = ".env"
)

func init() {
	var err error
	for i, v := range os.Args {
		if v == "-m" {
			switch os.Args[i+1] {
			case "dev":
				Env, err = godotenv.Read(dev)

			case "prod":
				Env, err = godotenv.Read(prod)
				Env["GIN_MODE"] = "info"
			default:
				Env, err = godotenv.Read(env)

			}

		}
	}
	if err != nil {
		panic(err)
	}
}
