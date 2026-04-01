package interactor

import (
	"fmt"
	"sync"
	"time"

	"golang-clean-architecture/pkg/domain/structure/concurrent"
)

const maxConcurrentDownload = 20

type URLDownloadConcurrentUsecase struct {
	urlDownload *concurrent.Limiter
}

func NewURLDownloadConcurrentUsecase() *URLDownloadConcurrentUsecase {
	return &URLDownloadConcurrentUsecase{
		urlDownload: concurrent.NewLimiter(maxConcurrentDownload),
	}
}

func (u *URLDownloadConcurrentUsecase) Run() string {
	before := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		u.urlDownload.ConcurrentCh <- struct{}{}
		wg.Add(1)
		go func(id int) {
			defer func() {
				<-u.urlDownload.ConcurrentCh
				wg.Done()
			}()
			url := fmt.Sprintf("https://example.com/api/users?id=%d", id)
			u.downloadJSON(url)
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(before)
	result := fmt.Sprintf("concurrent elapsed: %v", elapsed)
	fmt.Println(result)
	return result
}

func (u *URLDownloadConcurrentUsecase) downloadJSON(url string) {
	fmt.Println(url)
	time.Sleep(time.Second * 2)
}
