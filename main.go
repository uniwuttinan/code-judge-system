package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/wuttinanhi/code-judge-system/configs"
	"github.com/wuttinanhi/code-judge-system/consumers"
	"github.com/wuttinanhi/code-judge-system/controllers"
	"github.com/wuttinanhi/code-judge-system/databases"
	"github.com/wuttinanhi/code-judge-system/services"
)

func main() {
	configs.LoadConfig()

	log.Printf("MySQL database at %s:%s", viper.GetString("DB_HOST"), viper.GetString("DB_PORT"))
	db := databases.NewMySQLDatabase()
	serviceKit := services.CreateServiceKit(db)

	APP_MODE := viper.GetString("APP_MODE")

	if APP_MODE == "CONSUMER" {
		consumers.StartSubmissionConsumer(serviceKit)
		return
	}

	rateLimitStorage := controllers.GetRedisStorage()
	log.Printf(
		"Redis connection at %s@%s:%d",
		viper.GetString("RATE_LIMIT_USER"),
		viper.GetString("RATE_LIMIT_HOST"),
		viper.GetInt("RATE_LIMIT_PORT"),
	)
	api := controllers.SetupAPI(serviceKit, rateLimitStorage)
	api.Listen(":3000")
}
