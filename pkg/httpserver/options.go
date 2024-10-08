package httpserver

import (
	"net"
	"time"
)

type Options func(s *Server)

func ReadTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func Port(port string) Options {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

func ShutdownTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
