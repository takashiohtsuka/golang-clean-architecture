package controller

import (
	"net/http"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type FanInUsecase interface {
	Run() []string
}

type fanInController struct {
	fanInUsecase FanInUsecase
}

type FanIn interface {
	GetFanIn(c Context) error
}

func NewFanInController(f FanInUsecase) FanIn {
	return &fanInController{f}
}

func (fc *fanInController) GetFanIn(ctx Context) error {
	results := fc.fanInUsecase.Run()
	return ctx.JSON(http.StatusOK, results)
}
