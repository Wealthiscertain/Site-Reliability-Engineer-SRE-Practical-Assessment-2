# Real User Monitoring (RUM) integration

This directory holds files and instructions for integrating an Elastic APM RUM agent into the frontend service.

### Approach
- Use the `@elastic/apm-rum` package in the web application.
- Configure the agent with `serviceName`, `serverUrl`, and `environment`.
- Ensure CORS is configured on the APM Server to allow the frontend origin.
- Capture page-load, XHR/fetch, user interactions, and Core Web Vitals.
- Propagate tracing headers in outgoing requests to backend services for correlation.

Example initialization (e.g., in a script included by the frontend HTML):

```html
<script src="/path/to/elastic-apm-rum.umd.min.js"></script>
<script>
  var apm = elasticApm.init({
    serviceName: 'frontend',
    serverUrl: 'https://elastic-apm-server:8200',
    environment: 'production'
  });
</script>
```

Additional JS files with helper functions can be added in this directory.
