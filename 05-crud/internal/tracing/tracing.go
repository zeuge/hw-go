package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/zeuge/hw-go/05-crud/config"
)

type tracer struct {
	Provider *sdktrace.TracerProvider
}

func New(ctx context.Context, cfg *config.TracingConfig) (*tracer, error) {
	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(cfg.Endpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("otlptracehttp.New: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.AppName),
		)),
	)

	otel.SetTracerProvider(tp)

	return &tracer{Provider: tp}, nil
}

func (t *tracer) Shutdown(ctx context.Context) error {
	err := t.Provider.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("t.Provider.Shutdown: %w", err)
	}

	return nil
}

func GetTracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
