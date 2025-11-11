package auth

import (
	"net/http"
	"os"
	"time"
)

// ClearJWTCookie clears the JWT cookie
func ClearJWTCookie(w http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	// enforce domain for production
	if os.Getenv("API_ENV") == "production" {
		cookie.Domain = ".nuxify.tech" // allow access to all subdomains only
	}

	http.SetCookie(w, cookie)
}

// SetJWTCookie sets the JWT cookie
func SetJWTCookie(w http.ResponseWriter, cookieName, token string, expiresAt time.Time) {
	cookie := &http.Cookie{
		Name:     cookieName, // required by jwtauth.Verifier
		Value:    token,
		Path:     "/",
		Expires:  expiresAt,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	// enforce domain for production
	if os.Getenv("API_ENV") == "production" {
		cookie.Domain = ".nuxify.tech" // allow access to all subdomains only
	}

	http.SetCookie(w, cookie)
}
