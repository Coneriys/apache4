---
title: "Per-Router Observability"
description: "You can disable access logs, metrics, and tracing for a specific entrypoint attached to a HTTP Router. Read the technical documentation."
---

apache4's observability features include logs, access logs, metrics, and tracing. You can configure these options globally or at more specific levels, such as per router or per entry point.

By default, the router observability configuration is inherited from the attached EntryPoints and can be configured with the observability [options](../../../install-configuration/entrypoints.md#configuration-options)).
However, a router defining its own observability configuration will opt-out from these defaults.

!!! info
    To enable router-level observability, you must first enable access-logs, tracing, and metrics.

    When metrics layers are not enabled with the `addEntryPointsLabels`, `addRoutersLabels` and/or `addServicesLabels` options,
    enabling metrics for a router will not enable them.

!!! warning "AddInternals option"

    By default, and for any type of signal (access-logs, metrics and tracing),
    apache4 disables observability for internal resources.
    The observability options described below cannot interfere with the `AddInternals` ones,
    and will be ignored.

    For instance, if a router exposes the `api@internal` service and `metrics.AddInternals` is false,
    it will never produces metrics, even if the router observability configuration enables metrics.

## Configuration Example

```yaml tab="Structured (YAML)"
http:
  routers:
    my-router:
      rule: "Path(`/foo`)"
      service: service-foo
      observability:
        metrics: false
        accessLogs: false
        tracing: false
        traceVerbosity: detailed
```

```yaml tab="Structured (TOML)"
[http.routers.my-router]
  rule = "Path(`/foo`)"
  service = "service-foo"

  [http.routers.my-router.observability]
    metrics = false
    accessLogs = false
    tracing = false
    traceVerbosity = "detailed"
```

```yaml tab="Labels"
labels:
  - "apache4.http.routers.my-router.rule=Path(`/foo`)"
  - "apache4.http.routers.my-router.service=service-foo"
  - "apache4.http.routers.my-router.observability.metrics=false"
  - "apache4.http.routers.my-router.observability.accessLogs=false"
  - "apache4.http.routers.my-router.observability.tracing=false"
  - "apache4.http.routers.my-router.observability.traceVerbosity=detailed"
```

```json tab="Tags"
{
  // ...
  "Tags": [
    "apache4.http.routers.my-router.rule=Path(`/foo`)",
    "apache4.http.routers.my-router.service=service-foo",
    "apache4.http.routers.my-router.observability.metrics=false",
    "apache4.http.routers.my-router.observability.accessLogs=false",
    "apache4.http.routers.my-router.observability.tracing=false",
    "apache4.http.routers.my-router.observability.traceVerbosity=detailed"
  ]
}
```

## Configuration Options

| Field            | Description                                                                                                                                                                                | Default   | Required |
|:-----------------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------|:---------|
| `accessLogs`     | The `accessLogs` option controls whether the router will produce access-logs.                                                                                                              | `true`    | No       |
| `metrics`        | The `metrics` option controls whether the router will produce metrics.                                                                                                                     | `true`    | No       |
| `tracing`        | The `tracing` option controls whether the router will produce traces.                                                                                                                      | `true`    | No       |
| `traceVerbosity` | The `traceVerbosity` option controls the tracing verbosity level for the router. Possible values: `minimal` (default), `detailed`. If not set, the value is inherited from the entryPoint. | `minimal` | No       |

#### traceVerbosity

`observability.traceVerbosity` defines the tracing verbosity level for the router.

Possible values are:

- `minimal`: produces a single server span and one client span for each request processed by a router.
- `detailed`: enables the creation of additional spans for each middleware executed for each request processed by a router.
