package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// amqp://guest:guest@localhost:5672/
	rabbitMQURLFormat = "amqp://%s:%s@%s:%d/"
)

func initICSConsumerConfigureParams() {
	_ = viper.BindEnv("ics_consumer.rabbitmq.address")
	_ = viper.BindEnv("ics_consumer.rabbitmq.credentials.password")
	_ = viper.BindEnv("ics_consumer.registrator_api.address")
}

// GetRabbitMQURL ...
func GetRabbitMQURL() string {
	return fmt.Sprintf(rabbitMQURLFormat,
		viper.GetString("ics_consumer.rabbitmq.credentials.username"),
		viper.GetString("ics_consumer.rabbitmq.credentials.password"),
		viper.GetString("ics_consumer.rabbitmq.address"),
		viper.GetUint32("ics_consumer.rabbitmq.port"),
	)
}
