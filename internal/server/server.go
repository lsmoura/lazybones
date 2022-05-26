package server

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"net"
	"net/http"
)

type Server struct {
	server *http.Server

	Logger zerolog.Logger
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", logger(s.HandleFavicon))
	mux.HandleFunc("/health", logger(s.HandleHealth))
	mux.Handle("/", loggerHandler(http.FileServer(http.Dir("static/"))))

	s.server = &http.Server{
		Addr:        addr,
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return s.Logger.WithContext(context.Background()) },
	}

	s.Logger.Info().Str("addr", addr).Msg("starting server")
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("server.ListenAndServe: %w", err)
	}

	return nil
}

func (_ Server) HandleFavicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte{})
}

func (_ Server) HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func (s *Server) Close() error {
	return s.server.Close()
}
