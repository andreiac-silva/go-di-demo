package env

import (
	"os"

	// Automatic load environment variables from .env.
	_ "github.com/joho/godotenv/autoload"
)

var (
	PostgresDSN   string
	ServerAddress string
)

func init() {
	PostgresDSN = os.Getenv("POSTGRES_DSN")
	ServerAddress = os.Getenv("SERVER_ADDRESS")
}
