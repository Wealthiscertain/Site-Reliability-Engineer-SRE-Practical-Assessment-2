# Architectural Decision Log

This document tracks major decisions and reasoning made during the assessment implementation.

- **Collector topology**: Gateway + DaemonSet agent. Agents collect hostmetrics and forward all signals to the gateway for processing and sampling.
- **Exporter protocol**: OTLP/gRPC to Elastic APM Server on port 8200, authenticated via API key stored in a Kubernetes secret.
- **Service selection for instrumentation**: `frontend` (Go), `cartservice` (C# .NET), `paymentservice` (Node.js) to satisfy multi-language requirement.
- **RUM implementation**: Elastic APM RUM agent integrated into the frontend to take advantage of built-in Kibana dashboards and Core Web Vitals support.
- **Infrastructure monitoring**: Fleet-managed Elastic Agent preferred for ease of policy management; fallback to Metricbeat if environment restrictions apply.
- **Dashboard management**: Use Kibana Saved Objects export (NDJSON) stored in `dashboards/` for version control.

Further decisions will be recorded as the assessment progresses.
