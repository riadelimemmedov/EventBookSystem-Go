package app

import "book_event/common/postgresql"

// import "/common/postgresql/"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

//!NewConfigurationManager
func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

//!getPostgreSqlConfig
func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "bookevent",
		UserName:              "postgres",
		Password:              "123321",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
