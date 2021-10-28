package telemetry

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func Initialize(logger *logrus.Logger) func() {
	ctx := context.Background()
	// shutdown := enableConsoleProvider(ctx, os.Stdout, logger)
	// shutdown := enableJaegerTracerProvider(ctx, logger)
	shutdown := enableCollectorProvider(ctx, logger)

	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	return func() {
		if err := shutdown(); err != nil {
			logger.Fatal(errors.Wrap(err, "failed to shutdown telemetry components"))
		}
	}
}

func createResource(ctx context.Context, logger *logrus.Logger) *resource.Resource {
	r, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(viper.GetString("APP_NAME")),
			semconv.ServiceVersionKey.String(viper.GetString("APP_VERSION")),
			semconv.ServiceInstanceIDKey.String(uuid.Must(uuid.NewRandom()).String()),
		),
	)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to create resource"))
	}

	r, err = resource.Merge(resource.Default(), r)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to create resource"))
	}
	return r
}
