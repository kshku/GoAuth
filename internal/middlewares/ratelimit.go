package middlewares

import (
	"net/http"
	"net"
	"sync"
	"time"

	"go_auth/internal/utils"
)

type rateLimiter struct {
	visitors map[string]*visitor
	mu sync.RWMutex
	rate int
	burst int
	cleanup time.Duration
}

type visitor struct {
	tokens int
	lastSeen time.Time
	lastRefill time.Time
}

var limiter *rateLimiter

func init() {
	limiter = &rateLimiter{
		visitors: make(map[string]*visitor),
		rate: 60,
		burst: 10,
		cleanup: time.Minute,
	}

	go limiter.cleanupVisitors()
}

func (rl *rateLimiter) getVisitor(ip string) *visitor {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	v, exists := rl.visitors[ip]
	if !exists {
		rl.visitors[ip] = &visitor{
			tokens: rl.burst,
			lastSeen: time.Now(),
			lastRefill : time.Now(),
		}
	}

	return v
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		rl.visitors[ip] = &visitor{
			tokens: rl.burst - 1,
			lastSeen: time.Now(),
			lastRefill : time.Now(),
		}
		return true
	}

	v.lastSeen = time.Now()

	now := time.Now()
	elapsed := now.Sub(v.lastRefill)
	tokensToAdd := int(elapsed.Minutes() * float64(rl.rate))

	if tokensToAdd > 0 {
		v.tokens += tokensToAdd
		if v.tokens > rl.burst {
			v.tokens = rl.burst
		}
		v.lastRefill = now
	}

	if v.tokens > 0 {
		v.tokens--
		return true
	}

	return false
}

func (rl *rateLimiter) cleanupVisitors() {
	ticker := time.NewTicker(rl.cleanup)

	for range ticker.C {
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 3 * time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			ip = r.RemoteAddr
		}

		if !limiter.allow(ip) {
			utils.WriteError(w, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}

		next.ServeHTTP(w, r)
	}
}
