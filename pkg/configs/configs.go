package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   *ServerConfigs
	Database *DatabaseConfigs
	Auth     *AuthConfigs
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Server: &ServerConfigs{
			PortAPI: getEnv("PORT_API", "8080"),
			AllowOrigins: []string{
				"*",
			},
			AllowMethods: []string{
				"GET", "POST", "PUT", "DELETE", "OPTIONS",
			},
			AllowHeaders: []string{
				"Origin", "Content-Type", "Accept", "Authorization", "X-Www-Form-Urlencoded",
			},
			TimeZone:   getEnv("LOG_TIMEZONE", "Asia/Bangkok"),
			TimeFormat: getEnv("LOG_TIME_FORMAT", "2006-01-02 15:04:05"),
			Format:     getEnv("LOG_FORMAT", `${ip} - [${time}] "${method} ${url} ${protocol}" ${status} ${latency}`),
		},
		Database: &DatabaseConfigs{
			Host:                  getEnv("DB_HOST", "localhost"),
			Port:                  getEnv("DB_PORT", "3306"),
			User:                  getEnv("DB_USER", "root"),
			Password:              getEnv("DB_PASSWORD", "root"),
			Name:                  getEnv("DB_NAME", "test"),
			ConnMaxIdleTime:       getEnvAsInt("DB_CONN_MAX_IDLE_TIME", 0),
			ConnectionMaxLifeTime: getEnvAsInt("DB_CONNECTION_MAX_LIFE_TIME", 0),
			MaxIdleConns:          getEnvAsInt("DB_MAX_IDLE_CONNS", 0),
			MaxOpenConns:          getEnvAsInt("DB_MAX_OPEN_CONNS", 0),
		},
		Auth: &AuthConfigs{
			SecretKey:   getEnv("SECRET_KEY", "adsadadd"),
			TokenExpire: getEnvAsInt("TOKEN_EXPIRE", 72),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// func getEnvAsBool(key string, fallback bool) bool {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value == "true" || value == "1"
// 	}
// 	return fallback
// }

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
