package config

import "fmt"

const (
	// amqp://guest:guest@localhost:5672/
	rabbitMQURLFormat = "amqp://%s:%s@%s:%d/"
)

// ICSManagerConfig ...
type ICSManagerConfig struct {
	BindAddress          string `toml:"bind_address"`
	BindPort             uint16 `toml:"bind_port"`
	ICSFileExtension     string `toml:"ics_file_extension"`
	ICSFilesFolder       string `toml:"ics_files_folder"`
	RabbitMQAddress      string `toml:"rabbitmq_address"`
	RabbitMQICSQueueName string `toml:"rabbitmq_ics_queue_name"`
	RabbitMQPort         uint16 `toml:"rabbitmq_port"`
	RabbitMQUserName     string `toml:"rabbitmq_username"`

	RegistratorAPIAddress string `toml:"registrator_api_address"`
	RegistratorAPIPort    uint16 `toml:"registrator_api_port"`
}

// GetBindAddress ...
func GetBindAddress() string {
	return fmt.Sprintf("%s:%d", globalConfig.BindAddress, globalConfig.BindPort)
}

// GetRabbitMQURL ...
func GetRabbitMQURL() string {
	rabbitMQPassword := GetSecretValue(RabbitMQPassword)

	return fmt.Sprintf(rabbitMQURLFormat,
		globalConfig.RabbitMQUserName,
		rabbitMQPassword,
		globalConfig.RabbitMQAddress,
		globalConfig.RabbitMQPort,
	)
}
