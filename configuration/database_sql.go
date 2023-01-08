package configuration

import "github.com/spf13/viper"

type DatabaseSQL struct {
	ConnectionString string
}

func NewDatabaseSQL() DatabaseSQL {
	return DatabaseSQL{
		ConnectionString: viper.GetString("database.connection_string"),
	}
}
