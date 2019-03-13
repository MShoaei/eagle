package middlewares

import (
	"context"
	"net/http"
)

type AuthToken string

func GetAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := AuthToken("Authorization")
		ctx := context.WithValue(r.Context(), authorization, r.Header.Get("Authorization"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
