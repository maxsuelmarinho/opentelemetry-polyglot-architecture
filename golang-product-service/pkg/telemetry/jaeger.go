package telemetry

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func enableJaegerTracerProvider(ctx context.Context, logger *logrus.Logger) func() error {
	resource := createResource(ctx, logger)

	// flush, err := jaeger.InstallNewPipeline(
	// 	jaeger.WithCollectorEndpoint(viper.GetString("JAEGER_EXPORTER_ENDPOINT")),
	// 	jaeger.WithProcess(jaeger.Process{
	// 		ServiceName: viper.GetString("APP_NAME"),
	// 		// Tags: []attribute.KeyValue{
	// 		// 	attribute.String("exporter", "jaeger"),
	// 		// 	attribute.Float64("float", 312.23),
	// 		// },
	// 	}),
	// 	// jaeger.WithSDK(&sdktrace.Config{
	// 	// 	DefaultSampler: sdktrace.AlwaysSample(),
	// 	// 	Resource:       resource,
	// 	// }),
	// )
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(viper.GetString("JAEGER_EXPORTER_ENDPOINT"))))
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize jaeger exporter"))
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource),
	)

	otel.SetTracerProvider(tp)

	return func() error {
		if err := tp.Shutdown(ctx); err != nil {
			return err
		}

		return nil
	}
}
