package config

import (
	"log"

	"github.com/spf13/viper"
)

type RabbitMQ struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
}

type Producer struct {
	RabbitMQ RabbitMQ `json:"rabbitmq" mapstructure:"rabbitmq"`
}
type Event struct {
	Producer Producer `json:"producer" mapstructure:"producer"`
}
type Config struct {
	Event Event `json:"events" mapstructure:"events"`
}

func LoadConfig() (*Config, error) {

	viper.SetConfigName("producer_config")
	viper.SetConfigFile("app.yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("could not read env config, error: %v", err)
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)

	return config, err
}
