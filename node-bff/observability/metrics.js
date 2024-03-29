import { DiagConsoleLogger, DiagLogLevel, diag } from '@opentelemetry/api';
import { MeterProvider, ConsoleMetricExporter } from '@opentelemetry/sdk-metrics-base';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';
import { PrometheusExporter } from '@opentelemetry/exporter-prometheus';
import { OTLPMetricExporter } from '@opentelemetry/exporter-otlp-http';
import { DEFAULT_SERVICE_NAME } from './constants.js'
import dotenv from 'dotenv';
dotenv.config();

const serviceName = process.env.SERVICE_NAME || DEFAULT_SERVICE_NAME;
const exporterType = process.env.METRICS_EXPORTER || 'console';

const createMetricExporter = (serviceName, exporterType) => {
  let exporter;
  if (exporterType.toLowerCase() === "prometheus") {
    const prometheusPort = process.env.METRIC_PORT || PrometheusExporter.DEFAULT_OPTIONS.port;
    const prometheusEndpoint = PrometheusExporter.DEFAULT_OPTIONS.endpoint;

    exporter = new PrometheusExporter({
      startServer: true,
      port: prometheusPort,
      endpoint: prometheusEndpoint
    },
    () => {
      console.log(`prometheus scrape endpoint: http://localhost:${prometheusPort}${prometheusEndpoint}`);
    });
  } else if (exporterType.toLowerCase() === "collector") {
    const collectorUrl = process.env.COLLECTOR_METRIC_URL || 'http://localhost:55681/v1/metrics';
    console.log("collector url:", collectorUrl);
    exporter = new OTLPMetricExporter({
      serviceName,
      url: collectorUrl,
    });
  } else {
    exporter = new ConsoleMetricExporter();
  }

  return exporter;
};

if (process.env.COLLECTOR_DIAGNOSTIC_ENABLED) {
  // Optional and only needed to see the internal diagnostic logging
  diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG);
}

const exporter = createMetricExporter(serviceName, exporterType);
const collectInterval = process.env.METRIC_COLLECT_INTERVAL || 3000;
const meter = new MeterProvider({
  exporter,
  interval: collectInterval,
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
  }),
}).getMeter(serviceName);

const requestCount = meter.createCounter('http_requests_total', {
  description: 'Count all incoming requests'
});

// Creating a new labelSet and binding on every request is not ideal as creating the labelSet can often be an expensive operation.
// This is why instruments are created and stored in a Map according to the route key.
const boundInstruments = new Map();

const countAllRequests = () => {
  return (req, res, next) => {
    if (!boundInstruments.has(req.path)) {
      const labels = { route: req.path, environment: process.env.NODE_ENV, service_name: serviceName };
      const boundCounter = requestCount.bind(labels);
      boundInstruments.set(req.path, boundCounter);
    }

    boundInstruments.get(req.path).add(1);
    next();
  };
};

console.log(`metrics initialized for ${serviceName} sending metrics to ${exporterType} every ${collectInterval}ms`);

export { countAllRequests };
