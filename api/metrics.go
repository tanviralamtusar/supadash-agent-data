package api

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// metricsData holds basic application metrics
type metricsData struct {
	startTime time.Time
	requests  int64
}

var metrics = &metricsData{
	startTime: time.Now(),
}

// getMetrics returns application metrics in a Prometheus-compatible text format
func (a *Api) getMetrics(c *gin.Context) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	uptime := time.Since(metrics.startTime).Seconds()

	c.Header("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
	c.String(http.StatusOK, `# HELP supadash_uptime_seconds Time since the API server started.
# TYPE supadash_uptime_seconds gauge
supadash_uptime_seconds %.2f
# HELP supadash_go_goroutines Number of goroutines currently running.
# TYPE supadash_go_goroutines gauge
supadash_go_goroutines %d
# HELP supadash_go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE supadash_go_memstats_alloc_bytes gauge
supadash_go_memstats_alloc_bytes %d
# HELP supadash_go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE supadash_go_memstats_sys_bytes gauge
supadash_go_memstats_sys_bytes %d
# HELP supadash_go_gc_completed_total Number of completed GC cycles.
# TYPE supadash_go_gc_completed_total counter
supadash_go_gc_completed_total %d
`,
		uptime,
		runtime.NumGoroutine(),
		memStats.Alloc,
		memStats.Sys,
		memStats.NumGC,
	)
}
