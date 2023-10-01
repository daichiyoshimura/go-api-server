package env

type Server struct {
	host string
}

func (s *Server) Host() string {
	return s.host
}
