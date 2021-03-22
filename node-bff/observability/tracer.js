'use strict';

import * as opentelemetry from '@opentelemetry/api';
import { registerInstrumentations } from '@opentelemetry/instrumentation';
import { HttpInstrumentation } from '@opentelemetry/instrumentation-http';
import { ExpressInstrumentation } from '@aspecto/opentelemetry-instrumentation-express'
import { NodeTracerProvider } from '@opentelemetry/node';
import { SimpleSpanProcessor, ConsoleSpanExporter } from '@opentelemetry/tracing';
import { JaegerExporter } from '@opentelemetry/exporter-jaeger';
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin';
import { CollectorTraceExporter } from '@opentelemetry/exporter-collector';
import dotenv from 'dotenv';
import { DEFAULT_SERVICE_NAME } from './constants.js';

dotenv.config();
const exporterType = process.env.TRACING_EXPORTER || 'console';
const serviceName = process.env.SERVICE_NAME || DEFAULT_SERVICE_NAME;
const provider = new NodeTracerProvider();

if (process.env.COLLECTOR_DIAGNOSTIC_ENABLED) {
  // Optional and only needed to see the internal diagnostic logging
  opentelemetry.diag.setLogger(new opentelemetry.DiagConsoleLogger(), opentelemetry.DiagLogLevel.DEBUG);
}

registerInstrumentations({
  tracerProvider: provider,
  instrumentations: [
    new HttpInstrumentation(),
    new ExpressInstrumentation(),
    {
      plugins: {
        express: {
          enabled: true,
          path: '@opentelemetry/plugin-express',
        },
        http: {
          enabled: true,
          path: '@opentelemetry/plugin-http',
        },
      },
    },
  ],
});

const createTraceExporter = (serviceName, exporterType) => {
  let exporter;
  if (exporterType.toLowerCase() === "jaeger") {
    exporter = new JaegerExporter({
      serviceName
    });
  } else if (exporterType.toLowerCase() === "zipkin") {
    exporter = new ZipkinExporter({
      serviceName
    });
  } else if (exporterType.toLowerCase() === "collector") {
    exporter = new CollectorTraceExporter({
      url: process.env.COLLECTOR_TRACE_URL || 'http://localhost:55681/v1/trace',
      serviceName,
    });
  } else {
    exporter = new ConsoleSpanExporter();
  }

  return exporter;
};

const exporter = createTraceExporter(serviceName, exporterType);

provider.addSpanProcessor(new SimpleSpanProcessor(exporter));
provider.register();

//opentelemetry.trace.setGlobalTracerProvider(provider);

export const tracer = opentelemetry.trace.getTracer(serviceName);

export const addTraceId = (req, res, next) => {
  const spanContext = opentelemetry.getSpanContext(opentelemetry.context.active());
  req.traceId = spanContext && spanContext.traceId;
  next();
};

console.log(`tracing initialized for ${serviceName} sending span to ${exporterType}`);
