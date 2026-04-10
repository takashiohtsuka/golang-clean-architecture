package controller

import (
	"errors"
	"net/http"
	"strconv"

	"golang-clean-architecture/pkg/apperror"
	responseWomen "golang-clean-architecture/pkg/frontend/adapter/response/women"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/inputport"
)

type womanController struct {
	womanUsecase inputport.WomanUsecase
}

type Woman interface {
	GetWomanList(c Context) error
	GetStoreWomanList(c Context) error
	GetWomanDetail(c Context) error
}

func NewWomanController(u inputport.WomanUsecase) Woman {
	return &womanController{u}
}

func (wc *womanController) GetWomanList(ctx Context) error {
	women, err := wc.womanUsecase.GetList(ctx.Request().Context(), input.GetWomanListInput{})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return ctx.JSON(http.StatusOK, responseWomen.NewListResponse(women.All()))
}

func (wc *womanController) GetStoreWomanList(ctx Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	women, err := wc.womanUsecase.GetStoreWomanList(ctx.Request().Context(), input.GetStoreWomanListInput{
		StoreID: uint(id),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return ctx.JSON(http.StatusOK, responseWomen.NewListResponse(women.All()))
}

func (wc *womanController) GetWomanDetail(ctx Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	woman, err := wc.womanUsecase.GetDetail(ctx.Request().Context(), input.GetWomanDetailInput{
		WomanID: uint(id),
	})
	if err != nil {
		var nfe *apperror.NotFoundException
		if errors.As(err, &nfe) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": nfe.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDetailResponse(woman))
}
