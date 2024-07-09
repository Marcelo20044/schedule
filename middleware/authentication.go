package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"schedule/internal/presentation/utils"
	"strings"
	"time"
)

var jwtSecret = []byte("GoFuckYourself")

type Claims struct {
	Username string   `json:"sub"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.Response(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			utils.Response(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			utils.Response(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims.ExpiresAt.Before(time.Now()) {
			utils.Response(w, "Token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "roles", claims.Roles)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
