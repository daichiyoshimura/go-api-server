package auth

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SiginingKey(secret string) []byte {
	return []byte(secret)
}

func JWT(signingKey []byte) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			return c.Request().RequestURI == "/auth/signin"
		},
		SigningKey: signingKey,
	})
}
