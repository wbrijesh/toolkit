package middleware

import (
	"context"
	"errors"
	"net/http"

	"brijesh.dev/toolkit/buid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := buid.GenerateBUID()

		ctx := context.WithValue(r.Context(), "requestID", requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(r *http.Request) string {
	requestID := r.Context().Value("requestID").(string)
	if requestID == "" {
		panic(errors.New("calling GetRequestID without using RequestIDMiddleware"))
	}
	return requestID
}
