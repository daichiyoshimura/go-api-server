package env

import (
	"github.com/cockroachdb/errors"
)

func (r *Reader) db() (*DB, error) {
	host, err := r.read("DB_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	user, err := r.read("DB_USER")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	password, err := r.read("DB_PASSWORD")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	name, err := r.read("DB_NAME")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	return &DB{
		host:     host,
		user:     user,
		password: password,
		name:     name,
	}, nil
}

type DB struct {
	host     string
	user     string
	password string
	name     string
}

func (e *DB) Host() string {
	return e.host
}

func (e *DB) User() string {
	return e.user
}

func (e *DB) Password() string {
	return e.password
}

func (e *DB) Name() string {
	return e.name
}
