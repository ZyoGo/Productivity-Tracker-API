package util

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/w33h/Productivity-Tracker-API/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	// Postgres Driver
	Postgres DatabaseDriver = "postgres"
	// Mongodb Driver
	MongoDB DatabaseDriver = "mongodb"
)

type DatabaseConfig struct {
	Driver DatabaseDriver

	// MongoDB config
	MongoDB     *mongo.Database
	mongoClient *mongo.Client

	// Postgres config
	PostgreSQL *gorm.DB
}

func NewConnectionDB(config *config.AppConfig) *DatabaseConfig {
	var dbConfig DatabaseConfig

	switch config.Database.Driver {
	case "postgres":
		dbConfig.Driver = Postgres
		dbConfig.PostgreSQL = newPostgresConnection(config)
	case "mongodb":
		dbConfig.Driver = MongoDB
		dbConfig.mongoClient = newMongoDBClient(config)
		dbConfig.MongoDB = dbConfig.mongoClient.Database(config.Database.Name)
	default:
		panic("Unsupported database driver")
	}

	return &dbConfig
}

func newPostgresConnection(config *config.AppConfig) *gorm.DB {
	var uri string
	uri = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable connect_timeout=10 timezone=Asia/Makassar",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Name,
		config.Database.Password)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}

	return db
}

func newMongoDBClient(config *config.AppConfig) *mongo.Client {
	uri := "mongodb://"

	if config.Database.Username != "" {
		uri = fmt.Sprintf("%s%v:%v@", uri, config.Database.Username, config.Database.Password)
	}

	uri = fmt.Sprintf("%s%v:%v",
		uri,
		config.Database.Host,
		config.Database.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}

func (db *DatabaseConfig) Close() {
	switch db.Driver {
	case "postgres":
		db, _ := db.PostgreSQL.DB()
		db.Close()
	case "mongodb":
		db.mongoClient.Disconnect(context.Background())
	default:
		panic("Unsupported database driver")
	}
}
