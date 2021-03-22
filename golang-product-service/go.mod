module github.com/maxsuelmarinho/ecommerce-example/golang-product-service

go 1.14

require (
	github.com/gobuffalo/packr v1.30.1
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.1
	github.com/lib/pq v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/rubenv/sql-migrate v0.0.0-20210215143335-f84234893558
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.1
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.18.0
	go.opentelemetry.io/otel v0.18.0
	go.opentelemetry.io/otel/exporters/otlp v0.18.0
	go.opentelemetry.io/otel/exporters/stdout v0.18.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.18.0
	go.opentelemetry.io/otel/metric v0.18.0
	go.opentelemetry.io/otel/sdk v0.18.0
	go.opentelemetry.io/otel/sdk/metric v0.18.0
	go.opentelemetry.io/otel/trace v0.18.0
	google.golang.org/grpc v1.36.0
)
