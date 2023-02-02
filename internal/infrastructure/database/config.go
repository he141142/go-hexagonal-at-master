package database

import config "hex-base/internal/common"

type DatabaseConfig interface {
	Host() string
	Port() string
	Driver() string
	User() string
	Password() string
	SslMode() string
	DBName() string
	Metadata() map[string]interface{}
	AddMetadata(data map[string]interface{}, override bool) DatabaseConfig
}

type databaseConfig struct {
	DatabaseConfig
	host     string
	port     string
	driver   string
	user     string
	password string
	dbName   string
	sslMode  string
	metadata map[string]interface{}
}

func (cf *databaseConfig) Host() string {
	return cf.host
}
func (cf *databaseConfig) Port() string {
	return cf.port
}
func (cf *databaseConfig) Driver() string {
	return cf.driver
}
func (cf *databaseConfig) User() string {
	return cf.user
}
func (cf *databaseConfig) Password() string {
	return cf.password
}

func (cf *databaseConfig) SslMode() string {
	return cf.sslMode
}

func (cf *databaseConfig) DBName() string {
	return cf.dbName
}
func (cf *databaseConfig) Metadata() map[string]interface{} {
	return cf.metadata
}

func (cf *databaseConfig) AddMetadata(metadata map[string]interface{}, override bool) DatabaseConfig {
	for key, val := range metadata {
		if _, ok := cf.metadata[key]; ok || override {
			cf.metadata[key] = val
		}
	}
	return cf
}

func NewSqlConfig() DatabaseConfig {
	return &databaseConfig{
		host:     config.Viper().Get("DB_HOST"),
		port:     config.Viper().Get("DB_PORT"),
		driver:   config.Viper().Get("DB_DRIVER"),
		user:     config.Viper().Get("DB_USER"),
		password: config.Viper().Get("DB_PASSWORD"),
		dbName:   config.Viper().Get("DB_NAME"),
		sslMode:  config.Viper().Get("DB_SSL"),
		metadata: make(map[string]interface{}),
	}
}
