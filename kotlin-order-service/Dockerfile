FROM openjdk:8-jdk-alpine as builder
RUN apk update && \
    apk add --no-cache
WORKDIR /app
COPY . .
RUN chmod +x gradlew && \
    ./gradlew --stacktrace --info clean build

FROM openjdk:8-alpine
ENV LANG en_US.utf8
RUN apk update && \
    apk add --no-cache bash curl
COPY --from=builder /app/build/libs/order-service*.jar /app/server.jar
RUN sh -c 'touch /app/server.jar'
EXPOSE 8080
ENTRYPOINT ["java","-jar","/app/server.jar"]
