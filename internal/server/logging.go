package server

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func logger(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := zerolog.Ctx(r.Context())

		event := logger.Info().
			Str("method", r.Method).
			Str("url", r.URL.String())

		start := time.Now()
		next(w, r)
		event.Str("status", w.Header().Get("Status")).
			Str("duration", time.Since(start).String()).
			Msg("request")
	}
}

func loggerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(logger(next.ServeHTTP))
}
