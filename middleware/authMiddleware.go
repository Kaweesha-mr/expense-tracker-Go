package middleware

import (
	"Expense-Tracker-go/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserClaimsKey = contextKey("userClaims")

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Get the authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		info, err := utils.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid or Expired token", http.StatusUnauthorized)
			return
		}

		//add claims to the request context
		ctx := context.WithValue(r.Context(), UserClaimsKey, info)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
