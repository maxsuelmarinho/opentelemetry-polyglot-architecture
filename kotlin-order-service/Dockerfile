FROM openjdk:8-jdk-alpine as builder
RUN apk update && \
    apk add --no-cache
WORKDIR /app
COPY . .
RUN chmod +x gradlew && \
    ./gradlew --stacktrace --info clean build

FROM openjdk:8-alpine
ENV LANG=en_US.utf8 \
    APP_HOME=/app \
    JAVA_GC_ARGS="-XX:MinHeapFreeRatio=20 -XX:MaxHeapFreeRatio=40 -XX:+UseSerialGC -XX:GCTimeRatio=4 -XX:AdaptiveSizePolicyWeight=90" \
    MIN_HEAP_SIZE="40M" MAX_HEAP_SIZE="512M" THREADSTACK_SIZE="228k" \
    JAVA_OPTS="${JAVA_OPTS} \
      -server \
      -Djava.security.egd=file:/dev/./urandom \
      -XX:CompressedClassSpaceSize=64m \
      -XX:MaxMetaspaceSize=256m \
      -javaagent:/app/agent/opentelemetry-javaagent-all.jar"
RUN apk update && \
    apk add --no-cache bash curl
WORKDIR /app
COPY --from=builder /app/build/libs/order-service*.jar ./server.jar
COPY --from=builder /app/agent/opentelemetry-javaagent-all-*.jar ./agent/opentelemetry-javaagent-all.jar
COPY --from=builder /app/scripts/docker-entrypoint.sh ./
RUN sh -c 'touch /app/server.jar'
#USER appuser
EXPOSE 8080
#ENTRYPOINT ["java","${JAVA_OPTS}","-Xms${MIN_HEAP_SIZE}","-Xmx${MAX_HEAP_SIZE}","-Xss${THREADSTACK_SIZE}","${JAVA_GC_ARGS}","-jar","${APP_HOME}/server.jar"]
CMD ["sh", "/app/docker-entrypoint.sh"]
