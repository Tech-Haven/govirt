package configs

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct{
    Db          *mongo.Client
    HttpClient  *http.Client
}

func New() *Config {
    // Get database client
    db := connectDB()

    t := http.DefaultTransport.(*http.Transport).Clone()
    t.MaxIdleConns = 100
    t.MaxConnsPerHost = 100
    t.MaxIdleConnsPerHost = 100

    // Create http client
    httpClient := &http.Client{
        Timeout: time.Second * 2,
        Transport: t,
    }

    return &Config{
        Db: db,
        HttpClient: httpClient,
    }
}

func connectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) * mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}