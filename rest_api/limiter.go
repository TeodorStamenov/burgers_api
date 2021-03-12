package rest_api

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var visitors = make(map[string]*CustomLimiter)
var mu sync.Mutex

// CustomLimiter struct
type CustomLimiter struct {
	maxCalls     int
	recieveCalls int
	limiter      *rate.Limiter
	firstReq     time.Time
}

func NewCustomLimiter() *CustomLimiter {
	return &CustomLimiter{
		maxCalls:     3600,
		recieveCalls: 0,
		limiter:      rate.NewLimiter(1, 1),
		firstReq:     time.Now(),
	}
}

func (c CustomLimiter) GetLimit() int {
	return c.maxCalls
}

func (c CustomLimiter) GetRestCalls() int {
	return c.maxCalls - c.recieveCalls
}

func (c CustomLimiter) Allow() bool {
	return c.limiter.Allow() && c.GetRestCalls() > 0
}

func (c *CustomLimiter) update() {
	if time.Now().Second() > (c.firstReq.Second() + c.maxCalls) {
		c.reset()
	}
	if c.GetRestCalls() > 0 {
		c.recieveCalls++
	}
}

func (c *CustomLimiter) reset() {
	c.firstReq = time.Now()
	c.recieveCalls = 0
}

// GetVisitor function
func GetVisitor(ip string) CustomLimiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		// rt := rate.Every(1 * time.Hour / 3600)
		limiter = NewCustomLimiter()
		visitors[ip] = limiter
	}

	limiter.update()

	return *limiter
}
