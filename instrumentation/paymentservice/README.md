# Paymentservice (Node.js) instrumentation

The payment service is implemented in Node.js. Use the OpenTelemetry Node.js SDK with auto-instrumentation and manual spans.

### Steps
1. Install `@opentelemetry/sdk-node`, `@opentelemetry/instrumentation-http`, etc.
2. Configure resource attributes and enable auto-instrumentation for HTTP.
3. Add manual spans around payment validation and record errors with `span.setStatus`.
4. Create a custom counter metric for payment attempts.
5. Export via OTLP to the agent's hostIP.

Example snippet in `app.js`:

```js
// TODO: create file instrumentation/paymentservice/app.js with instrumentation setup
```
