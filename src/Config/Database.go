package Config

import (
    "WoW-Assistant/src/Utils"
    "context"
    "fmt"
    "github.com/goonode/mogo"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"

    "github.com/jinzhu/gorm"
    r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var mongoConnection *mogo.Connection = nil
var DB *gorm.DB

// DBConfig represents db Configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

type RethinkConfig struct {
	Address  string
	Database string
	Username string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "10.11.13.119",
		Port:     3306,
		User:     "root",
		Password: "Babycakes15!",
		DBName:   "wow",
	}

	return &dbConfig
}

func BuildRethinkConfig() *RethinkConfig {
	dbConfig := RethinkConfig{
		Address:  "10.11.13.119",
		Username: "root",
		Password: "Babycakes15!",
		Database: "wow",
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

func ConnectRethink(dbConfig *RethinkConfig) *r.Session {
	var err error

	session, err := r.Connect(r.ConnectOpts{
		Address:  dbConfig.Address,
		Database: dbConfig.Database,
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return session
}

// GetConnection is for get mongo connection
func GetConnectionItem(realmName string) *mongo.Collection {
    if mongoConnection == nil {
        connectionString := Utils.EnvVar("DB_CONNECTION_STRING", "")
        dbName := Utils.EnvVar("DB_NAME", "")

        URI := "mongodb://" + connectionString + "/"

        ctx := context.Background()
        clientOpts := options.Client().ApplyURI(URI)
        client, err := mongo.Connect(ctx, clientOpts)
        if err != nil {
            fmt.Println("can't connect to mongodb server")
        }
        fmt.Println("Connected to server")

        db := client.Database(dbName)
        fmt.Print("Connected to database")
        coll := db.Collection(realmName)
        fmt.Println("Created Collection")

        return coll
    }
    return nil
}
