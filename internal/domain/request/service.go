package request

import (
	"GoClearArch/internal/domain"
	"context"
	"log"
	"strconv"
)

type service struct {
	storage domain.Storage
}

func NewService(storage domain.Storage) Service {
	return &service{storage: storage}
}

func (s *service) GetRequest(ctx context.Context) (string, error) {
	//get application
	app, err := s.storage.GetRandomAliveApplication(ctx)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return app.Name, nil
}

func (s *service) GetAdminRequests(ctx context.Context) ([]string, error) {
	// get applications
	Active, Cancel, err := s.storage.GetShowedAndCancelApplications(ctx)

	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	// crete slice rez
	rez := make([]string, 0)

	for _, app := range Active {
		rez = append(rez, "Active("+app.Name+"-"+strconv.Itoa(app.Count)+")")
	}

	for _, app := range Cancel {
		rez = append(rez, "Cancelled("+app.Name+"-"+strconv.Itoa(app.Count)+")")
	}

	return rez, nil
}
