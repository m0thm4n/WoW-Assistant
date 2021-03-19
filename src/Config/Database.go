package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db Configuration
type DBConfig struct {
	Host		string
	Port		int
	User		string
	DBName		string
	Password	string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host: 		"10.11.13.38",
		Port: 		3306,
		User: 		"root",
		Password:	"Babycakes15!",
		DBName:		"podify",
	}

	return &dbConfig
}

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
