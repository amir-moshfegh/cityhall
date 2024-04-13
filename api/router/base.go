package router

import (
	"github.com/amir-moshfegh/cityhall/api/controller"
	"github.com/amir-moshfegh/cityhall/bootstrap"
	"github.com/amir-moshfegh/cityhall/repository"
	"github.com/amir-moshfegh/cityhall/usecase"
	"github.com/labstack/echo/v4"
	"time"
)

func BaseRoute(app *bootstrap.Application, r *echo.Group, timeOut time.Duration) {
	br := repository.NewBaseRepository(app.DB)
	bc := controller.BaseController{
		BaseUseCase: usecase.NewBaseUseCase(br, timeOut),
	}

	r.POST("/create", bc.Create)
	r.PUT("/update", bc.Update)
	r.DELETE("/delete/:id", bc.Delete)
	r.GET("/find/:id", bc.FindByID)
	r.GET("/find-by-category/:id", bc.FindByCategory)
}
