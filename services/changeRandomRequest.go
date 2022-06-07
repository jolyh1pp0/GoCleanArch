package services

import (
	"GoClearArch/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func changeRandomRequest(collection *models.MongoDBCollections) {
	randINT := Random(0, maxRequests-1)
	skip := int64(randINT)
	opts := &options.FindOneOptions{
		Skip: &skip,
	}

	randomRequest := models.Request{}
	err := collection.ActiveRequests.FindOne(context.TODO(), bson.D{}, opts).Decode(&randomRequest)
	if err != nil {
		log.Fatal(err)
	}

	if randomRequest.Count > 0 {
		tempApp := models.Request{
			Name:  randomRequest.Name,
			Count: randomRequest.Count,
		}

		_, err = collection.CancelledRequests.InsertOne(context.TODO(), tempApp)
		if err != nil {
			log.Fatal(err)
		}
	}

	filter := bson.D{{"_id", randomRequest.ID}}

	_, err = collection.ActiveRequests.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	_, err = collection.ActiveRequests.InsertOne(context.TODO(), models.Request{
		Name:  GenerateRandomRequest(),
		Count: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
}