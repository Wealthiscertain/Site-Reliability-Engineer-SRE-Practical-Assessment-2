# OpenTelemetry Collector Deployment

This directory contains Helm values used to deploy a dual-topology collector into the
Kubernetes cluster: a DaemonSet agent on every node and a central gateway Deployment.
All telemetry is forwarded via OTLP to the Elastic APM Server.

## Deployment steps

1. Add the OpenTelemetry Helm repository if not already present:
   ```sh
   helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
   helm repo update
   ```

2. Create a namespace for the collectors (example `observability`):
   ```sh
   kubectl create namespace observability
   ```

3. Deploy the gateway:
   ```sh
   helm upgrade --install otel-gateway open-telemetry/opentelemetry-collector \
     --namespace observability \
     --values ./otel-collector/values-gateway.yaml
   ```

4. Deploy the agent DaemonSet:
   ```sh
   helm upgrade --install otel-agent open-telemetry/opentelemetry-collector \
     --namespace observability \
     --values ./otel-collector/values-agent.yaml
   ```

5. Verify the components are running:
   ```sh
   kubectl get pods -n observability
   kubectl get svc -n observability
   ```

6. The gateway exposes health zpages on port 8888; you can port-forward to inspect:
   ```sh
   kubectl port-forward svc/otel-gateway 8888:8888 -n observability
   # visit http://localhost:8888/debug/tracez
   ```

7. Ensure the gateway service name matches the `gatewayService` value in the agent
   values file (defaults to `otel-gateway.otlp.svc.cluster.local`).

8. Create a Kubernetes Secret for the Elastic APM API key used by the gateway exporter:
   ```sh
   kubectl create secret generic elastic-apm-token \
     --namespace observability \
     --from-literal=ELASTIC_APM_API_KEY="<your-api-key>"
   ```
   and reference it in a deployment/VarENV or set `ELASTIC_APM_API_KEY` in the gateway
   environment via Helm chart values.

9. After services and instrumented applications send telemetry, verify data arrives in
   Kibana under Observability → APM → Services. You should see the collector itself
   (e.g., `otel-gateway`) showing internal metrics and traces.

## Configuration details

Both `values-*.yaml` files configure:

* **Receivers**: otlp (grpc/http), zipkin, and hostmetrics on agents.
* **Processors**: batch for throughput, memory_limiter to guard against leaks, resource
  detection to enrich with Kubernetes attributes, and tail-based sampling on the gateway.
* **Exporters**: OTLP pointing at the Elastic APM Server; TLS enabled with API key.
* **Pipelines**: separate pipelines for traces, metrics, and logs; each includes the
  appropriate processors and exports to the `otlp` exporter.

### Sampling policy
Errors are always sampled (100%). A probabilistic policy samples 10% of non-error
requests, reducing data volume while preserving representative traces. See
`sampling-policy.yaml` for documentation.

### Resource enrichment
The `resource` processor inserts attributes such as
`k8s.pod.name`, `k8s.namespace.name`, `k8s.deployment.name`, and `service.name` to
facilitate filtering in Kibana.

### Monitoring the collector
The gateway exposes Prometheus-friendly metrics and zpages; scrape them with
Prometheus or view via port-forward. Agents also forward hostmetrics to the gateway.

---

Once the collectors are deployed and verified, you can proceed to instrument the
microservices so that telemetry flows through this pipeline.