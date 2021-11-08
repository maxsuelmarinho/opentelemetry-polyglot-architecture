FROM openjdk:8-jdk-alpine as builder
RUN apk update && \
    apk add --no-cache
WORKDIR /app
COPY *.gradle.kts gradlew /app/
COPY ./gradle /app/gradle
RUN chmod +x gradlew && ./gradlew --version
RUN ./gradlew --stacktrace --info clean
COPY . .
RUN ./gradlew --stacktrace --info build

FROM openjdk:8-alpine
ENV LANG=en_US.utf8 \
    APP_HOME=/app \
    JAVA_GC_ARGS="-XX:MinHeapFreeRatio=20 -XX:MaxHeapFreeRatio=40 -XX:+UseSerialGC -XX:GCTimeRatio=4 -XX:AdaptiveSizePolicyWeight=90" \
    MIN_HEAP_SIZE="40M" \
    MAX_HEAP_SIZE="512M" \
    THREADSTACK_SIZE="228k"
RUN apk update && \
    apk add --no-cache bash curl
WORKDIR /app
COPY --from=builder /app/build/libs/order-service*.jar ./server.jar
COPY --from=builder /app/agent/opentelemetry-javaagent-all-*.jar ./agent/opentelemetry-javaagent-all.jar
COPY --from=builder /app/agent/elastic-apm-agent-*.jar ./agent/elastic-apm-agent.jar
COPY --from=builder /app/scripts/docker-entrypoint.sh ./
RUN sh -c 'touch /app/server.jar'
RUN addgroup -S appgroup
RUN adduser -S -D -h /app appuser appgroup
RUN chown -R appuser:appgroup /app
EXPOSE 8080
CMD ["sh", "/app/docker-entrypoint.sh"]
