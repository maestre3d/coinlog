package sql

import "github.com/spf13/viper"

type Config struct {
	ConnectionString string
}

func setConfigDefaults() {
	viper.SetDefault("database.connection_string",
		"host=localhost port=5432 user=postgres dbname=coinlog password=root sslmode=disable")
}

func NewConfig() Config {
	setConfigDefaults()
	return Config{
		ConnectionString: viper.GetString("database.connection_string"),
	}
}
