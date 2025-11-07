package store

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]*Store, error)
}
