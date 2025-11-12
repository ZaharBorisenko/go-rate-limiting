package limiter

import (
	"container/list"
	"time"
)

type RateLimiter interface {
	Allow() bool
}

var _ RateLimiter = (*SlidingWindowLimiter)(nil)

type SlidingWindowLimiter struct {
	window int64
	limit  int
	logs   *list.List
}

func NewSlidingWindowLimiter(window int64, limit int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		window: window,
		limit:  limit,
		logs:   list.New(),
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	currentTime := time.Now()
	delta := currentTime.Unix() - s.window
	edgeTime := time.Unix(delta, 0) //Everything that happened before edgeTime is deleted, because these requests have already gone out of the window and should not be taken into account when limiting.

	//remove old request
	for s.logs.Len() > 0 {
		front := s.logs.Front()
		if front.Value.(time.Time).Before(edgeTime) {
			s.logs.Remove(front)
		} else {
			break
		}
	}

	//check if we can accept this request
	if s.logs.Len() < s.limit {
		s.logs.PushBack(currentTime)
		return true
	}

	return false
}
