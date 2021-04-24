package store

import (
	"context"

	"{{GoModule}}/internal/domain"
)

type Interface interface {
	GetHello(ctx context.Context, who string) (*domain.Hello, error)
}
