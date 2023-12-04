package routes

import (
	"BE-Sosmed/features/postings"
	"BE-Sosmed/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uh users.Handler, ph postings.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uh)
	routePosting(e, ph)
}

func routeUser(e *echo.Echo, uh users.Handler) {
	e.POST("/users", uh.Register())
	e.POST("/login", uh.Login())
	e.GET("/users/:id", uh.ReadById(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/users", uh.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/users", uh.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func routePosting(e *echo.Echo, ph postings.Handler) {
	e.POST("/post", ph.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/post", ph.GetAll())
	e.PUT("/post/:id", ph.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/post/:id", ph.Delete(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
