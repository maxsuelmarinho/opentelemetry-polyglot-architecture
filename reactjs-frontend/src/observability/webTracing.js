import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/tracing';
import { WebTracerProvider } from '@opentelemetry/web';
import { DiagConsoleLogger, DiagLogLevel, diag } from '@opentelemetry/api';
import { registerInstrumentations } from '@opentelemetry/instrumentation';
import { DocumentLoad } from '@opentelemetry/plugin-document-load';
import { CollectorTraceExporter } from '@opentelemetry/exporter-collector';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import { BaseOpenTelemetryComponent } from '@opentelemetry/plugin-react-load';
import { FetchInstrumentation } from '@opentelemetry/instrumentation-fetch';
import { XMLHttpRequestInstrumentation } from '@opentelemetry/instrumentation-xml-http-request';
import { B3Propagator } from '@opentelemetry/propagator-b3';
import { DEFAULT_SERVICE_NAME } from './constants';

export default () => {

  if (process.env.REACT_APP_COLLECTOR_DIAGNOSTIC_ENABLED) {
    // Optional and only needed to see the internal diagnostic logging
    diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG);
  }

  const serviceName = process.env.REACT_APP_SERVICE_NAME || DEFAULT_SERVICE_NAME;
  console.log(serviceName);

  const provider = new WebTracerProvider();

  registerInstrumentations({
    tracerProvider: provider,
    instrumentations: [
      new DocumentLoad(),
      new FetchInstrumentation(),
      new XMLHttpRequestInstrumentation({
        propagateTraceHeaderCorsUrls: [
          /localhost:8000/,
        ],
      }),
    ],
  });

  const exporter = new CollectorTraceExporter({
    url: process.env.REACT_APP_COLLECTOR_TRACE_URL || 'http://localhost:55681/v1/trace',
    serviceName: serviceName,
  });

  //provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
  provider.addSpanProcessor(new SimpleSpanProcessor(exporter));

  provider.register({
    contextManager: new ZoneContextManager(),
    //propagator: new B3Propagator(),
  });


  const tracer = provider.getTracer(serviceName);
  BaseOpenTelemetryComponent.setTracer(serviceName);
  BaseOpenTelemetryComponent.setLogger(provider.logger);

  console.log(`tracing initialized for ${serviceName} sending span to collector`);
  return tracer;
};

