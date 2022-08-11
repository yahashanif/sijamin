package middleware

import "github.com/labstack/echo/v4/middleware"

var IsAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
