import { DiagConsoleLogger, DiagLogLevel, diag } from '@opentelemetry/api';
import { CollectorMetricExporter } from '@opentelemetry/exporter-collector';
import { MetricExporter } from '@opentelemetry/metrics';
import { DEFAULT_SERVICE_NAME } from './constants.js';

if (process.env.REACT_APP_ENVIRONMENT === 'development') {
  diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG);
}

const serviceName = process.env.REACT_APP_SERVICE_NAME || DEFAULT_SERVICE_NAME;
const metricExporter = new CollectorMetricExporter({
  serviceName: serviceName,
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

