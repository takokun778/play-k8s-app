package log

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Log() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		r = r.WithContext(SetLogCtx(r.Context()))

		lrw := &loggingResponseWriter{w, http.StatusOK}

		next.ServeHTTP(lrw, r)

		log := GetLogCtx(r.Context())

		elapsed := time.Since(start)

		log.Sugar().Infof("%s %s %s %d %s", r.Proto, r.Method, r.URL, lrw.statusCode, elapsed)
	})
}
