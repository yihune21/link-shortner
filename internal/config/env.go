package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

func LoadEnv() error  {
	err := godotenv.Load(".env")
	if err != nil{
	    return err
	}
	return nil
}

func GetEnv(key string) string {
	 return os.Getenv(key)
}