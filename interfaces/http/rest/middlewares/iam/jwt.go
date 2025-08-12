package jwt

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"

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

// TODO: pass on r.Context
// Verifier override jwtauth Verifier
func Verifier(ja *jwtauth.JWTAuth, cookieKey *string) func(http.Handler) http.Handler {
	random := "sample data"
	tokenFromCookieWithKey := func(r *http.Request) string {
		log.Println(random)

		key := "jwt" // default key
		if cookieKey != nil {
			key = *cookieKey
		}

		cookie, err := r.Cookie(key)
		if err != nil {
			return ""
		}

		return cookie.Value
	}

	return jwtauth.Verify(ja, jwtauth.TokenFromHeader, tokenFromCookieWithKey)
}
