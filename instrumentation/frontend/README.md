# Frontend service (Go) instrumentation

This service uses Go and will be instrumented using the OpenTelemetry Go SDK.

### Steps
1. Add the `go.opentelemetry.io/otel` packages and any auto-instrumentation middleware (e.g., otelhttp).
2. Configure the SDK with resource attributes (`service.name`, `service.version`, `deployment.environment`).
3. Wrap HTTP handlers with `otelhttp.NewHandler` and create manual spans for template rendering and business logic.
4. Export to OTLP via gRPC to the local DaemonSet agent (use hostIP or service).

Example snippet in `main.go`:

```go
// TODO: create file instrumentation/frontend/main.go with real code
```