package main

import (
	"errors"
	"fmt"
	"github.com/lsmoura/go-fullstack/internal/server"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func addr() string {
	host := ""
	port := "8080"
	if hostEnvValue := os.Getenv("HOST"); hostEnvValue != "" {
		host = hostEnvValue
	}
	if portEnvValue := os.Getenv("PORT"); portEnvValue != "" {
		port = portEnvValue
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("fullstack server starting")

	s := server.Server{
		Logger: logger,
	}

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-done
		if err := s.Close(); err != nil {
			logger.Error().Err(err).Msg("error closing server")
		}
	}()

	if err := s.Start(addr()); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal().Err(err).Msg("server error")
		}

		logger.Info().Msg("server closed")
	}
}
