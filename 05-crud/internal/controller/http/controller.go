package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/zeuge/hw-go/05-crud/config"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/server"
)

type Controller struct {
	server *http.Server
	uc     *usecase.UserUseCase
	mc     *metricCollector
}

func New(cfg *config.HTTPServerConfig, uc *usecase.UserUseCase) *Controller {
	mc := newMetricCollector()

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		Handler:           mux,
	}

	controller := &Controller{
		server: server,
		uc:     uc,
		mc:     mc,
	}

	if cfg.UsePprof {
		mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	}

	usersMux := http.NewServeMux()
	usersMux.HandleFunc(http.MethodGet+" /users", controller.getUsersHandler)
	usersMux.HandleFunc(http.MethodPost+" /users", controller.createUserHandler)
	usersMux.HandleFunc(http.MethodGet+" /users/{id}", controller.getUserHandler)
	usersMux.HandleFunc(http.MethodDelete+" /users/{id}", controller.deleteUserHandler)

	usersHandler := controller.metricsMiddleware(usersMux)

	mux.Handle("/users", usersHandler)
	mux.Handle("/users/", usersHandler)
	mux.HandleFunc(http.MethodGet+" /live", controller.liveHandler)
	mux.HandleFunc(http.MethodGet+" /ready", controller.readyHandler)

	return controller
}

func (c *Controller) Start() error {
	slog.Info("🚀 Server running at http://" + c.server.Addr)

	go http.ListenAndServe(":9091", promhttp.Handler()) //nolint:errcheck

	err := c.server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("c.server.ListenAndServe: %w", err)
	}

	return nil
}

func (c *Controller) Stop(ctx context.Context) error {
	err := c.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("c.server.Shutdown: %w", err)
	}

	slog.Info("📴 Server stopped gracefully")

	return nil
}
