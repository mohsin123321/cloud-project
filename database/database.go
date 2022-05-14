package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database interface
type DatabaseInterface interface {
	Close()
	InsertData(model.Data)
}

// Database struct
type Database struct {
	DB             *mongo.Client
	DBContextClose context.CancelFunc
	DBContext      context.Context
}

// SetupDB initializes the db and returns it
func SetupDB() *Database {
	var db Database
	var err error

	db.DBContext, db.DBContextClose = context.WithTimeout(context.Background(), 8*time.Second)

	connString := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?maxPoolSize=20&w=majority",
		config.Config.Database.DbType,
		config.Config.Database.DbUser,
		config.Config.Database.DbPass,
		config.Config.Database.DbAddr,
		config.Config.Database.DbPort,
		config.Config.Database.DbName,
	)
	clientOpts := options.Client().ApplyURI(connString)
	db.DB, err = mongo.Connect(db.DBContext, clientOpts)
	if err != nil {
		log.Fatal("Error in db connection :", err)
	}

	// check connection
	if err := db.DB.Ping(db.DBContext, readpref.Primary()); err != nil {
		log.Println(" Cannot connect to the db")
		log.Fatal(err)
	}
	fmt.Println("connected to the db successfully")

	return &db
}

// Close is used to close mongoDB client and cancel context
func (db *Database) Close() {
	// CancelFunc to cancel to context
	defer db.DBContextClose()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// Disconnect method also has deadline.
		// returns error if any,
		if err := db.DB.Disconnect(db.DBContext); err != nil {
			panic(err)
		}
	}()
}
