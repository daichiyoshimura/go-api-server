package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	errMsgEmpty string = "ENV must be set :%v"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read() (*Server, *DB, error) {
	stg := NewStage(os.Getenv("STAGE"))
	if stg.isDev() {
		if err := godotenv.Load(".env"); err != nil {
			return nil, nil, err
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
		err := fmt.Errorf(errMsgEmpty, key)
		return val, err
	}

	return val, nil
}

func (r *Reader) server() (*Server, error) {
	host, err := r.read("SERVER_HOST")
	if err != nil {
		return nil, err
	}

	return &Server{
		host: host,
	}, nil
}

func (r *Reader) db() (*DB, error) {
	host, err := r.read("DB_HOST")
	if err != nil {
		return nil, err
	}

	user, err := r.read("DB_HOST")
	if err != nil {
		return nil, err
	}

	password, err := r.read("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	instance, err := r.read("DB_INSTANCE")
	if err != nil {
		return nil, err
	}

	return &DB{
		host:     host,
		user:     user,
		password: password,
		instance: instance,
	}, nil
}
