package mongodb

import (
	"GoClearArch/pkg/client/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewClient(ctx context.Context, host, port, username, password, database, authSource string) (*mongo.Database, error) {
	var mongoDBURL string
	var anonymous bool
	if username == "" || password == "" {
		anonymous = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}
	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if !anonymous {
		clientOptions.SetAuth(options.Credential{
			AuthSource: authSource,
			Username: username,
			Password: password,
			PasswordSet: true,
		})
	}
	client, err := mongo.Connect(reqCtx, clientOptions)
	if err != nil {
		fmt.Errorf("Failed to create client to mongoDB due to error %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Errorf("Failed to create client to mongoDB due to error %w", err)
	}

	logging.Init()
	logger := logging.GetLogger()
	logger.Info("mongodb connection successful")

	return client.Database(database), nil
}
