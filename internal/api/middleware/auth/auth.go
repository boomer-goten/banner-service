package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var ContextRoleKey = ContextKey("role")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenParam := r.Header.Get("token")
		if r.URL.Path == "/token" {
			w.WriteHeader(http.StatusCreated)
			next.ServeHTTP(w, r)
			return
		}
		token, err := jwt.Parse(tokenParam, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextRoleKey, role)
		r = r.WithContext(ctx)
		if (role == "user" && r.Method == "GET") || (role == "admin") {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	})
}
