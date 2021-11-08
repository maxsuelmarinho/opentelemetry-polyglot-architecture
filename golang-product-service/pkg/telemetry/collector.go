package telemetry

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/host"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric/global"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

func enableCollectorProvider(ctx context.Context, logger *logrus.Logger) func() error {
	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(viper.GetString("COLLECTOR_EXPORTER_ENDPOINT")),
		otlptracegrpc.WithDialOption(grpc.WithBlock()), // useful for testing
	)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to create exporter"))
	}

	resource := createResource(ctx, logger)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(viper.GetFloat64("TRACE_SAMPLING_RATIO"))),
		sdktrace.WithResource(resource),
	)

	client := otlpmetricgrpc.NewClient(
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(viper.GetString("COLLECTOR_EXPORTER_ENDPOINT")),
		otlpmetricgrpc.WithDialOption(grpc.WithBlock()), // useful for testing
	)
	metricExporter, err := otlpmetric.New(ctx, client)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to create metric exporter"))
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithExactDistribution(),
			metricExporter,
		),
		controller.WithExporter(metricExporter),
		controller.WithCollectPeriod(viper.GetDuration("COLLECTOR_COLLECT_PERIOD")),
	)

	otel.SetTracerProvider(tracerProvider)
	global.SetMeterProvider(pusher)

	if err := pusher.Start(ctx); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to start metric controller"))
	}

	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to start publishing runtime metric"))
	}

	if err := host.Start(); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to start publishing host metric"))
	}

	return func() error {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			return errors.Wrap(err, "failed to shutdown trace provider")
		}

		if err := pusher.Stop(ctx); err != nil {
			return errors.Wrap(err, "failed to stop metric controller")
		}

		return nil
	}

}
