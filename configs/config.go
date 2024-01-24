package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")

	// Load .env.local if it exists
	if _, err := os.Stat(".env.local"); err == nil {
		viper.SetConfigName(".env.local")
		err = viper.ReadInConfig()
		if err != nil {
			log.Println("failed to load .env.local file")
		}
	} else {
		// Load .env if .env.local does not exist or failed to load
		viper.SetConfigName(".env")
		err = viper.ReadInConfig()
		if err != nil {
			log.Println("failed to load .env file")
		}
	}

	// Load environment variables from OS environment
	viper.AutomaticEnv()
}
