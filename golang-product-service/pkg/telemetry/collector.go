package telemetry

import (
	"context"
	"crypto/x509"
	"fmt"
	"strings"
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
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func enableCollectorProvider(ctx context.Context, logger *logrus.Logger) func() error {
	connectionConfig, err := getConnectionConfig()
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to create system cert pool %w", err))
	}

	traceExporter, err := otlptracegrpc.New(ctx,
		connectionConfig,
		otlptracegrpc.WithCompressor(gzip.Name),
		otlptracegrpc.WithHeaders(getHeaders()),
		otlptracegrpc.WithEndpoint(viper.GetString("COLLECTOR_EXPORTER_ENDPOINT")),
		otlptracegrpc.WithTimeout(10*time.Second),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
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

func getConnectionConfig() (otlptracegrpc.Option, error) {
	if viper.GetBool("COLLECTOR_EXPORTER_INSECURE_CONNECTION") {
		return otlptracegrpc.WithInsecure(), nil
	}

	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to create system cert pool %w", err)
	}

	return otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(pool, "")), nil
}

func getHeaders() map[string]string {
	headers := map[string]string{}
	for _, header := range strings.Split(viper.GetString("COLLECTOR_EXPORTER_HEADERS"), ",") {
		kv := strings.Split(header, "=")
		if len(kv) != 2 {
			continue
		}

		headers[kv[0]] = kv[1]
	}
	return headers
}
