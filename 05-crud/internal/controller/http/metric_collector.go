package http

import "github.com/prometheus/client_golang/prometheus"

const (
	tagMethod = "method"
)

type metricCollector struct {
	RequestDuration *prometheus.HistogramVec
	RequestsTotal   *prometheus.CounterVec
}

func newMetricCollector() *metricCollector {
	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration",
			Help: "Duration of HTTP requests (milliseconds).",
		},
		[]string{tagMethod},
	)

	requestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{tagMethod},
	)

	prometheus.MustRegister(requestDuration, requestsTotal)

	return &metricCollector{
		RequestDuration: requestDuration,
		RequestsTotal:   requestsTotal,
	}
}
