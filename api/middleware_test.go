package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestRateLimitMiddleware(t *testing.T) {
	t.Run("allows requests under limit", func(t *testing.T) {
		router := gin.New()
		router.Use(RateLimitMiddleware(10, 10)) // 10 req/sec, burst 10
		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "192.168.1.1:12345"
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("blocks requests over limit", func(t *testing.T) {
		router := gin.New()
		// Very restrictive: 1 req/sec, burst of 1
		router.Use(RateLimitMiddleware(1, 1))
		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// First request should pass
		w1 := httptest.NewRecorder()
		req1, _ := http.NewRequest("GET", "/test", nil)
		req1.RemoteAddr = "10.0.0.1:12345"
		router.ServeHTTP(w1, req1)
		assert.Equal(t, http.StatusOK, w1.Code)

		// Second request immediately should be rate limited
		w2 := httptest.NewRecorder()
		req2, _ := http.NewRequest("GET", "/test", nil)
		req2.RemoteAddr = "10.0.0.1:12345"
		router.ServeHTTP(w2, req2)
		assert.Equal(t, http.StatusTooManyRequests, w2.Code)
	})

	t.Run("different IPs have separate limits", func(t *testing.T) {
		router := gin.New()
		router.Use(RateLimitMiddleware(1, 1))
		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// IP A — first request
		wA := httptest.NewRecorder()
		reqA, _ := http.NewRequest("GET", "/test", nil)
		reqA.RemoteAddr = "10.0.0.100:12345"
		router.ServeHTTP(wA, reqA)
		assert.Equal(t, http.StatusOK, wA.Code)

		// IP B — first request (should pass, different limiter)
		wB := httptest.NewRecorder()
		reqB, _ := http.NewRequest("GET", "/test", nil)
		reqB.RemoteAddr = "10.0.0.200:12345"
		router.ServeHTTP(wB, reqB)
		assert.Equal(t, http.StatusOK, wB.Code)

		// IP A — second request (should be limited)
		wA2 := httptest.NewRecorder()
		reqA2, _ := http.NewRequest("GET", "/test", nil)
		reqA2.RemoteAddr = "10.0.0.100:12345"
		router.ServeHTTP(wA2, reqA2)
		assert.Equal(t, http.StatusTooManyRequests, wA2.Code)
	})
}

func TestRateLimiterEntry(t *testing.T) {
	rl := newIPRateLimiter(10, 10)

	t.Run("creates limiter for new IP", func(t *testing.T) {
		l := rl.getLimiter("1.2.3.4")
		assert.NotNil(t, l)
	})

	t.Run("returns same limiter for same IP", func(t *testing.T) {
		l1 := rl.getLimiter("5.6.7.8")
		l2 := rl.getLimiter("5.6.7.8")
		assert.Equal(t, l1, l2)
	})

	t.Run("returns different limiters for different IPs", func(t *testing.T) {
		rl2 := newIPRateLimiter(1, 1) // 1 req/sec, burst 1
		l1 := rl2.getLimiter("10.0.0.1")
		l2 := rl2.getLimiter("10.0.0.2")
		// Exhaust l1's token
		l1.Allow()
		// l2 should still have its own token since it's a separate limiter
		assert.True(t, l2.Allow(), "different IPs should have independent limiters")
	})
}
