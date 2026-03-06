package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/metric/global"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/semconv/v1.17.0"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func initTracer() func(context.Context) error {
    ctx := context.Background()
    // export to local agent via hostPort or service
    endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
    if endpoint == "" {
        endpoint = "otel-agent.observability.svc.cluster.local:4317"
    }

    exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(endpoint), otlptracegrpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to create exporter: %v", err)
    }

    resAttrs := resource.NewWithAttributes(
        semconv.ServiceNameKey.String("frontend"),
        semconv.ServiceVersionKey.String("1.0.0"),
        semconv.DeploymentEnvironmentKey.String("production"),
    )

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resAttrs),
    )
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.TraceContext{})
    return tp.Shutdown
}

func main() {
    shutdown := initTracer()
    defer shutdown(context.Background())

    // a custom metric counter
    meter := global.Meter("frontend")
    cartCounter, _ := meter.SyncInt64().Counter("cart.additions")

    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        // manual span for business logic
        tracer := otel.Tracer("frontend")
        ctx, span := tracer.Start(ctx, "render-homepage")
        defer span.End()

        // simulate template rendering
        span.AddEvent("render template start")
        fmt.Fprintln(w, "Welcome to Online Boutique!")
        span.AddEvent("render template end")

        // increment metric
        cartCounter.Add(ctx, 1)
    })

    wrapped := otelhttp.NewHandler(handler, "homepage")
    http.Handle("/", wrapped)

    port := "8080"
    if p := os.Getenv("PORT"); p != "" {
        port = p
    }
    log.Printf("starting frontend on %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
