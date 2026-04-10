package controller

import (
	"net/http"
	"strconv"

	responseStores "golang-clean-architecture/pkg/frontend/adapter/response/stores"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/inputport"
)

type storeController struct {
	storeUsecase inputport.StoreUsecase
}

type Store interface {
	GetStoreDetail(c Context) error
}

func NewStoreController(u inputport.StoreUsecase) Store {
	return &storeController{u}
}

func (sc *storeController) GetStoreDetail(ctx Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	store, err := sc.storeUsecase.GetDetail(ctx.Request().Context(), input.GetStoreDetailInput{
		StoreID: uint(id),
	})
	if err != nil {
		return err
	}

	if store.IsNil() {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	return ctx.JSON(http.StatusOK, responseStores.NewDetailResponse(store))
}
