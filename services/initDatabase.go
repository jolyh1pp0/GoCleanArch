package services

import (
	"GoClearArch/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const maxRequests = 50

func InitDatabase(cl *mongo.Client, dbName, activeRequests, cancelledRequests string) *models.MongoDBCollections {
	collection := &models.MongoDBCollections{
		ActiveRequests:    cl.Database(dbName).Collection(activeRequests),
		CancelledRequests: cl.Database(dbName).Collection(cancelledRequests),
	}

	_, err := collection.ActiveRequests.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = collection.ActiveRequests.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}

	_, err = collection.CancelledRequests.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = collection.CancelledRequests.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < maxRequests; i++ {
		randomRequest := models.Request{
			Name:  GenerateRandomRequest(),
			Count: 0,
		}
		_, err := collection.ActiveRequests.InsertOne(context.TODO(), randomRequest)
		if err != nil {
			log.Fatal(err)
		}
	}

	ticker := time.Tick(200 * time.Millisecond)
	go func() {
		for range ticker {
			changeRandomRequest(collection)
		}
	}()
	return collection
}