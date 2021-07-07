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

func GetHttpRequestsSize() prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_size",
			Help: "Total traffic in megabytes",
		},
	)
}

func PrometheusMiddleware(requestsCounter *prometheus.CounterVec, requestsSize prometheus.Counter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestSize := int(c.Request.ContentLength)

		c.Next()

		responseSize := c.Writer.Size()
		requestsCounter.WithLabelValues(c.Request.Method, c.Request.RequestURI, strconv.Itoa(c.Writer.Status()))
		requestsSize.Add(float64((requestSize + responseSize) / 1024 / 1024))
	}
}

func PrometheusGinHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
