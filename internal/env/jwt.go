package env

import "github.com/cockroachdb/errors"

func (r *Reader) jwt() (*JWT, error) {
	secret, err := r.read("JWT_SECRET")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}
	return &JWT{
		secret: secret,
	}, nil
}

type JWT struct {
	secret string
}

func (j *JWT) Secret() string {
	return j.secret
}
