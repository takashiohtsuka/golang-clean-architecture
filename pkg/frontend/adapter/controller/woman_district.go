package controller

import (
	"net/http"
	"strconv"

	responseWomen "golang-clean-architecture/pkg/frontend/adapter/response/women"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/inputport"
)

type womanDistrictController struct {
	womanDistrictUsecase inputport.WomanDistrictUsecase
}

type WomanDistrict interface {
	GetWomanDistrictList(c Context) error
}

func NewWomanDistrictController(u inputport.WomanDistrictUsecase) WomanDistrict {
	return &womanDistrictController{u}
}

func (wc *womanDistrictController) GetWomanDistrictList(ctx Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	women, err := wc.womanDistrictUsecase.GetList(ctx.Request().Context(), input.GetWomanDistrictListInput{
		DistrictID: uint(id),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDistrictListResponse(women.All()))
}
