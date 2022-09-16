package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	BookServiceHost string
	BookServicePort int

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println("No .env file found")
	// }

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	config.BookServiceHost = cast.ToString(getOrReturnDefault("BOOK_SERVICE_HOST", "localhost"))
	config.BookServicePort = cast.ToInt(getOrReturnDefault("BOOK_SERVICE_PORT", 9100))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
