package types

type JWTCookieName string

const (
	DefaultCookieName JWTCookieName = "jwt"
	AdminCookieName   JWTCookieName = "jwt-admin"
)
