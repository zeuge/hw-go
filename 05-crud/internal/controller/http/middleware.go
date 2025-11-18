package http

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func (c *Controller) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start).Milliseconds()
		method := r.Method

		c.mc.RequestDuration.
			With(prometheus.Labels{
				tagMethod: method,
			}).
			Observe(float64(elapsed))

		c.mc.RequestsTotal.
			With(prometheus.Labels{
				tagMethod: method,
			}).
			Inc()
	})
}
