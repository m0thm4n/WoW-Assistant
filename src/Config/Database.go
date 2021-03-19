package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
  r "gopkg.in/rethinkdb/rethinkdb-go.v6"
  "log"
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

type RethinkConfig struct {
  Address   string
  Database  string
  Username  string
  Password  string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host: 		"10.11.13.38",
		Port: 		3306,
		User: 		"root",
		Password:	"Babycakes15!",
		DBName:		"wow",
	}

	return &dbConfig
}

func BuildRethinkConfig() *RethinkConfig {
  dbConfig := RethinkConfig{
    Address:    "10.11.13.38",
    Username:   "root",
    Password:   "Babycakes15!",
    Database:   "wow",
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

func ConntectRethink(dbConfig *RethinkConfig) *r.Session {
  var err error

  session, err := r.Connect(r.ConnectOpts{
    Address: dbConfig.Address,
    Database: dbConfig.Database,
    Username: dbConfig.Username,
    Password: dbConfig.Password,
  })
  if err != nil {
    log.Fatalln(err.Error())
  }

  return session
}
