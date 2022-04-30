package relationship

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common/middleware"
)

func Routes(e *echo.Echo) {
	r := e.Group("/")

	r.Use(middleware.AuthMiddleware())
	r.Use(middleware.UserMiddleware())

	r.POST("follow/:Username", follow)
	r.DELETE("unfollow/:Username", unfollow)
}
