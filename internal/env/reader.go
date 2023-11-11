package env

import (
	"os"

	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

const (
	errMsgEnv   string = "env: %w"
	errMsgEmpty string = "ENV must be set : %v"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read() (*Server, *DB, error) {
	stg, err := NewStage(os.Getenv("STAGE"))
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if stg.isDev() || stg.isTest() {
		if err := godotenv.Load("~/go/src/dev/go-api-server/.env"); err != nil { // FIXME unexpected behaivor 
			return nil, nil, errors.Newf(errMsgEnv, err)
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
		return val, errors.Newf(errMsgEmpty, key)
	}

	return val, nil
}
