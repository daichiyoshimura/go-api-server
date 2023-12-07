package model

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type uuidHelper struct{}

func newUUIDHelper() *uuidHelper {
	return &uuidHelper{}
}

func (h *uuidHelper) generate() ([]byte, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	binId, err := id.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return binId, nil
}

func (h *uuidHelper) toBinary(id string) ([]byte, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	bin, err := parsed.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bin, nil
}

func (h *uuidHelper) toString(id []byte) (string, error) {
	parsed, err := uuid.ParseBytes(id)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return parsed.String(), nil
}
