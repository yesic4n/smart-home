package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

var connect *mongo.Client
var lock = &sync.Mutex{}

func getURI() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Failed to load .env file")
	// }

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	cluster := os.Getenv("DB_CLUSTER")

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", user, password, cluster)
	return uri
}

func createConnect() *mongo.Client {
	// Set Client Options
	uri := getURI()
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetConnection() *mongo.Client {
	if connect == nil {
		// Apply Singleton Design Pattern
		lock.Lock()
		defer lock.Unlock()
		if connect == nil {
			connect = createConnect()
			fmt.Println("MongoDB Connected")
		} else {
			return connect
		}
	}
	return connect
}

// GetCollection Get Collection DB
// Users
// Sensors
// Actions
// Notifications
func GetCollection(collectionName string) *mongo.Collection {
	switch collectionName {
	case "Users":
		return GetConnection().Database("SmartHomeDB").Collection("Users")
	case "Sensors":
		return GetConnection().Database("SmartHomeDB").Collection("Sensors")
	case "Actions":
		return GetConnection().Database("SmartHomeDB").Collection("Actions")
	case "Notifications":
		return GetConnection().Database("SmartHomeDB").Collection("Notifications")
	}
	return nil
}
