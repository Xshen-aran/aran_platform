package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Env = map[string]string{}
)

const (
	dev  string = ".env.development"
	prod string = ".env.production"
)

func init() {
	var err error
	for i, v := range os.Args {
		if v == "-m" {
			switch os.Args[i+1] {
			case "dev":
				Env, err = godotenv.Read(dev)
				Env["GIN_MODE"] = "debug"
			case "prod":
				Env, err = godotenv.Read(prod)
				Env["GIN_MODE"] = "info"
			default:
				Env, err = godotenv.Read(dev)
				Env["GIN_MODE"] = "debug"
			}

		}
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Env: %v\n", Env)
}
