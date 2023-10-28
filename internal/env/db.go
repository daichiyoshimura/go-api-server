package env

import "github.com/cockroachdb/errors"

func (r *Reader) db() (*DB, error) {
	host, err := r.read("DB_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	user, err := r.read("DB_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	password, err := r.read("DB_PASSWORD")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	instance, err := r.read("DB_INSTANCE")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	return &DB{
		host:     host,
		user:     user,
		password: password,
		instance: instance,
	}, nil
}

type DB struct {
	host     string
	user     string
	password string
	instance string
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

func (e *DB) Instance() string {
	return e.instance
}
