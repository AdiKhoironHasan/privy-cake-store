package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Middleware struct {
}

func NewMidleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization,Origin,Accept,datetime,signature,Content-Type")
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}

func (m Middleware) LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, ip=${remote_ip}, latency=${latency_human}\n",
	}))
}
