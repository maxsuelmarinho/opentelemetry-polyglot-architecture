FROM node:12-alpine as builder
ARG environment_profile=docker
RUN apk update && apk add --no-cache git ca-certificates tzdata openssh make
WORKDIR /app
COPY . .
RUN cp .env.${environment_profile} .env
RUN npm install
RUN CI=false npm run build

FROM nginx:alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/build /usr/share/nginx/html
COPY --from=builder /app/docker/nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080
