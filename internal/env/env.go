package env

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load("../.env")
	if err != nil{
		fmt.Printf("Error with loading .env file %v\n" , err)
	    return
	}
}

func GetEnv(key string) string {
	 return os.Getenv(key)
}