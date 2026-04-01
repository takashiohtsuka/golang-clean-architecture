package interactor

import (
	"fmt"
	"time"
)

type URLDownloadSequentialUsecase struct{}

func NewURLDownloadSequentialUsecase() *URLDownloadSequentialUsecase {
	return &URLDownloadSequentialUsecase{}
}

func (u *URLDownloadSequentialUsecase) Run() string {
	before := time.Now()

	for i := 1; i <= 100; i++ {
		url := fmt.Sprintf("https://example.com/api/users?id=%d", i)
		u.downloadJSON(url)
	}

	elapsed := time.Since(before)
	result := fmt.Sprintf("sequential elapsed: %v", elapsed)
	fmt.Println(result)
	return result
}

func (u *URLDownloadSequentialUsecase) downloadJSON(url string) {
	fmt.Println(url)
	time.Sleep(time.Second * 2)
}
