package controller

import (
	"github.com/amir-moshfegh/cityhall/dto"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BaseController struct {
	BaseUseCase dto.BaseUseCase
}

func (bc *BaseController) Create(c echo.Context) error {
	var base dto.BaseCreateReq
	if err := c.Bind(&base); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}

	//TODO:: change validation to ozzo
	if err := c.Validate(&base); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}
	nb, err := bc.BaseUseCase.Create(c.Request().Context(), &base)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Internal server error: " + err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Created successfully",
		Data:    nb,
	})
}
func (bc *BaseController) Update(c echo.Context) error {
	var req dto.BaseUpdateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}

	//TODO:: change validation to ozzo
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}

	if _, err := bc.BaseUseCase.FindByID(c.Request().Context(), req.ID); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "this record not found!" + err.Error(),
			Data:    nil,
		})
	}

	if err := bc.BaseUseCase.Update(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Internal server error: " + err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Updated successfully",
		Data:    nil,
	})
}
func (bc *BaseController) Delete(c echo.Context) error {
	//TODO:: id convertor to uint
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}

	if _, err := bc.BaseUseCase.FindByID(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "this record not found!" + err.Error(),
			Data:    nil,
		})
	}

	if err := bc.BaseUseCase.Delete(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Internal server error: " + err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Deleted successfully",
		Data:    nil,
	})

}
func (bc *BaseController) FindByID(c echo.Context) error {
	//TODO:: id convertor to uint
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}

	nb, err := bc.BaseUseCase.FindByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Internal server error: " + err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Found successfully",
		Data:    nb,
	})
}
func (bc *BaseController) FindByCategory(c echo.Context) error {
	//TODO:: id convertor to uint
	categoryId, err := strconv.Atoi(c.Param("category"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request" + err.Error(),
			Data:    nil,
		})
	}

	nb, err := bc.BaseUseCase.FindByCategory(c.Request().Context(), uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Internal server error: " + err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Found successfully",
		Data:    nb,
	})
}
