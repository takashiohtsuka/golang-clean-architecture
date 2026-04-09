package controller

import (
	"net/http"
	"strconv"

	requestWomen "golang-clean-architecture/pkg/backend/adapter/request/women"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
)

type womanController struct {
	womanUsecase inputport.WomanUsecase
}

type Woman interface {
	CreateWoman(c Context) error
	UpdateWoman(c Context) error
}

func NewWomanController(u inputport.WomanUsecase) Woman {
	return &womanController{u}
}

func (wc *womanController) UpdateWoman(ctx Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req requestWomen.Put
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	file, header, _ := ctx.Request().FormFile("image")
	if file != nil {
		defer file.Close()
	}

	if err := wc.womanUsecase.Update(ctx.Request().Context(), req.ToInput(uint(id), file, header)); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (wc *womanController) CreateWoman(ctx Context) error {
	var req requestWomen.Post
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := wc.womanUsecase.Create(ctx.Request().Context(), req.ToInput()); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}
