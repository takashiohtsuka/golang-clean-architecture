package controller

import (
	"net/http"
	"strconv"

	requestStores "golang-clean-architecture/pkg/backend/adapter/request/stores"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
)

type storeController struct {
	storeUsecase inputport.StoreUsecase
}

type Store interface {
	CreateStore(c Context) error
	UpdateStore(c Context) error
}

func NewStoreController(u inputport.StoreUsecase) Store {
	return &storeController{u}
}

func (sc *storeController) CreateStore(ctx Context) error {
	var req requestStores.Post
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := sc.storeUsecase.Create(ctx.Request().Context(), req.ToInput()); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (sc *storeController) UpdateStore(ctx Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req requestStores.Put
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := sc.storeUsecase.Update(ctx.Request().Context(), req.ToInput(uint(id))); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}
