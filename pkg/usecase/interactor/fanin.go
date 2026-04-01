package interactor

import (
	"context"
	"fmt"
	"time"

	"golang-clean-architecture/pkg/domain/structure/concurrent"
)

const maxConcurrent = 20

type FanInUsecase struct {
	fanIn *concurrent.Limiter
}

func NewFanInUsecase() *FanInUsecase {
	return &FanInUsecase{
		fanIn: concurrent.NewLimiter(maxConcurrent),
	}
}

func (f *FanInUsecase) Run() []string {
	// 空きスロットを1つ確保（20本埋まっていたら空くまで待機）
	f.fanIn.ConcurrentCh <- struct{}{}
	fmt.Printf("goroutine started: 現在 %d 本\n", len(f.fanIn.ConcurrentCh))
	defer func() {
		// 処理完了後にスロットを解放
		<-f.fanIn.ConcurrentCh
		fmt.Printf("goroutine stopped: 現在 %d 本\n", len(f.fanIn.ConcurrentCh))
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := fanIn(ctx, generator(ctx, "Hello"), generator(ctx, "Bye"))
	results := make([]string, 100)
	for i := 0; i < 100; i++ {
		results[i] = <-ch
	}
	return results
}

func fanIn(ctx context.Context, ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("fanIn ch1: goroutine stopped")
				return
			case val := <-ch1:
				newCh <- val
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("fanIn ch2: goroutine stopped")
				return
			case val := <-ch2:
				newCh <- val
			}
		}
	}()
	return newCh
}

func generator(ctx context.Context, msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("generator(%s): goroutine stopped\n", msg)
				return
			case ch <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}
