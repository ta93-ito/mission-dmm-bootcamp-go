package repository

import (
	"context"

	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	FindByID(ctx context.Context, id int64) (*object.Account, error)
	// TODO: Add Other APIs
	CreateAccount(ctx context.Context, account *object.Account) (*object.Account, error)
	Follow(ctx context.Context, followee_name, follower_name string) error
	IsFollowed(ctx context.Context, followee_name, follower_name string) (bool, error)
	GetFollowings(ctx context.Context, username string, limit int64) ([]dto.Account, error)
}
