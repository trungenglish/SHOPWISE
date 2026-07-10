package middleware

import (
	"net/http"
	"sync"
	"time"

	"shopwise/retail/internal/platform/apperror"

	"github.com/gin-gonic/gin"
)

type rateBucket struct {
	count     int
	resetTime time.Time
}

type IPRateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	buckets map[string]*rateBucket
}

func NewIPRateLimiter(limit int, window time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		limit:   limit,
		window:  window,
		buckets: map[string]*rateBucket{},
	}
}

func (l *IPRateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		l.mu.Lock()
		bucket, ok := l.buckets[ip]
		if !ok || now.After(bucket.resetTime) {
			bucket = &rateBucket{count: 0, resetTime: now.Add(l.window)}
			l.buckets[ip] = bucket
		}
		bucket.count++
		allowed := bucket.count <= l.limit
		l.mu.Unlock()

		if !allowed {
			_ = c.Error(apperror.TooManyRequests("rate limit exceeded", nil))
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}
