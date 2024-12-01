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

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	CoinsServiceHost string
	CoinsGRPCPort    string

	PostgresMaxConnections int32
}

// Load ...
func Load() Config {
	if err := godotenv.Load("/ps_go_coin_service.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "tg_go_coins_service"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "debug"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "mynewproject.cl0oquum2f1m.ap-northeast-1.rds.amazonaws.com"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "Bahodir42434542809i"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "mypm"))
	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	// config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	// config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	// config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "bahodir"))
	// config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1100"))
	// config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "telegram_bot"))
	// config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.CoinsServiceHost = cast.ToString(getOrReturnDefaultValue("COINS_SERVICE_HOST", "localhost"))
	config.CoinsGRPCPort = cast.ToString(getOrReturnDefaultValue("COINS_GRPC_PORT", ":50501"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
