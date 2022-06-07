package main

import (
	"GoClearArch/internal/composites"
	"GoClearArch/pkg/client/logging"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("create mongodb composite")

	mongoDBComposite, err := composites.NewMongoDBComposite(context.Background(), "localhost", "27017", "", "", "GoClearArch", "test")
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
	fmt.Println("started server at http://localhost:8080/request")
	log.Fatal(http.ListenAndServe(":8080", router))
}