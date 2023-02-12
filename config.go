package coinlog

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	MajorVersion uint16
}

func NewConfig() Config {
	viper.SetEnvPrefix("NCORP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return Config{
		MajorVersion: viper.GetUint16("application.major_version"),
	}
}
