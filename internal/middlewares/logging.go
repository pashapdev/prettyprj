package middlewares

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"prettyprj/internal/logger"
	"prettyprj/internal/utils"

	"github.com/google/uuid"
)

var privatePaths map[string]struct{} = map[string]struct{}{
	"/readyz":  {},
	"/healthz": {},
	"/metrics": {},
}

func MakeLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(string(utils.HeaderRequestID))
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx := context.WithValue(r.Context(), utils.HeaderRequestID, requestID)
		w.Header().Set(string(utils.HeaderRequestID), requestID)

		if _, exist := privatePaths[r.URL.Path]; !exist {
			defer logRequest(ctx, "finish request", r)
			logRequest(ctx, "start request", r)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func logRequest(_ context.Context, message string, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))
	logger.Info(
		message,
		"method", r.Method,
		"uri", r.URL.String(),
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
		"headers", r.Header,
		"request_body_size", len(body),
		"request_body", body)
}
