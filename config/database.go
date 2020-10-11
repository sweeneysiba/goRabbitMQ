package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//DB ...
var DB *gorm.DB

// DBConfig represents db configuration

//DBConfig ...
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

//BuildDBConfig ...
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root",
		DBName:   "eastern_enterprise",
	}
	return &dbConfig
}

//DbURL ...
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
