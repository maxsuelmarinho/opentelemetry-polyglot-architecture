# https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/cmd/otelcontribcol/Dockerfile
FROM otel/opentelemetry-collector-contrib:0.22.0 as service

FROM alpine:3.13
RUN apk add --no-cache gettext
WORKDIR /app
COPY --from=service /otelcontribcol ./otelcol
COPY ./docker-entrypoint.sh ./
COPY ./otel-collector-config.yaml ./otel-collector-config.yaml.template
RUN addgroup -S appgroup
RUN adduser -S -D -h /app appuser appgroup
RUN chown -R appuser:appgroup /app
USER appuser
ENTRYPOINT [ "/app/docker-entrypoint.sh" ]
CMD [ "--config=/app/otel-collector-config.yaml", "${OTELCOL_ARGS}" ]
EXPOSE 4317 55678 55679
