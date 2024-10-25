package jwt

import (
	"net/http"

	"github.com/go-chi/jwtauth"

	"gomora/interfaces/http/rest/viewmodels"
	"gomora/internal/errors"
)

// JWTAuthCustomMiddleware handles JWT authentication custom errors
func JWTAuthCustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			var httpCode int
			var errorMsg string

			switch {
			case err == jwtauth.ErrExpired:
				httpCode = http.StatusUnauthorized
				errorMsg = "Token has expired."
			case err == jwtauth.ErrNoTokenFound:
				httpCode = http.StatusUnauthorized
				errorMsg = "No token found."
			case err == jwtauth.ErrUnauthorized:
				httpCode = http.StatusUnauthorized
				errorMsg = "Invalid token."
			default:
				httpCode = http.StatusUnauthorized
				errorMsg = "Invalid token"
			}

			response := viewmodels.HTTPResponseVM{
				Status:    httpCode,
				Success:   false,
				Message:   errorMsg,
				ErrorCode: errors.UnauthorizedAccess,
			}

			response.JSON(w)
			return
		}

		// If token is valid, proceed to the next handler
		if claims != nil {
			next.ServeHTTP(w, r)
		}
	})
}
