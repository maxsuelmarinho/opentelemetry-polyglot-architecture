package telemetry

import (
	"context"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric/global"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func enableConsoleProvider(ctx context.Context, w io.Writer, logger *logrus.Logger) func() error {
	exporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(w),
	)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize stdout exporter"))
	}

	resource := createResource(ctx, logger)

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(tp)

	metricExporter, err := stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(),
		stdoutmetric.WithWriter(w),
	)
	if err != nil {
		logger.Fatalf("creating stdoutmetric exporter: %v", err)
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithExactDistribution(),
			metricExporter,
		),
		controller.WithExporter(metricExporter),
		controller.WithCollectPeriod(5*time.Second),
	)

	if err := pusher.Start(ctx); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize metric controller"))
	}

	global.SetMeterProvider(pusher)

	return func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			otel.Handle(err)
			return err
		}

		if err := pusher.Stop(ctx); err != nil {
			otel.Handle(err)
			return errors.Wrap(err, "failed to stop metric pusher controller")
		}

		return nil
	}
}
