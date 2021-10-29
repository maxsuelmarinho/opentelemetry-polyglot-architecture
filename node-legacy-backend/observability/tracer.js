'use strict';

import * as opentelemetry from '@opentelemetry/api';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';
import { registerInstrumentations } from '@opentelemetry/instrumentation';
import { HttpInstrumentation } from '@opentelemetry/instrumentation-http';
import { ExpressInstrumentation, ExpressLayerType } from '@opentelemetry/instrumentation-express';
import { MongoDBInstrumentation } from '@opentelemetry/instrumentation-mongodb';
import { DnsInstrumentation } from '@opentelemetry/instrumentation-dns';
import { NodeTracerProvider } from '@opentelemetry/sdk-trace-node';
import { SimpleSpanProcessor, ConsoleSpanExporter } from '@opentelemetry/sdk-trace-base';
import { JaegerExporter } from '@opentelemetry/exporter-jaeger';
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin';
import { OTLPTraceExporter } from '@opentelemetry/exporter-otlp-http';
import dotenv from 'dotenv';
import { DEFAULT_SERVICE_NAME } from './constants.js';

dotenv.config();
const exporterType = process.env.TRACING_EXPORTER || 'console';
const serviceName = process.env.SERVICE_NAME || DEFAULT_SERVICE_NAME;
const provider = new NodeTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
  }),
});

if (process.env.COLLECTOR_DIAGNOSTIC_ENABLED) {
  // Optional and only needed to see the internal diagnostic logging
  opentelemetry.diag.setLogger(new opentelemetry.DiagConsoleLogger(), opentelemetry.DiagLogLevel.DEBUG);
}

registerInstrumentations({
  tracerProvider: provider,
  instrumentations: [
    new HttpInstrumentation({
        requestHook: (span, request) => {
          span.updateName(`${request.method} ${request.url}`);
        },
    }),
    new ExpressInstrumentation({
      ignoreLayersType: [
        // ExpressLayerType.MIDDLEWARE,
        // ExpressLayerType.ROUTER,
        // ExpressLayerType.REQUEST_HANDLER,
      ],
    }), // https://github.com/open-telemetry/opentelemetry-js-contrib/tree/main/plugins/node/opentelemetry-instrumentation-express
    new MongoDBInstrumentation({
      enhancedDatabaseReporting: true,
      // responseHook: (span, data) => {
      // },
    }), // https://github.com/open-telemetry/opentelemetry-js-contrib/tree/main/plugins/node/opentelemetry-instrumentation-mongodb
    new DnsInstrumentation({
      // Avoid dns lookup loop with http zipkin calls
      ignoreHostnames: ['localhost'],
    }), // https://github.com/open-telemetry/opentelemetry-js-contrib/tree/main/plugins/node/opentelemetry-instrumentation-dns
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
    exporter = new OTLPTraceExporter({
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
  const spanContext = opentelemetry.trace.getSpanContext(opentelemetry.context.active());
  req.traceId = spanContext && spanContext.traceId;
  next();
};

console.log(`tracing initialized for ${serviceName} sending span to ${exporterType}`);
