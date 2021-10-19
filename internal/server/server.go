package server

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() error {

	return s.httpServer.ListenAndServe()
}
