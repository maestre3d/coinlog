package configuration

import "github.com/spf13/viper"

type Application struct {
	MajorVersion uint16
}

func NewApplication() Application {
	return Application{
		MajorVersion: viper.GetUint16("application.major_version"),
	}
}
