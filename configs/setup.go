package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[Database] Connected to MongoDB")

	// Ensuring Indexes are Set
	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EnvMongoDatabase()).Collection(collectionName)
	return collection
}

// CreateIndex - creates an index for a specific field in a collection
func CreateIndex(collection *mongo.Collection, field string, unique bool) bool {

	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	// 2. Create the context for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Create a single index
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		// 5. Something went wrong, we log it and return false
		log.Printf("[Database] collection %v index %v - Unique %v - Err: %v", collection.Name(), field, unique, err)
		return false
	}

	// 6. All went well, we return true
	return true
}
