package configuration

import "github.com/spf13/viper"

type DatabaseSQL struct {
	ConnectionString string
}

func setDatabaseSQLDefaults() {
	viper.SetDefault("database.connection_string",
		"host=localhost port=6432 user=postgres dbname=coinlog password=root sslmode=disable")
}

func NewDatabaseSQL() DatabaseSQL {
	setDatabaseSQLDefaults()
	return DatabaseSQL{
		ConnectionString: viper.GetString("database.connection_string"),
	}
}
