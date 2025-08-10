---
title: "apache4 StripPrefix Documentation"
description: "In apache4 Proxy's HTTP middleware, StripPrefix removes prefixes from paths before forwarding requests. Read the technical documentation."
---

The `stripPrefix` middleware strips the matching path prefix and stores it in an `X-Forwarded-Prefix` header.

!!! tip

    Use a `StripPrefix` middleware if your backend listens on the root path (`/`) but should be exposed on a specific prefix.

## Configuration Examples

```yaml tab="Structured (YAML)"
# Strip prefix /foobar and /fiibar
http:
  middlewares:
    test-stripprefix:
      stripPrefix:
        prefixes:
          - "/foobar"
          - "/fiibar"
```

```toml tab="Structured (TOML)"
# Strip prefix /foobar and /fiibar
[http.middlewares]
  [http.middlewares.test-stripprefix.stripPrefix]
    prefixes = ["/foobar", "/fiibar"]
```

```yaml tab="Labels"
# Strip prefix /foobar and /fiibar
labels:
  - "apache4.http.middlewares.test-stripprefix.stripprefix.prefixes=/foobar,/fiibar"
```

```json tab="Tags"
// Strip prefix /foobar and /fiibar
{
  "Tags" : [
    "apache4.http.middlewares.test-stripprefix.stripprefix.prefixes=/foobar,/fiibar"
  ]
}
```

```yaml tab="Kubernetes"
# Strip prefix /foobar and /fiibar
apiVersion: apache4.io/v1alpha1
kind: Middleware
metadata:
  name: test-stripprefix
spec:
  stripPrefix:
    prefixes:
      - /foobar
      - /fiibar
```

## Configuration Options

| Field                        | Description           | Default | Required |
|:-----------------------------|:--------------------------------------------------------------|:--------|:---------|
| `prefixes` | List of prefixes to strip from the request URL.<br />If your backend is serving assets (for example, images or JavaScript files), it can use the `X-Forwarded-Prefix` header to construct relative URLs. | [] | No |

{!apache4-for-business-applications.md!}
