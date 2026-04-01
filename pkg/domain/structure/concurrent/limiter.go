package concurrent

type Limiter struct {
	ConcurrentCh chan struct{}
}

func NewLimiter(maxConcurrent int) *Limiter {
	return &Limiter{
		ConcurrentCh: make(chan struct{}, maxConcurrent),
	}
}
