package env

import "github.com/cockroachdb/errors"

func (r *Reader) server() (*Server, error) {
	host, err := r.read("SERVER_HOST")
	if err != nil {
		return nil, errors.Errorf(errMsgEnv, err)
	}

	return &Server{
		host: host,
	}, nil
}

type Server struct {
	host string
}

func (s *Server) Host() string {
	return s.host
}
