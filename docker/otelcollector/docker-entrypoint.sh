#!/usr/bin/env sh

# default values
RECEIVERS_CORS_ALLOWED_ORIGINS="${RECEIVERS_CORS_ALLOWED_ORIGINS:-http://*, https://*}" \
TRACES_RECEIVERS_ENABLED="${TRACES_RECEIVERS_ENABLED:-otlp}" \
METRICS_RECEIVERS_ENABLED="${METRICS_RECEIVERS_ENABLED:-otlp}" \
LOGS_RECEIVERS_ENABLED="${LOGS_RECEIVERS_ENABLED:-otlp}" \
TRACES_EXPORTERS_ENABLED="${TRACES_EXPORTERS_ENABLED:-zipkin,jaeger,otlp/elastic}" \
METRICS_EXPORTERS_ENABLED="${METRICS_EXPORTERS_ENABLED:-prometheus,otlp/elastic}" \
LOGS_EXPORTERS_ENABLED="${LOGS_EXPORTERS_ENABLED:-logging}" \
  envsubst < /app/otel-collector-config.yaml.template > /app/otel-collector-config.yaml;

/app/otelcol $@
