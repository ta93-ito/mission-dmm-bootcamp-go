package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, content *object.Status) (*object.Status, error)
	FindStatus(ctx context.Context, id int64) (*object.Status, error)
	DeleteStatus(ctx context.Context, id int64) error
}
