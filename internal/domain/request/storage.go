package request

import (
	"GoClearArch/internal/domain"
	"GoClearArch/models"
	"GoClearArch/services"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type requestStorage struct {
	db *mongo.Database
}


func NewStorage(db *mongo.Database) domain.Storage {
	return &requestStorage{db: db}
}

func (r *requestStorage) GetRandomAliveApplication(ctx context.Context) (models.Request, error) {
	// get random
	randINT := services.Random(0, 49)

	skip := int64(randINT)

	// made skip options for find one options
	opts := &options.FindOneOptions{
		Skip: &skip,
	}

	app := models.Request{}
	// find one
	err := r.db.Collection("activeRequests").FindOne(context.TODO(), bson.D{}, opts).Decode(&app)
	if err != nil {
		log.Fatal(err)
		return models.Request{}, err
	}

	// made filter for update one
	filter := bson.D{{"_id", app.ID}}
	// made update data foe update one
	update := bson.D{
		{"$inc", bson.D{
			{"count", 1},
		}},
	}

	// update one
	_, err = r.db.Collection("activeRequests").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return models.Request{}, err
	}

	return models.Request{
		Name:  app.Name,
		Count: app.Count,
	}, nil
}

func (r *requestStorage) GetShowedAndCancelApplications(ctx context.Context) (rezActive []models.Request, rezCancel []models.Request, err error) {
	// ADD ACTIVE
	// find many active
	filter := bson.M{
		"count": bson.M{"$gt": 0},
	}

	curA, err := r.db.Collection("activeRequests").Find(context.TODO(), filter, &options.FindOptions{})
	if err != nil {
		log.Fatal(err)
		return rezActive, rezCancel, err
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

		rezActive = append(rezActive, models.Request{
			Name:  episode.Name,
			Count: episode.Count,
		})
	}

	// ADD CANCEL
	// find all canceled
	curC, err := r.db.Collection("cancelledRequests").Find(context.TODO(), bson.M{}, &options.FindOptions{})
	if err != nil {
		log.Fatal(err)
		return rezActive, rezCancel, err
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

		rezCancel = append(rezCancel, models.Request{
			Name:  episode.Name,
			Count: episode.Count,
		})
	}

	return rezActive, rezCancel, nil
}