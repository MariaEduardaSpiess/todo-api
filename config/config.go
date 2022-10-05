package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

type Environment struct {
	ApiPort string `env:"API_PORT"`
	Aws     struct {
		Endpoint        string `env:"AWS_ENDPOINT"`
		Region          string `env:"AWS_DEFAULT_REGION"`
		AccessKeyId     string `env:"AWS_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	}
	Database struct {
		TableName string `env:"TABLE_NAME"`
	}
}

var Env Environment

func LoadEnv() {
	gotenv.Load()

	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		fmt.Print(err.Error())
	}
}
