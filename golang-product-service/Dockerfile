FROM golang:1.13-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata openssh make
RUN adduser -D -g '' appuser
WORKDIR /app
COPY . .
RUN make -S build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
WORKDIR /app
COPY --from=builder /app/dbconfig.yml ./dbconfig.yml
COPY --from=builder /app/db ./db
COPY --from=builder /app/build/server ./server
COPY --from=builder /app/build/healthchecker /usr/local/bin/healthchecker
USER appuser
HEALTHCHECK --interval=5s --timeout=5s CMD [ "healthchecker" ] || exit 1
ENTRYPOINT ["/app/server"]
