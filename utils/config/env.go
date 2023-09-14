package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Port       string
		Mode       string
		AppName    string
		AppVersion string
		Timezone   string
	}

	PubSub struct {
		ProjectId string
		TopicName string
	}
}

var AppCfg AppConfig

func init() {

	var config AppConfig

	config.App.Port = GetEnv("PORT", "8000")
	config.App.Mode = GetEnv("MODE", "Local")
	config.App.AppName = GetEnv("APP_NAME", "Newsletter Publisher")
	config.App.AppVersion = GetEnv("APP_VERSION", "1.0.0")
	config.App.Timezone = GetEnv("TIMEZONE", "Asia/Jakarta")

	config.PubSub.ProjectId = GetEnv("PROJECT_ID", "")
	config.PubSub.TopicName = GetEnv("TOPIC_NAME", "")

	AppCfg = config
}

func GetEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
