package request

import "context"

type Service interface {
	GetRequest(ctx context.Context) (string, error)
	GetAdminRequests(ctx context.Context) ([]string, error)
}

