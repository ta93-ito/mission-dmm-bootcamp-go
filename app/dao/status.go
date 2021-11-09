package dao

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

const (
	readStatus   = "select * from status where id = ?"
	insertStatus = "insert into status(content, account_id) values(?, ?)"
)

func (r *status) CreateStatus(ctx context.Context, newStatus *object.Status) (*object.Status, error) {
	result, err := r.db.ExecContext(ctx, insert, newStatus.Content, newStatus.AccountID)
	if err != nil {
		return nil, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	insertedStatus, err := r.FindStatus(ctx, lastID)
	if err != nil {
		return nil, err
	}
	return insertedStatus, nil
}

func (r *status) FindStatus(ctx context.Context, id int64) (*object.Status, error) {
	status := new(object.Status)
	if err := r.db.QueryRowxContext(ctx, readStatus, id).StructScan(status); err != nil {
		return nil, err
	}
	return status, nil
}
