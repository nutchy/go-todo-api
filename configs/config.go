package configs

import (
	"time"

	"github.com/spf13/viper"
)

var C Config

type Config struct {
	A        string        `mapstucture:"a"`
	Duration time.Duration `mapstructure:"duration"`
	Joey     struct {
		FirstName string `mapstructure:"first_name"`
		LastName  string `mapstructure:"last_name"`
	} `mapstructure:"joey"`
}

func NewConfig() error {
	viper.AddConfigPath("./configs")

	// read from config.yaml file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	err := viper.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}
