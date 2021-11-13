package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

const (
	readByID            = "select * from account where id = ?"
	readByUsername      = "select * from account where username = ?"
	insertAccount       = "insert into account(username, password_hash) values(?, ?)"
	insertRelationships = "insert into relationship(follower_name, followee_name) values(?, ?)"
	readRelationships   = "select * from relationship where id = ?"
	isFollowed          = "select count(1) from relationship where follower_name = ? and followee_name = ? "
	readFollowings      = "select username, display_name, avatar, note, create_at from account where username = (select followee_name from relationship where follower_name = ?) limit ?"
)

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, readByUsername, username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *account) CreateAccount(ctx context.Context, newAccount *object.Account) (*object.Account, error) {
	result, err := r.db.ExecContext(ctx, insertAccount, newAccount.Username, newAccount.PasswordHash)
	if err != nil {
		return nil, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	insertedAccount := new(object.Account)
	insertedRow := r.db.QueryRowxContext(ctx, readByID, lastID)
	insertedRow.StructScan(&insertedAccount)

	return newAccount, nil
}

func (r *account) FindByID(ctx context.Context, id int64) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, readByID, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *account) Follow(ctx context.Context, followerName, folloeeName string) error {
	result, err := r.db.ExecContext(ctx, insertRelationships, followerName, folloeeName)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (r *account) IsFollowed(ctx context.Context, followerName, folloeeName string) (bool, error) {
	row := r.db.QueryRowContext(ctx, isFollowed, followerName, folloeeName)
	var isFollowed int64
	if err := row.Scan(&isFollowed); err != nil {
		return false, err
	}
	if isFollowed == 0 {
		return false, nil
	}
	return true, nil
}

func (r *account) GetFollowings(ctx context.Context, username string, limit int64) ([]dto.Account, error) {
	rows, err := r.db.QueryxContext(ctx, readFollowings, username, limit)
	if err != nil {
		return nil, err
	}
	var followings []dto.Account
	for rows.Next() {
		var following dto.Account
		if err := rows.StructScan(&following); err != nil {
			return nil, err
		}
		followings = append(followings, following)
	}
	return followings, nil
}
