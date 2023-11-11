package value

import (
	"regexp"

	"github.com/cockroachdb/errors"
)

type AccountName struct {
	name string
}

const (
	reName string = `[A-Za-z0-9\-]`
)

const (
	errNameLength string = "invalid name length: %v"
)

func NewAccountName(name string) (*AccountName, error) {
	length := len(name)
	if length == 0 || length > 64 || !regexp.MustCompile(reName).Match([]byte(name)) {
		return nil, errors.Newf(errNameLength, name)
	}

	return &AccountName{
		name: name,
	}, nil
}

func (a *AccountName) Value() string {
	return a.name
}
