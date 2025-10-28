package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/zeuge/hw-go/05-crud/config"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/server"
)

type Controller struct {
	server *http.Server
	uc     *usecase.UserUseCase
}

func New(cfg *config.HTTPServerConfig, uc *usecase.UserUseCase) *Controller {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		Handler:           mux,
	}

	controller := &Controller{
		server: server,
		uc:     uc,
	}

	mux.HandleFunc(http.MethodGet+" /live", controller.liveHandler)
	mux.HandleFunc(http.MethodGet+" /ready", controller.readyHandler)

	mux.HandleFunc(http.MethodGet+" /users", controller.getUsersHandler)
	mux.HandleFunc(http.MethodPost+" /users", controller.createUserHandler)
	mux.HandleFunc(http.MethodGet+" /users/{id}", controller.getUserHandler)
	mux.HandleFunc(http.MethodDelete+" /users/{id}", controller.deleteUserHandler)

	return controller
}

func (c *Controller) Start() error {
	slog.Info("🚀 Server running at http://" + c.server.Addr)

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
