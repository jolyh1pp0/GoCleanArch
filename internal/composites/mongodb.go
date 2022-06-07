package composites

import (
	"GoClearArch/pkg/client/logging"
	"GoClearArch/pkg/client/mongodb"
	"GoClearArch/services"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, Host, Port, Username, Password, Database, AuthSource string) (*MongoDBComposite, error) {
	client, err := mongodb.NewClient(ctx, Host, Port, Username, Password, Database, AuthSource)
	if err != nil {
		return nil, err
	}
	services.InitDatabase(client.Client(), "GoClearArch", "activeRequests", "cancelledRequests")

	logging.Init()
	logger := logging.GetLogger()
	logger.Info("database initialized")

	return &MongoDBComposite{db: client}, nil
}
