package http

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddress             string
	RootPathFormat            string
	MaxRequestBodySize        string
	MaxRate                   float64
	GracefulShutdownThreshold time.Duration
}

func setConfigDefaults() {
	viper.SetDefault("server.http.listen_address", ":8080")
	viper.SetDefault("server.http.root_path_format", "/api/v%d")
	viper.SetDefault("server.http.graceful_shutdown_threshold", time.Second*15)
	viper.SetDefault("server.http.max_request_body_size", "20M")
	viper.SetDefault("server.http.max_rate", 100.0)
}

func NewConfig() Config {
	setConfigDefaults()
	return Config{
		ListenAddress:             viper.GetString("server.http.listen_address"),
		RootPathFormat:            viper.GetString("server.http.root_path_format"),
		MaxRequestBodySize:        viper.GetString("server.http.max_request_body_size"),
		MaxRate:                   viper.GetFloat64("server.http.max_rate"),
		GracefulShutdownThreshold: viper.GetDuration("server.http.graceful_shutdown_threshold"),
	}
}
