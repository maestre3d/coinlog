package coinlog

import "github.com/spf13/viper"

type Config struct {
	MajorVersion uint16
}

func NewConfig() Config {
	return Config{
		MajorVersion: viper.GetUint16("application.major_version"),
	}
}
