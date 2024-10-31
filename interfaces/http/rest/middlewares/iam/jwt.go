package jwt

import (
	"net/http"

	"github.com/go-chi/jwtauth"

	"gomora/interfaces/http/rest/viewmodels"
	"gomora/internal/errors"
)

// JWTAuthMiddleware handles JWT authentication custom errors
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			var httpCode int
			var errorMsg string

			switch err {
			case jwtauth.ErrExpired:
				httpCode = http.StatusUnauthorized
				errorMsg = "Token has expired."
			case jwtauth.ErrNoTokenFound:
				httpCode = http.StatusUnauthorized
				errorMsg = "No token found."
			case jwtauth.ErrUnauthorized:
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

		// if token is nil, creates unauthorized response
		if claims == nil {
			response := viewmodels.HTTPResponseVM{
				Status:    http.StatusUnauthorized,
				Success:   false,
				Message:   "Invalid token",
				ErrorCode: errors.UnauthorizedAccess,
			}

			response.JSON(w)
			return
		}

		// if token is valid, proceeds to the next handler
		next.ServeHTTP(w, r)
	})
}
