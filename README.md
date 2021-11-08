# OpenTelemetry Polyglot Architecture Sample

## Getting Started

1. Run the services:
    ```shell
    $ cp .env.template .env
    $ docker-compose up -d
    ```

2. Run seeds:
    ```shell
    $ cd node-legacy-backend
    $ cp .env.dist .env
    $ npm install
    $ npm run data:import
    ```
### Services

- [Amaxson Frontend in React.js](http://localhost:3000)
- [BFF in Node.js](http://localhost:8000)
- [Product Service in Golang](http://localhost:8090/health)
- [Order Service in Kotlin](http://localhost:8080/actuator/health)
- [Legacy Back-end in Node.js](http://localhost:5000)
- [Jaeger Query UI](http://localhost:16686)
- [Zipkin](http://localhost:9411)
- [Prometheus](http://localhost:9090)
- [Kibaba](http://localhost:5601)
- [Mongo Express](http://localhost:8081)
