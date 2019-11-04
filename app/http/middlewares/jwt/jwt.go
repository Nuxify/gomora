package jwt

import (
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kabaluyot/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

//Authenticator is built-in middleware that authenticates the given token
func Authenticator() func(http.Handler) http.Handler {
	return jwtauth.Authenticator
}

//GetClaims returns the claim in a map[string]interface{} type
func GetClaims(r *http.Request) jwt.MapClaims {
	_, claims, _ := jwtauth.FromContext(r.Context())

	return claims
}

//Verifier returns the JWT Auth token created from the JWT_SECRET_KEY
func Verifier() func(http.Handler) http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)

	return jwtauth.Verifier(tokenAuth)
}
