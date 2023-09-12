package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

type Environment struct {
	ApiPort string `env:"API_PORT"`
	//Aws     struct {
	//	Endpoint        string `env:"AWS_ENDPOINT"`
	//	Region          string `env:"AWS_DEFAULT_REGION"`
	//	AccessKeyId     string `env:"AWS_ACCESS_KEY_ID"`
	//	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	//}
	Database struct {
		//TableName string `env:"TABLE_NAME"`
		Host     string `env:"DATABASE_HOST"`
		User     string `env:"DATABASE_USER"`
		Pass     string `env:"DATABASE_PASS"`
		Name     string `env:"DATABASE_NAME"`
		Port     string `env:"DATABASE_PORT"`
		SslMode  string `env:"DATABASE_SSL_MODE"`
		TimeZone string `env:"DATABASE_TIME_ZONE"`
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
