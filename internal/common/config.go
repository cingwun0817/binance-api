package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// loadConfig loads the config file.
func LoadConfig() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("[%s] error: %w", "internal.common.config", err))
	}
}
