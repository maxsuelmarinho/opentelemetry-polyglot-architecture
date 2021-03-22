package telemetry

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"google.golang.org/grpc"
)

func Initialize(logger *logrus.Logger) func() {

	ctx := context.Background()
	//shutdown := enableConsoleProvider(ctx, logger)
	//shutdown := enableJaegerTracerProvider(ctx, logger)
	shutdown := enableCollectorProvider(ctx, logger)

	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	return func() {
		if err := shutdown(); err != nil {
			logger.Fatal(errors.Wrap(err, "failed to shutdown telemetry components"))
		}
	}
}

func enableConsoleProvider(ctx context.Context, logger *logrus.Logger) func() error {
	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize stdout exporter"))
	}

	resource := createResource(ctx, logger)

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(tp)

	pusher := controller.New(
		processor.New(simple.NewWithExactDistribution(), exporter),
		controller.WithPusher(exporter),
		controller.WithCollectPeriod(5*time.Second),
	)

	if err := pusher.Start(ctx); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize metric controller"))
	}

	global.SetMeterProvider(pusher.MeterProvider())

	return func() error {
		if err := tp.Shutdown(ctx); err != nil {
			return err
		}

		if err := pusher.Stop(ctx); err != nil {
			return errors.Wrap(err, "failed to stop metric pusher controller")
		}

		return nil
	}
}

func enableJaegerTracerProvider(ctx context.Context, logger *logrus.Logger) func() error {
	resource := createResource(ctx, logger)

	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(viper.GetString("JAEGER_EXPORTER_ENDPOINT")),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: viper.GetString("APP_NAME"),
			Tags: []attribute.KeyValue{
				attribute.String("exporter", "jaeger"),
				attribute.Float64("float", 312.23),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
			Resource:       resource,
		}),
	)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize jaeger exporter"))
	}

	return func() error {
		flush()
		return nil
	}
}

func enableCollectorProvider(ctx context.Context, logger *logrus.Logger) func() error {
	driver := otlpgrpc.NewDriver(
		otlpgrpc.WithInsecure(),
		otlpgrpc.WithEndpoint(viper.GetString("COLLECTOR_EXPORTER_ENDPOINT")),
		otlpgrpc.WithDialOption(grpc.WithBlock()), // useful for testing
	)
	exporter, err := otlp.NewExporter(ctx, driver)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to create exporter"))
	}

	resource := createResource(ctx, logger)
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(resource),
	)

	cont := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithPusher(exporter),
		controller.WithCollectPeriod(viper.GetDuration("COLLECTOR_COLLECT_PERIOD")),
	)

	otel.SetTracerProvider(tracerProvider)
	global.SetMeterProvider(cont.MeterProvider())

	if err := cont.Start(ctx); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to start controller"))
	}

	return func() error {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			return errors.Wrap(err, "failed to shutdown trace provider")
		}

		if err := cont.Stop(ctx); err != nil {
			return errors.Wrap(err, "failed to stop controller")
		}

		return nil
	}

}

func createResource(ctx context.Context, logger *logrus.Logger) *resource.Resource {
	resource, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(viper.GetString("APP_NAME")),
			semconv.ServiceVersionKey.String("APP_VERSION"),
			semconv.ServiceInstanceIDKey.String(uuid.Must(uuid.NewRandom()).String()),
		),
	)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to create resource"))
	}

	return resource
}
