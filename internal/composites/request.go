package composites

import (
	"GoClearArch/internal/adapters/api"
	"GoClearArch/internal/adapters/api/request"
	"GoClearArch/internal/domain"
	request2 "GoClearArch/internal/domain/request"
)

type RequestComposite struct {
	domain.Storage
	request2.Service
	Handler api.Handler
}

func NewRequestComposite(mongoComposite *MongoDBComposite) (*RequestComposite, error) {
	storage := request2.NewStorage(mongoComposite.db)
	service := request2.NewService(storage)
	handler := request.NewHandler(service)

	return &RequestComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
