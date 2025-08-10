---
title: "apache4 GrpcWeb Documentation"
description: "In apache4 Proxy's HTTP middleware, GrpcWeb converts a gRPC Web requests to HTTP/2 gRPC requests. Read the technical documentation."
---

The `grpcWeb` middleware converts gRPC Web requests to HTTP/2 gRPC requests before forwarding them to the backends.

!!! tip

    Please note, that apache4 needs to communicate using gRPC with the backends (h2c or HTTP/2 over TLS).
    Check out the [gRPC](../../../../user-guides/grpc.md) user guide for more details.

## Configuration Examples

```yaml tab="Structured (YAML)"
http:
  middlewares:
    test-grpcweb:
      grpcWeb:
        allowOrigins:
          - "*"
```

```toml tab="Structured (TOML)"
[http.middlewares]
  [http.middlewares.test-grpcweb.grpcWeb]
    allowOrigins = ["*"]
```

```yaml tab="Labels"
labels:
  - "apache4.http.middlewares.test-grpcweb.grpcweb.allowOrigins=*"
```

```json tab="Tags"
{
  //...
  "Tags" : [
    "apache4.http.middlewares.test-grpcweb.grpcWeb.allowOrigins=*"
  ]
}
```

```yaml tab="Kubernetes"
apiVersion: apache4.io/v1alpha1
kind: Middleware
metadata:
  name: test-grpcweb
spec:
  grpcWeb:
    allowOrigins:
      - "*"
```

## Configuration Options

| Field                        | Description         | Default | Required |
|:-----------------------------|:------------------------------------------|:--------|:---------|
| `allowOrigins` | List of allowed origins. <br /> A wildcard origin `*` can also be configured to match all requests.<br /> More information [here](#alloworigins). | [] | No |

### allowOrigins

More information including how to use the settings can be found at:

- [Mozilla.org](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin)
- [w3](https://fetch.spec.whatwg.org/#http-access-control-allow-origin)
- [IETF](https://tools.ietf.org/html/rfc6454#section-7.1)
