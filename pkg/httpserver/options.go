package httpserver

import "time"

// Option is a functional option type
type Option func(*Server)

// Port sets the address
func Port(port string) Option {
	return func(s *Server) {
		s.server.Addr = ":" + port
	}
}

// ReadTimeout sets the ReadTimeout
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout sets the WriteTimeout
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// IdleTimeout sets the IdleTimeout
func IdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.IdleTimeout = timeout
	}
}

// ShutdownTimeout sets the Graceful Shutdown timeout
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
