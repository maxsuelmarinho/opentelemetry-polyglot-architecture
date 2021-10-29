import { DiagConsoleLogger, DiagLogLevel, diag } from '@opentelemetry/api';
import { OTLPMetricExporter } from '@opentelemetry/exporter-otlp-http';
import { MeterProvider } from '@opentelemetry/sdk-metrics-base';
import { Resource }from '@opentelemetry/resources';
import { SemanticResourceAttributes }from '@opentelemetry/semantic-conventions';
import { DEFAULT_SERVICE_NAME } from './constants.js';

if (process.env.REACT_APP_COLLECTOR_DIAGNOSTIC_ENABLED) {
  // Optional and only needed to see the internal diagnostic logging (during development)
  diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG);
}

const serviceName = process.env.REACT_APP_SERVICE_NAME || DEFAULT_SERVICE_NAME;

const collectorUrl = process.env.REACT_APP_COLLECTOR_METRIC_URL || 'http://localhost:55681/v1/metrics';
const metricExporter = new OTLPMetricExporter({
  serviceName: serviceName,
  url: collectorUrl,
});

let interval;
let meter;

const stopMetrics = () => {
  console.log("stopping metrics");
  clearInterval(interval);
  meter.shutdown();
};

const startMetrics = () => {
  console.log("starting metrics");
  meter = new MeterProvider({
    exporter: metricExporter,
    interval: process.env.REACT_APP_METRIC_COLLECT_INTERVAL || 1000,
    resource: new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
    }),
  }).getMeter(serviceName);

  const requestCounter = meter.createCounter('requests', {
    description: 'Example of a Counter',
  });

  const upDownCounter = meter.createUpDownCounter('test_up_down_counter', {
    description: 'Example of a UpDownCounter',
  });

  const labels = {
    pid: process.pid,
    environment: process.env.REACT_APP_ENVIRONMENT,
  };

  interval = setInterval(() => {
    requestCounter.bind(labels).add(1);
    upDownCounter.bind(labels).add(Math.random() > 0.5 ? 1 : -1);
  }, 1000);
};

