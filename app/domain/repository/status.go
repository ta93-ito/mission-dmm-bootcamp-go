package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, content *object.Status) (*object.Status, error)
}
