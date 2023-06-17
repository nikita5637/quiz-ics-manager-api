package config

import (
	"fmt"

	"go.uber.org/zap"
)

var (
	defaultLogLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

// LoggerConfig ...
type LoggerConfig struct {
	LogLevel string `toml:"log_level"`

	ElasticAddress     string `toml:"elastic_address"`
	ElasticIndex       string `toml:"elastic_index"`
	ElasticLogsEnabled bool   `toml:"elastic_logs_enabled"`
	ElasticPort        uint16 `toml:"elastic_port"`
}

// GetElasticAddress ...
func GetElasticAddress() string {
	return fmt.Sprintf("http://%s:%d", globalConfig.ElasticAddress, globalConfig.ElasticPort)
}

// GetLogLevel ...
func GetLogLevel() zap.AtomicLevel {
	level, err := zap.ParseAtomicLevel(globalConfig.LogLevel)
	if err != nil {
		level = defaultLogLevel
	}

	return level
}
