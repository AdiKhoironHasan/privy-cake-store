package http

import (
	"net/http"
	"os"
	"strconv"

	"github.com/AdiKhoironHasan/privy-cake-store/internal/services"
	"github.com/AdiKhoironHasan/privy-cake-store/internal/transport/http/middleware"
	servConst "github.com/AdiKhoironHasan/privy-cake-store/pkg/common/const"
	"github.com/AdiKhoironHasan/privy-cake-store/pkg/dto"
	servErrors "github.com/AdiKhoironHasan/privy-cake-store/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}
	middleware.NewMidleware().LogMiddleware(e)
	e.GET("/ping", handler.Ping)
	e.POST("/cakes", handler.AddNewCake)
	e.GET("/cakes", handler.ShowAllCake)
	e.GET("/cakes/:id", handler.ShowCakeByID)
	e.PATCH("/cakes/:id", handler.UpdateCake)
	e.DELETE("/cakes/:id", handler.DeleteCake)
}

func (h *HttpHandler) Ping(c echo.Context) error {
	version := os.Getenv("VERSION")

	if version == "" {
		version = "pong"
	}

	data := version

	return c.JSON(http.StatusOK, data)

}

func (h *HttpHandler) AddNewCake(c echo.Context) error {
	postDTO := dto.CakeReqDTO{}

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.AddNewCake(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: servConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) ShowAllCake(c echo.Context) error {
	getDTO := dto.CakeParamReqDTO{}

	if err := c.Bind(&getDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := getDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := h.service.ShowAllCake(&getDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: servConst.GetDataSuccess,
		Data:    result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) ShowCakeByID(c echo.Context) error {
	var cakeID int

	cakeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	result, err := h.service.ShowCakeByID(cakeID)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: servConst.GetDataSuccess,
		Data:    result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) UpdateCake(c echo.Context) error {
	var cakeID int
	patchDTO := dto.CakeReqDTO{}

	cakeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Bind(&patchDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	patchDTO.ID = cakeID

	err = patchDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.UpdateCake(&patchDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: servConst.UpdateSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)

}

func (h *HttpHandler) DeleteCake(c echo.Context) error {
	var cakeID int

	cakeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err = h.service.DeleteCake(cakeID)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: servConst.DeleteSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)

}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case servErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case servErrors.ErrNotFound:
		return http.StatusNotFound
	case servErrors.ErrConflict:
		return http.StatusConflict
	case servErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case servErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
