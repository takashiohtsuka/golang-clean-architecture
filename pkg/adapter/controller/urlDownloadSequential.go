package controller

import (
	"net/http"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type URLDownloadSequentialUsecase interface {
	Run() string
}

type urlDownloadSequentialController struct {
	urlDownloadSequentialUsecase URLDownloadSequentialUsecase
}

type URLDownloadSequential interface {
	GetURLDownloadSequential(c Context) error
}

func NewURLDownloadSequentialController(u URLDownloadSequentialUsecase) URLDownloadSequential {
	return &urlDownloadSequentialController{u}
}

func (uc *urlDownloadSequentialController) GetURLDownloadSequential(ctx Context) error {
	result := uc.urlDownloadSequentialUsecase.Run()
	return ctx.JSON(http.StatusOK, result)
}
