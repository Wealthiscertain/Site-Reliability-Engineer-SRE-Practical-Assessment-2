const http = require('http');
const { NodeTracerProvider } = require('@opentelemetry/sdk-trace-node');
const { registerInstrumentations } = require('@opentelemetry/instrumentation');
const { HttpInstrumentation } = require('@opentelemetry/instrumentation-http');
const { expressInstrumentation } = require('@opentelemetry/instrumentation-express');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');
const { OTLPTraceExporter } = require('@opentelemetry/exporter-otlp-grpc');
const { SimpleSpanProcessor } = require('@opentelemetry/sdk-trace-base');

// init tracer
const provider = new NodeTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'paymentservice',
    [SemanticResourceAttributes.SERVICE_VERSION]: '1.0.0',
    [SemanticResourceAttributes.DEPLOYMENT_ENVIRONMENT]: 'production',
  }),
});

const exporter = new OTLPTraceExporter({
  url: process.env.OTEL_EXPORTER_OTLP_ENDPOINT || 'http://otel-agent.observability.svc.cluster.local:4317',
});
provider.addSpanProcessor(new SimpleSpanProcessor(exporter));
provider.register();

registerInstrumentations({
  instrumentations: [
    new HttpInstrumentation(),
    expressInstrumentation(),
  ],
});

const tracer = provider.getTracer('paymentservice');

const requestCounter = provider.getMeter('paymentservice').createCounter('payments.attempts');

const server = http.createServer((req, res) => {
  const span = tracer.startSpan('validate-payment');
  span.setAttributes({ 'payment.method': 'card', 'order.total': 42.50 });
  // simulate an error
  if (Math.random() < 0.1) {
    span.setStatus({ code: 2, message: 'payment failed' });
    res.statusCode = 500;
    res.end('failure');
  } else {
    res.end('ok');
  }
  span.end();
  requestCounter.add(1);
});

server.listen(3000, () => {
  console.log('paymentservice listening on 3000');
});
