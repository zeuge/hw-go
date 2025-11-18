package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
	"github.com/zeuge/hw-go/05-crud/internal/tracing"
)

const tracerName string = "http_client"

func (c *Controller) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "getUsersHandler")
	defer span.End()

	users, err := c.uc.GetUsers(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "c.uc.GetUsers", "error", err)
		writeErrorJSON(ctx, w, http.StatusInternalServerError, "internal error")

		return
	}

	writeJSON(ctx, w, http.StatusOK, users)
}

func (c *Controller) createUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "createUserHandler")
	defer span.End()

	var input dto.CreateUser

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.ErrorContext(ctx, "json.NewDecoder", "error", err)
		writeErrorJSON(ctx, w, http.StatusBadRequest, err.Error())

		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Role)
	if err != nil {
		slog.ErrorContext(ctx, "entity.NewUser", "error", err)
		writeErrorJSON(ctx, w, http.StatusBadRequest, err.Error())

		return
	}

	err = c.uc.CreateUser(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "c.uc.CreateUser", "error", err)
		writeErrorJSON(ctx, w, http.StatusInternalServerError, "internal error")

		return
	}

	writeJSON(ctx, w, http.StatusCreated, &user)
}

func (c *Controller) getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "getUserHandler")
	defer span.End()

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.ErrorContext(ctx, "uuid.Parse", "error", err)
		writeErrorJSON(ctx, w, http.StatusBadRequest, err.Error())

		return
	}

	user, err := c.uc.GetUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "c.uc.GetUser", "error", err)

		if errors.Is(err, entity.ErrNotFound) {
			writeErrorJSON(ctx, w, http.StatusNotFound, "user not found")

			return
		}

		writeErrorJSON(ctx, w, http.StatusInternalServerError, "internal error")

		return
	}

	writeJSON(ctx, w, http.StatusOK, user)
}

func (c *Controller) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "deleteUserHandler")
	defer span.End()

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.ErrorContext(ctx, "uuid.Parse", "error", err)
		writeErrorJSON(ctx, w, http.StatusBadRequest, err.Error())

		return
	}

	err = c.uc.DeleteUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "c.uc.DeleteUser", "error", err)
		writeErrorJSON(ctx, w, http.StatusInternalServerError, "internal error")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
