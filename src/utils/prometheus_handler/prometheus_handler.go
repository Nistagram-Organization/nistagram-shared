package prometheus_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
)

func GetHttpRequestsCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_count",
			Help: "Total number of http requests",
		},
		[]string{"method", "originalUrl", "statusCode"},
	)
}

func GetHttpRequestsSize() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_size",
			Help: "Total traffic in megabytes",
		},
		[]string{},
	)
}

func PrometheusMiddleware(requestsCounter *prometheus.CounterVec, requestsSize *prometheus.CounterVec) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestSize := int(c.Request.ContentLength)

		c.Next()

		responseSize := c.Writer.Size()
		requestsCounter.WithLabelValues(c.Request.Method, c.Request.RequestURI, strconv.Itoa(c.Writer.Status()))
		requestsSize.WithLabelValues(strconv.Itoa((requestSize + responseSize) / 1024 / 1024)).Inc()
	}
}

func PrometheusGinHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
