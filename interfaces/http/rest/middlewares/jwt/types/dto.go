package types

type JWTCookieName string
type ctxKey int64

const (
	DefaultCookieName JWTCookieName = "jwt"
	AdminCookieName   JWTCookieName = "jwt-admin"
)

const (
	UserIDCtx ctxKey = iota
)
