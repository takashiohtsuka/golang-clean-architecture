package controller

import (
	"net/http"
)

// 使う側（controller）がusecaseに必要なインターフェースを定義
type URLDownloadConcurrentUsecase interface {
	Run() string
}

type urlDownloadConcurrentController struct {
	urlDownloadConcurrentUsecase URLDownloadConcurrentUsecase
}

type URLDownloadConcurrent interface {
	GetURLDownloadConcurrent(c Context) error
}

func NewURLDownloadConcurrentController(u URLDownloadConcurrentUsecase) URLDownloadConcurrent {
	return &urlDownloadConcurrentController{u}
}

func (uc *urlDownloadConcurrentController) GetURLDownloadConcurrent(ctx Context) error {
	result := uc.urlDownloadConcurrentUsecase.Run()
	return ctx.JSON(http.StatusOK, result)
}
