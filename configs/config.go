package configs

import (
	"os"
)

type conf struct {
	DBDriver          string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	WebServerPort     string
	GRPCServerPort    string
	GraphQLServerPort string
	RabbitMQHost      string
	RabbitMQPort      string
	RabbitMQUser      string
	RabbitMQPassword  string
}

func LoadConfig(path string) (*conf, error) {
	cfg := &conf{
		DBDriver:          getEnv("DB_DRIVER", "mysql"),
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "3306"),
		DBUser:            getEnv("DB_USER", "root"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "orders"),
		WebServerPort:     getEnv("WEB_SERVER_PORT", "8080"),
		GRPCServerPort:    getEnv("GRPC_SERVER_PORT", "50051"),
		GraphQLServerPort: getEnv("GRAPHQL_SERVER_PORT", "8081"),
		RabbitMQHost:      getEnv("RABBITMQ_HOST", "localhost"),
		RabbitMQPort:      getEnv("RABBITMQ_PORT", "5672"),
		RabbitMQUser:      getEnv("RABBITMQ_USER", "guest"),
		RabbitMQPassword:  getEnv("RABBITMQ_PASSWORD", "guest"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
