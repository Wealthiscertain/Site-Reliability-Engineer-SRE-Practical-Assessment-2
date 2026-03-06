# SRE Assessment Repository

This repository contains the implementation work for the ITStack Limited Site Reliability Engineer practical assessment. It is structured to cover all required sections: OpenTelemetry collector configuration, multi-language service instrumentation, RUM integration, dashboard exports, and infrastructure monitoring configurations.

## Directory layout

```text
sre-assessment/
├── otel-collector/                # Helm values, collector configs
│   ├── values-agent.yaml
│   ├── values-gateway.yaml
│   └── sampling-policy.yaml
├── instrumentation/               # Per-service instrumentation code/patches
│   ├── frontend/
│   ├── cartservice/
│   └── paymentservice/
├── rum/                           # Browser SDK integration code
├── dashboards/                    # Kibana Saved Objects (NDJSON exports)
│   ├── service-health.ndjson
│   ├── rum-performance.ndjson
│   └── business-transactions.ndjson
├── infrastructure/                # Agent/Beat configs, alert rules
│   ├── elastic-agent-policies/    # Fleet policy exports or agent.yml
│   ├── postgres-integration/
│   ├── redis-integration/
│   ├── nginx-integration/
│   └── alerting-rules/            # Kibana rule exports (NDJSON)
├── docs/
│   └── DECISIONS.md               # Architectural decision log
└── README.md                      # Setup instructions
```

## Getting Started

1. Clone the repository and ensure you have access to a Kubernetes cluster with at least two nodes. The Google Online Boutique sample application should be deployed there.
2. Follow the directories to configure and deploy components:
   - `otel-collector/` contains Helm values for deploying the OpenTelemetry Collector.
   - `instrumentation/` holds language-specific code snippets and guidance to instrument services.
   - `rum/` covers the frontend RUM agent setup.
   - `dashboards/` holds exported NDJSON objects that can be imported in Kibana.
   - `infrastructure/` includes Elastic Agent or Metricbeat configurations and alert rules.
3. The `docs/DECISIONS.md` file records architectural choices made during implementation.

## Submission

Once complete, push the repository to a public Git platform and provide the link as instructed in the assessment email.
