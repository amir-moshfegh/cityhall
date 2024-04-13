package router

import (
	"github.com/amir-moshfegh/cityhall/bootstrap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func New() *echo.Echo {
	e := echo.New()
	//e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}

func Setup(app *bootstrap.Application, r *echo.Group, timeOut time.Duration) {

	publicRouter := r.Group("/base")
	BaseRoute(app, publicRouter, timeOut)

}
