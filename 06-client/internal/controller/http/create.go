package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/zeuge/hw-go/06-client/internal/entity"
	"github.com/zeuge/hw-go/06-client/internal/entity/dto"
)

func (c *Controller) Create(ctx context.Context, cmd entity.Command) error {
	name, _ := cmd.Data["name"].(string)
	email, _ := cmd.Data["email"].(string)
	role, _ := cmd.Data["role"].(string)

	req := dto.CreateUser{
		Name:  name,
		Email: email,
		Role:  role,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	resp, err := c.requestWithContext(ctx, http.MethodPost, "/users", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return entity.ErrUnexpectedHTTPStatus
	}

	var user entity.User

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	err = c.uc.CreateUser(ctx, &user)
	if err != nil {
		return fmt.Errorf("c.uc.CreateUser: %w", err)
	}

	return nil
}
