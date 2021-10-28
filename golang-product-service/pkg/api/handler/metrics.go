package handler

import (
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

const (
	instrumentationVersion = "v0.1.0"
)

var (
	meter = global.GetMeterProvider().Meter(
		viper.GetString("APP_NAME"),
		metric.WithInstrumentationVersion(instrumentationVersion),
	)
	httpRequestsCounter = metric.Must(meter).NewInt64Counter("http_requests_total")
	httpPathKey         = attribute.Key("path")
)
