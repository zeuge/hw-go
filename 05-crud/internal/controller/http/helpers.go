package http

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func writeJSON(ctx context.Context, w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			slog.ErrorContext(ctx, "json.Encode", "error", err)
		}
	}
}

func writeErrorJSON(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	resp := map[string]string{"error": msg}

	writeJSON(ctx, w, status, resp)
}
