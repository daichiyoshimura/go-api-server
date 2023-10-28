package env

import (
	"os"

	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

const (
	errMsgEnv     string = "env: %w"
	errMsgEmpty   string = "ENV must be set : %v"
	errMsgLoadEnv string = "failed to load .env: %w"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read() (*Server, *DB, error) {
	stg := NewStage(os.Getenv("STAGE"))
	if stg.isDev() {
		if err := godotenv.Load(".env"); err != nil {
			return nil, nil, errors.Errorf(errMsgLoadEnv, err)
		}
	}

	srv, err := r.server()
	if err != nil {
		return nil, nil, err
	}

	db, err := r.db()
	if err != nil {
		return nil, nil, err
	}

	return srv, db, nil
}

func (r *Reader) read(key string) (string, error) {
	val := os.Getenv(key)
	if len(val) == 0 {
		
		return val, errors.Errorf(errMsgEnv, errors.Newf(errMsgEmpty, key))
	}

	return val, nil
}

func (r *Reader) server() (*Server, error) {
	host, err := r.read("SERVER_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgLoadEnv, err)
	}

	return &Server{
		host: host,
	}, nil
}

func (r *Reader) db() (*DB, error) {
	host, err := r.read("DB_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgLoadEnv, err)
	}

	user, err := r.read("DB_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgLoadEnv, err)
	}

	password, err := r.read("DB_PASSWORD")
	if err != nil {
		return nil, errors.Errorf(errMsgLoadEnv, err)
	}

	instance, err := r.read("DB_INSTANCE")
	if err != nil {
		return nil, errors.Errorf(errMsgLoadEnv, err)
	}

	return &DB{
		host:     host,
		user:     user,
		password: password,
		instance: instance,
	}, nil
}
