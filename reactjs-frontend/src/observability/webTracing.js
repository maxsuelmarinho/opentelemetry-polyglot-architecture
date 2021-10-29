import { context, trace, DiagConsoleLogger, DiagLogLevel, diag } from '@opentelemetry/api';
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base';
import { WebTracerProvider } from '@opentelemetry/sdk-trace-web';
import { Resource }from '@opentelemetry/resources';
import { SemanticResourceAttributes }from '@opentelemetry/semantic-conventions';
import { XMLHttpRequestInstrumentation } from '@opentelemetry/instrumentation-xml-http-request';
import { DocumentLoadInstrumentation } from '@opentelemetry/instrumentation-document-load';
import { UserInteractionInstrumentation } from '@opentelemetry/instrumentation-user-interaction';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import { OTLPTraceExporter } from '@opentelemetry/exporter-otlp-http';
// import { JaegerExporter } from '@opentelemetry/exporter-jaeger';
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin';
import { CompositePropagator, W3CTraceContextPropagator } from '@opentelemetry/core';
import { B3Propagator } from '@opentelemetry/propagator-b3';
import { registerInstrumentations } from '@opentelemetry/instrumentation';
import { FetchInstrumentation } from '@opentelemetry/instrumentation-fetch';
import { BaseOpenTelemetryComponent } from '@opentelemetry/plugin-react-load';
import { DEFAULT_SERVICE_NAME } from './constants';

export default () => {
  if (process.env.REACT_APP_COLLECTOR_DIAGNOSTIC_ENABLED) {
    // Optional and only needed to see the internal diagnostic logging
    diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG);
  }

  const serviceName = process.env.REACT_APP_SERVICE_NAME || DEFAULT_SERVICE_NAME;
  const exporterType = process.env.REACT_APP_TRACING_EXPORTER || 'console';
  console.log(serviceName);

  const provider = new WebTracerProvider({
    resource: new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
    }),
  });

  registerInstrumentations({
    tracerProvider: provider,
    instrumentations: [
      new UserInteractionInstrumentation(),
      new DocumentLoadInstrumentation(),
      new FetchInstrumentation({
        ignoreUrls: [
          /localhost:3000/,
          /.*\.hot-.*\.json$/, // ignores hot reload files
        ],
        propagateTraceHeaderCorsUrls: [
          /localhost:8000/,
        ],
        clearTimingResources: true,
      }),
      new XMLHttpRequestInstrumentation({
        ignoreUrls: [
          /localhost:3000/,
          /.*\.hot-.*\.json$/, // ignores hot reload files
        ],
        propagateTraceHeaderCorsUrls: [
          /localhost:8000/,
        ],
      }),
    ],
  });

  const createTraceExporter = (serviceName, exporterType) => {
    let exporter;
    if (exporterType.toLowerCase() === "jaeger") {
      // exporter = new JaegerExporter({
      //   serviceName
      // });
    } else if (exporterType.toLowerCase() === "zipkin") {
      exporter = new ZipkinExporter({
        serviceName
      });
    } else if (exporterType.toLowerCase() === "collector") {
      exporter = new OTLPTraceExporter({
        url: process.env.REACT_APP_COLLECTOR_TRACE_URL || 'http://localhost:55681/v1/trace',
        headers: { // https://github.com/open-telemetry/opentelemetry-js/issues/2321#issuecomment-889861080
          "Content-Type": "application/json"
        },
        serviceName: serviceName,
      });
    } else {
      exporter = new ConsoleSpanExporter();
    }

    return exporter;
  };

  const exporter = createTraceExporter(serviceName, exporterType);

  //provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
  provider.addSpanProcessor(new SimpleSpanProcessor(exporter));

  provider.register({
    contextManager: new ZoneContextManager(),
    // propagator: new B3Propagator(),
    propagator: new CompositePropagator({
      propagators: [
        new B3Propagator(),
        new W3CTraceContextPropagator(),
      ],
    }),
  });

  const tracer = provider.getTracer(serviceName);
  BaseOpenTelemetryComponent.setTracer(serviceName)
  BaseOpenTelemetryComponent.setLogger(provider.logger)

  console.log(`tracing initialized for ${serviceName} sending span to collector`);
  return tracer;
};

