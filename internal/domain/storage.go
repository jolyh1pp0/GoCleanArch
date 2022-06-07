package domain

import (
	"GoClearArch/models"
	"context"
)

type Storage interface {
	GetRandomAliveApplication(ctx context.Context) (models.Request, error)
	GetShowedAndCancelApplications(ctx context.Context) ([]models.Request, []models.Request, error)
}
