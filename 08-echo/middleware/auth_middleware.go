package middleware

import (
	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.NewHTTPError(401, "Please provide valid credentials")
			}
			// 验证 token 的逻辑
			return next(c)
		}
	}
}
