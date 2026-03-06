using System;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using OpenTelemetry;
using OpenTelemetry.Trace;
using OpenTelemetry.Resources;
using OpenTelemetry.Metrics;

namespace CartService
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateBuilder(args);

            builder.Services.AddOpenTelemetryTracing(tracerProviderBuilder =>
            {
                tracerProviderBuilder
                    .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService("cartservice", serviceVersion: "1.0.0", serviceEnvironment: "production"))
                    .AddAspNetCoreInstrumentation()
                    .AddRedisInstrumentation()
                    .AddHttpClientInstrumentation()
                    .AddOtlpExporter(options =>
                    {
                        options.Endpoint = new Uri(builder.Configuration["OTEL_EXPORTER_OTLP_ENDPOINT"] ?? "http://otel-agent.observability.svc.cluster.local:4317");
                    });
            });

            builder.Services.AddOpenTelemetryMetrics(metricsBuilder =>
            {
                metricsBuilder
                    .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService("cartservice"))
                    .AddMeter("cartservice")
                    .AddOtlpExporter();
            });

            var app = builder.Build();

            app.MapGet("/add-item", (HttpContext http) =>
            {
                var tracer = TracerProvider.Default.GetTracer("cartservice");
                using var span = tracer.StartActiveSpan("validate-cart-contents", out var s);
                s.SetAttribute("user.id", "12345");
                s.SetAttribute("cart.size", 3);
                // ... business logic
                return "item added";
            });

            app.Run();
        }
    }
}
