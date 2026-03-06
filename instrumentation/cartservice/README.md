# Cartservice (C# .NET) instrumentation

The cart service uses .NET; auto-instrumentation can be enabled with the Elastic APM .NET agent or OpenTelemetry .NET SDK.

### Steps
1. Install OpenTelemetry NuGet packages and configure the `TracerProvider` with resource attributes.
2. Enable auto-instrumentation for ASP.NET Core and Redis client (StackExchange.Redis).
3. Add custom metrics using `Meter` for cart operations and manual spans for business logic.
4. Send OTLP to the local agent service endpoint.

Example snippet in `Program.cs` or `Startup.cs`:

```csharp
// TODO: create file instrumentation/cartservice/Program.cs with actual initialization code
```