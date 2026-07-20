package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// httpRequestTotal là metric đếm tổng số request HTTP đã nhận được, phân loại theo method, path và status code
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Tổng số HTTP request theo method, path, status",
		},
		[]string{"method", "path", "status"},
    )

	// httpRequestDuration là metric đo thời gian xử lý request HTTP, phân loại theo method và path
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Thời gian xử lý request",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method","path"},
	)
)

func init(){
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

// MetricsMiddleware là middleware để thu thập các metric về request HTTP
func MetricsMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		start := time.Now() // Bắt đầu đo thời gian

		c.Next() // Tiếp tục xử lý request

		duration := time.Since(start).Seconds() // Tính toán thời gian xử lý request
		path := c.FullPath() // Lấy path của request
		if path == "" {
			path = "unknown"
		}

		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			path,
			strconv.Itoa(c.Writer.Status()),
		).Inc() // Tăng counter cho request này

		httpRequestDuration.WithLabelValues(
			c.Request.Method, path,
		).Observe(duration) // Ghi nhận thời gian xử lý request
	}
}