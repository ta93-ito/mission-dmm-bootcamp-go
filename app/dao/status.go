package dao

import (
	"context"
	"fmt"
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

func (r *status) CreateStatus(ctx context.Context, newStatus *object.Status) (*object.Status, error) {
	const (
		insert = "insert into status(content, account_id) values(?, ?)"
		read   = "select * from status where id = ?"
	)
	// content入ってない。。。
	fmt.Println(newStatus.Content)
	fmt.Printf("(%%#v) %#v\n", newStatus)
	result, err := r.db.ExecContext(ctx, insert, newStatus.Content, newStatus.AccountID)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	insertedStatus := new(object.Status)
	insertedRow := r.db.QueryRowxContext(ctx, read, lastID)
	insertedRow.StructScan(&insertedStatus)
	return insertedStatus, nil
}
func (r *status) GetStatus(ctx context.Context, id int64) (*object.Status, error) {
	const read = "select * from status where id = ?"

	status := new(object.Status)
	row := r.db.QueryRowxContext(ctx, read, id)
	if err := row.StructScan(&status); err != nil {
		return nil, err
	}
	return status, nil
}
