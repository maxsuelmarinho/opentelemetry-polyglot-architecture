module github.com/maxsuelmarinho/ecommerce-example/golang-product-service

go 1.14

require (
	github.com/gobuffalo/packr v1.30.1
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/rubenv/sql-migrate v0.0.0-20210215143335-f84234893558
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.1
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.2
	github.com/uptrace/opentelemetry-go-extra/otelsqlx v0.1.2
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.25.0
	go.opentelemetry.io/otel v1.0.1
	go.opentelemetry.io/otel/exporters/jaeger v1.0.1
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.24.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.24.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.1
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.24.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.1
	go.opentelemetry.io/otel/metric v0.24.0
	go.opentelemetry.io/otel/sdk v1.0.1
	go.opentelemetry.io/otel/sdk/metric v0.24.0
	go.opentelemetry.io/otel/trace v1.0.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/grpc v1.41.0
)
