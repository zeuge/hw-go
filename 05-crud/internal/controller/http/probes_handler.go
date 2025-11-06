package http

import "net/http"

func (c *Controller) liveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) readyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := c.uc.DbCheck(ctx)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	w.WriteHeader(http.StatusOK)
}
