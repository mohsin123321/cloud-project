package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

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
	DB *mongo.Database
}

// SetupDB initializes the db and returns it
func SetupDB() *Database {
	var err error
	var replicas []string
	for _, r := range config.Config.Database.Replicas {
		replicaAddress := r.Host + ":" + r.Port
		replicas = append(replicas, replicaAddress)
	}

	// entire replica string
	replicasPath := strings.Join(replicas, ",")

	// final string should look like this
	// mongodb://username:password@replicahost:replicaPort,replicahost:replicaPort/database?replicaSet=replicaSetName
	connString := fmt.Sprintf(
		"%s://%s:%s@%s/%s?replicaSet=%s",
		config.Config.Database.DbType,
		config.Config.Database.DbUser,
		url.QueryEscape(config.Config.Database.DbPass),
		replicasPath,
		config.Config.Database.DbName,
		config.Config.Database.ReplicaName,
	)
	log.Println(connString)

	clientOpts := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal("Error in db connection :", err)
	}

	// check connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println("Cannot connect to the db")
		log.Fatal(err)
	}

	fmt.Println("connected to the db successfully")
	return &Database{
		DB: client.Database(config.Config.Database.DbName),
	}
}

// Close is used to close mongoDB client and cancel context
func (db *Database) Close() {
	// Disconnect method also has deadline.
	// returns error if any,
	if err := db.DB.Client().Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
