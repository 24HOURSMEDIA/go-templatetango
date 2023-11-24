package tango

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadDotEnv(path string) (err error) {
	// check if .env file exists
	if _, err = os.Stat(path + "/.env"); os.IsNotExist(err) {
		return nil
	}
	return godotenv.Load(path + "/.env")
}
