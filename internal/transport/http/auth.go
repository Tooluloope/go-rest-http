package http

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(
	original func(http.ResponseWriter, *http.Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]

		if authHeader == nil {
			fmt.Println("No Authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader[0], " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("Malformed token")
			fmt.Fprintf(w, "Unauthorized")
			return
		}

		token := bearerToken[1]

		if validateToken(token) {
			original(w, r)
		} else {
			fmt.Println("Invalid token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(token string) bool {
	var mySigningKey = []byte("goisagoodlanguage")
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok != true {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return accessToken.Valid
}
