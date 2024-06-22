package usecase

import "context"

type Book interface {
	Get(ctx context.Context) []string
	GetByID(ctx context.Context, id string) string
}
