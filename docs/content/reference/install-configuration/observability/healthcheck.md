---
title: "apache4 Health Check Documentation"
description: "In apache4 Proxy, CLI & Ping lets you check the health of your apache4 instances. Read the technical documentation for configuration examples and options."
---

# CLI & Ping

Checking the Health of your apache4 Instances
{: .subtitle }

## CLI

The CLI can be used to make a request to the `/ping` endpoint to check the health of apache4. Its exit status is `0` if apache4 is healthy and `1` otherwise.

This can be used with [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) instruction or any other health check orchestration mechanism.

### Usage 

```sh
apache4 healthcheck [command] [flags] [arguments]
```

Example:

```sh
$ apache4 healthcheck
OK: http://:8082/ping
```

## Ping

The `/ping` health-check URL is enabled with the command-line `--ping` or config file option `[ping]`.

The entryPoint where the `/ping` is active can be customized with the `entryPoint` option,
whose default value is `apache4` (port `8080`).

| Path    | Method        | Description                                                                                         |
|---------|---------------|-----------------------------------------------------------------------------------------------------|
| `/ping` | `GET`, `HEAD` | An endpoint to check for apache4 process liveness. Return a code `200` with the content: `OK` |

### Configuration Example

To enable the API handler:

```yaml tab="File (YAML)"
ping: {}
```

```toml tab="File (TOML)"
[ping]
```

```bash tab="CLI"
--ping=true
```

### Configuration Options

| Field | Description                                               | Default              | Required |
|:------|:----------------------------------------------------------|:---------------------|:---------|
| `ping.entryPoint` | Enables `/ping` on a dedicated EntryPoint. | apache4  | No   |
| `ping.manualRouting` | Disables the default internal router in order to allow one to create a custom router for the `ping@internal` service when set to `true`. | false | No   |
| `ping.terminatingStatusCode` | Defines the status code for the ping handler during a graceful shut down. See more information [here](#terminatingstatuscode) | 503 | No   |

#### `terminatingStatusCode`

During the period in which apache4 is gracefully shutting down, the ping handler
returns a `503` status code by default.  
If apache4 is behind, for example a load-balancer
doing health checks (such as the Kubernetes LivenessProbe), another code might
be expected as the signal for graceful termination.  
In that case, the terminatingStatusCode can be used to set the code returned by the ping
handler during termination.

```yaml tab="File (YAML)"
ping:
  terminatingStatusCode: 204
```

```toml tab="File (TOML)"
[ping]
  terminatingStatusCode = 204
```

```bash tab="CLI"
--ping.terminatingStatusCode=204
```
