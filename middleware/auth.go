package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user_claims"}

type UserClaims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Bearer token required", http.StatusUnauthorized)
			return
		}

		jwtSecret := []byte(os.Getenv("NEXTAUTH_SECRET"))
		if len(jwtSecret) == 0 {
			http.Error(w, "JWT Secret not configured", http.StatusInternalServerError)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(*UserClaims); ok {
			ctx := context.WithValue(r.Context(), UserCtxKey, claims)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
