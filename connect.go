package connect

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewCollection(databaseName, collectionName string) *mongo.Collection {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	collection := client.Database(databaseName).Collection(collectionName)
	// var entitieTcar = entities.TCar{
	// 	Model:    "gol",
	// 	Year:     2015,
	// 	Color:    "blue",
	// 	Status:   true,
	// 	BuyValue: 20000,
	// 	DoorsQty: 4,
	// 	SeatsQty: 5,
	// }

	// collection.InsertOne(context.TODO(), entitieTcar)
	return collection
}
