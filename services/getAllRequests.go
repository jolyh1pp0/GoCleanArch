package services

import (
	"GoClearArch/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetAllRequests(collection *models.MongoDBCollections) (activeRequests []models.Request, cancelledRequests []models.Request, err error) {
	filter := bson.M{
		"count": bson.M{"$gt": 0},
	}

	curA, err := collection.ActiveRequests.Find(context.TODO(), filter, &options.FindOptions{})
	if err != nil {
		log.Fatal(err)
		return activeRequests, cancelledRequests, err
	}

	defer func() {
		err = curA.Close(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}()

	// parse all
	for curA.Next(context.TODO()) {
		var episode models.Request
		if err = curA.Decode(&episode); err != nil {
			log.Fatal(err)
		}

		activeRequests = append(activeRequests, models.Request{
			Name:  episode.Name,
			Count: episode.Count,
		})
	}

	// ADD CANCEL
	// find all canceled
	curC, err := collection.CancelledRequests.Find(context.TODO(), bson.M{}, &options.FindOptions{})
	if err != nil {
		log.Fatal(err)
		return activeRequests, cancelledRequests, err
	}

	defer func() {
		err = curC.Close(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}()

	// parse all
	for curC.Next(context.TODO()) {
		var episode models.Request
		if err = curC.Decode(&episode); err != nil {
			log.Fatal(err)
		}

		cancelledRequests = append(cancelledRequests, models.Request{
			Name:  episode.Name,
			Count: episode.Count,
		})
	}

	return activeRequests, cancelledRequests, nil
}
