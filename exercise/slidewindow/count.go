package sliding_window_counter

import (
	"github.com/pkg/errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	once             sync.Once
	ErrExceededLimit = errors.New("ErrExceededLimit")
)

//将滑动窗口划分为多个时间片；
//在每个区间内每有一次请求就将计数器加一维持一个时间窗口,占据多个区间；
//每经过一个区间的时间,则抛弃最老的一个区间,并纳入最新的一个区间；
//如果当前窗口内区间的请求计数总和超过了限制数量,则本窗口内所有的请求都被丢弃。
type slideCount struct {
	partRequests          int32         //时间片上的请求数
	durationRequests      chan int32    //个数是时间片的个数
	timePart              time.Duration //时间划分的区间(时间片）
	timeWindow            time.Duration //时间窗口
	windowCurrentRequests int32         //时间窗口上的请求数
	allowRequests         int32         //允许的请求数
}

// 初始化对象
func New(accuracy time.Duration, snippet time.Duration, allowRequests int32) *slideCount {
	return &slideCount{durationRequests: make(chan int32, snippet/accuracy), timePart: accuracy, timeWindow: snippet, allowRequests: allowRequests}
}

// 请求
func (l *slideCount) Take() error {
	once.Do(func() {
		go sliding(l)
		go calculate(l)
	})
	curRequest := atomic.LoadInt32(&l.windowCurrentRequests)
	if curRequest >= l.allowRequests {
		return ErrExceededLimit
	}
	if !atomic.CompareAndSwapInt32(&l.windowCurrentRequests, curRequest, curRequest+1) {
		return ErrExceededLimit
	}
	atomic.AddInt32(&l.partRequests, 1)
	return nil

}

// 窗口滑动
func sliding(l *slideCount) {
	for {
		select {
		// 每经过一个时间窗口，重新计数
		case <-time.After(l.timePart):
			t := atomic.SwapInt32(&l.partRequests, 0)
			l.durationRequests <- t
		}
	}
}

// 滑动计算
func calculate(l *slideCount) {
	for {
		<-time.After(l.timePart)
		if len(l.durationRequests) == cap(l.durationRequests) {
			break
		}
	}
	for {
		<-time.After(l.timePart)
		t := <-l.durationRequests
		if t != 0 {
			atomic.AddInt32(&l.windowCurrentRequests, -t)
		}
	}
}
