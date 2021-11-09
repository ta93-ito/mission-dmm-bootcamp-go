package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	readByID       = "select * from account where id = ?"
	readByUsername = "select * from account where username = ?"
	insert         = "insert into account(username, password_hash) values(?, ?)"
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
	result, err := r.db.ExecContext(ctx, insert, newAccount.Username, newAccount.PasswordHash)
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
