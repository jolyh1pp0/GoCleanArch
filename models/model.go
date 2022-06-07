package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Request struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string 			 `bson:"name,omitempty"`
	Count int       		 `bson:"count,omitempty"`
}

type MongoDBCollections struct {
	ActiveRequests    *mongo.Collection
	CancelledRequests *mongo.Collection
}
