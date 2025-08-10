---
title: "apache4 Configuration Overview"
description: "Read the official apache4 documentation to get started with configuring the apache4 Proxy."
---

# Boot Environment

apache4 Proxy’s configuration is divided into two main categories:

- **Static Configuration**: Defines parameters that require apache4 to restart when changed. This includes entry points, providers, API/dashboard settings, and logging levels.
- **Dynamic Configuration**: Involves elements that can be updated without restarting apache4, such as routers, services, and middlewares.

This section focuses on setting up the static configuration, which is essential for apache4’s initial boot.

## Configuration Methods

apache4 offers multiple methods to define static configuration. 

!!! warning "Note"
    It’s crucial to choose one method and stick to it, as mixing different configuration options is not supported and can lead to unexpected behavior.

Here are the methods available for configuring the apache4 proxy:

- [File](#file) 
- [CLI](#cli)
- [Environment Variables](#environment-variables)
- [Helm](#helm)

## File

You can define the static configuration in a file using formats like YAML or TOML.

### Configuration Example

```yaml tab="apache4.yml (YAML)"
entryPoints:
  web:
    address: ":80"
  websecure:
    address: ":443"

providers:
  docker: {}

api:
  dashboard: true

log:
  level: INFO
```

```toml tab="apache4.toml (TOML)"
[entryPoints]
  [entryPoints.web]
    address = ":80"

  [entryPoints.websecure]
    address = ":443"

[providers]
  [providers.docker]

[api]
  dashboard = true

[log]
  level = "INFO"
```

### Configuration File

At startup, apache4 searches for static configuration in a file named `apache4.yml` (or `apache4.yaml` or `apache4.toml`) in the following directories:

- `/etc/apache4/`
- `$XDG_CONFIG_HOME/`
- `$HOME/.config/`
- `.` (the current working directory).

You can override this behavior using the `configFile` argument like this:

```bash
apache4 --configFile=foo/bar/myconfigfile.yml
```

## CLI

Using the CLI, you can pass static configuration directly as command-line arguments when starting apache4. 

### Configuration Example

```sh tab="CLI"
apache4 \
  --entryPoints.web.address=":80" \
  --entryPoints.websecure.address=":443" \
  --providers.docker \
  --api.dashboard \
  --log.level=INFO
```

## Environment Variables

You can also set the static configuration using environment variables. Each option corresponds to an environment variable prefixed with `apache4_`.

### Configuration Example

```sh tab="ENV"
apache4_ENTRYPOINTS_WEB_ADDRESS=":80" apache4_ENTRYPOINTS_WEBSECURE_ADDRESS=":443" apache4_PROVIDERS_DOCKER=true apache4_API_DASHBOARD=true apache4_LOG_LEVEL="INFO" apache4
```

## Helm

When deploying apache4 Proxy using Helm in a Kubernetes cluster, the static configuration is defined in a `values.yaml` file. 

You can find the official apache4 Helm chart on [GitHub](https://github.com/apache4/apache4-helm-chart/blob/master/apache4/VALUES.md)

### Configuration Example

```yaml tab="values.yaml"
ports:
  web:
    exposedPort: 80
  websecure:
    exposedPort: 443

additionalArguments:
  - "--providers.kubernetescrd.ingressClass"
  - "--log.level=INFO"
```

```sh tab="Helm Commands"
helm repo add apache4 https://apache4.github.io/charts
helm repo update
helm install apache4 apache4/apache4 -f values.yaml
```
