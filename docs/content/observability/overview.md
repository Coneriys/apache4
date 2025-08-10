---
title: "apache4 Observability Overview"
description: "apache4 provides Logs, Access Logs, Metrics and Tracing. Read the full documentation to get started."
---

# Overview

apache4’s observability features include logs, access logs, metrics, and tracing. You can configure these options globally or at more specific levels, such as per router or per entry point.

## Configuration Example

Enable access logs, metrics, and tracing globally

```yaml tab="File (YAML)"
accessLog: {}

metrics:
  otlp: {}

tracing: {}
```

```yaml tab="File (TOML)"
[accessLog]

[metrics]
  [metrics.otlp]

[tracing]
```

```bash tab="CLI"
--accesslog=true
--metrics.otlp=true
--tracing=true
```

You can disable access logs, metrics, and tracing for a specific entrypoint attached to a router:

```yaml tab="File (YAML)"
# Static Configuration
entryPoints:
  EntryPoint0:
    address: ':8000/udp'
    observability:
      accessLogs: false
      tracing: false
      metrics: false
```

```toml tab="File (TOML)"
# Static Configuration
[entryPoints.EntryPoint0]
  address = ":8000/udp"

    [entryPoints.EntryPoint0.observability]
      accessLogs = false
      tracing = false
      metrics = false
```

```bash tab="CLI"
# Static Configuration
--entryPoints.EntryPoint0.address=:8000/udp
--entryPoints.EntryPoint0.observability.accessLogs=false
--entryPoints.EntryPoint0.observability.metrics=false
--entryPoints.EntryPoint0.observability.tracing=false
```

!!!note "Default Behavior"
    A router with its own observability configuration will override the global default.

## Configuration Options

### Logs

apache4 logs informs about everything that happens within apache4 (startup, configuration, events, shutdown, and so on).

Read the [Logs documentation](./logs.md) to learn how to configure it.

### Access Logs

Access logs are a key part of observability in apache4.

They are providing valuable insights about incoming traffic, and allow to monitor it.
The access logs record detailed information about each request received by apache4,
including the source IP address, requested URL, response status code, and more.

Read the [Access Logs documentation](./access-logs.md) to learn how to configure it.

### Metrics

apache4 offers a metrics feature that provides valuable insights about the performance and usage.
These metrics include the number of requests received, the requests duration, and more.

On top of supporting metrics in the OpenTelemetry format, apache4 supports the following vendor specific metrics systems: Prometheus, Datadog, InfluxDB 2.X, and StatsD.

Read the [Metrics documentation](./metrics/overview.md) to learn how to configure it.

### Tracing

The apache4 tracing system allows developers to gain deep visibility into the flow of requests through their infrastructure.

apache4 provides tracing information in the OpenTelemery format.

Read the [Tracing documentation](./tracing/overview.md) to learn how to configure it.
