package middleware

import (
	"net/http"
	"time"

	"github.com/plasmatrip/muslib/internal/logger"
)

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.responseData.status = status
}

func WithLogging(log logger.Logger) func(next http.Handler) http.Handler {
	log.Sugar.Debug("handler logging started")

	return func(next http.Handler) http.Handler {
		start := time.Now()

		fn := func(w http.ResponseWriter, r *http.Request) {
			responseData := &responseData{
				status: 0,
				size:   0,
			}
			lw := loggingResponseWriter{
				ResponseWriter: w,
				responseData:   responseData,
			}

			next.ServeHTTP(&lw, r)

			duration := time.Since(start)

			var logMsg []interface{}

			switch r.Method {
			case http.MethodGet:
				logMsg = append(logMsg, "URI", r.RequestURI, "  METHOD:", r.Method, "  DURATION:", duration)
			case http.MethodPost:
				logMsg = append(logMsg, "URI", r.RequestURI, "  METHOD:", r.Method,
					"  DURATION:", duration, "  STATUS", responseData.status, "  SIZE", responseData.size)
			}

			log.Sugar.Infoln(logMsg...)
		}
		return http.HandlerFunc(fn)
	}

}
