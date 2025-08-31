package mysql

import (
	"database/sql"
	"io"
)

// struct for db ops
type chatDB struct {
	CreateUser *sql.Stmt
}

var (
	_ Database  = &chatDB{}
	_ io.Closer = &chatDB{}
)

func NewChatDB(db *sql.DB) (*chatDB, error) {
	var (
		database = &chatDB{}
		err      error
	)
	if database.CreateUser, err = db.Prepare(create_user); err != nil {
		return nil, err
	}
	return database, nil
}

func (db *chatDB) Close() error {
	db.CreateUser.Close()
	return nil
}
