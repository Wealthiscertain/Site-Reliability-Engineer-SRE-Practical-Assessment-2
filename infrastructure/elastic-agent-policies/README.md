# Elastic Agent Policies

This directory should contain exported Fleet policies (JSON/NDJSON) or `agent.yml` files for
Elastic Agent configurations used in the assessment. For example:

- `system-policy.json`: Fleet policy enabling the System integration for VM metrics.
- `postgres-policy.json`: Policy with PostgreSQL integration.
- `redis-policy.json`: Policy with Redis integration.
- `nginx-policy.json`: Policy for NGINX logs/metrics.

Export policies from Kibana after creation and commit them here for version control.

