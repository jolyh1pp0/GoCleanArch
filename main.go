package main

import (
	"GoClearArch/config"
	"GoClearArch/internal/composites"
	"GoClearArch/pkg/client/logging"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("connecting config")
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	logger.Info("config connected")

	logger.Info("create mongodb composite")
	mongoDBComposite, err := composites.NewMongoDBComposite(context.Background(), viper.GetString("mongo.host"), viper.GetString("mongo.port"), viper.GetString("mongo.username"), viper.GetString("mongo.password"), viper.GetString("mongo.database"), viper.GetString("mongo.authsource"))
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("request composite initializing")
	requestComposite, err := composites.NewRequestComposite(mongoDBComposite)
	if err != nil {
		logger.Fatal("request composite failed")
	}

	logger.Info("router initializing")
	router := http.NewServeMux()
	requestComposite.Handler.Register(router)

	logger.Info("server start")
	fmt.Println("started server at http://localhost:" + viper.GetString("server.port") +"/request")
	log.Fatal(http.ListenAndServe(":" + viper.GetString("server.port"), router))
}