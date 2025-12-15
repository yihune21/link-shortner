package env

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Printf("Error with loading .env file %v" , err)
	    return
	}
}

func GetEnv() (string,string) {
	 port := os.Getenv("PORT")
	 db_url := os.Getenv("DB_URL")
	 return port , db_url
}