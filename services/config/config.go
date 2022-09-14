package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	DefaultOffset string
	DefaultLimit  string

	SecretKey string

	PasscodePool   string
	PasscodeLength int

	ServiceHost string
	GRPCPort    string

	PostgresMaxConnections int
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "author_service"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "your_db_password"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.ServiceName))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "cMhXjcrfFRDYrF6f"))

	config.PasscodePool = cast.ToString(getOrReturnDefaultValue("PASSCODE_POOL", "0123456789"))
	config.PasscodeLength = cast.ToInt(getOrReturnDefaultValue("PASSCODE_LENGTH", "6"))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "0.0.0.0"))
	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":9100"))

	config.PostgresMaxConnections = cast.ToInt(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 5))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
