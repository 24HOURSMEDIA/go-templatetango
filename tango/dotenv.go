package tango

import (
	"github.com/joho/godotenv"
	"os"
)

// LoadDotEnv loads the .env file in the given directory
func LoadDotEnv(dir string) (err error) {
	// check if .env file exists
	if _, err = os.Stat(dir + "/.env"); os.IsNotExist(err) {
		return nil
	}
	return godotenv.Load(dir + "/.env")
}
