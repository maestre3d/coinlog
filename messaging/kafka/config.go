package kafka

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Address               string
	WriterTimeout         time.Duration
	EnableTopicGeneration bool
}

func setDefaultConfig() {
	// K8s uses SRV DNS records for services, Service Discovery will route to any of the registered IPs
	viper.SetDefault("messaging.kafka.address", "localhost:9092")
	viper.SetDefault("messaging.kafka.writer_timeout", time.Second*10)
	viper.SetDefault("messaging.kafka.enable_topic_generation", true)
}

func NewConfig() Config {
	setDefaultConfig()
	return Config{
		Address:               viper.GetString("messaging.kafka.address"),
		WriterTimeout:         viper.GetDuration("messaging.kafka.writer_timeout"),
		EnableTopicGeneration: viper.GetBool("messaging.kafka.enable_topic_generation"),
	}
}
