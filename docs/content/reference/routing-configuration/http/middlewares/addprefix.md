---
title: "apache4 AddPrefix Documentation"
description: "Learn how to implement the HTTP AddPrefix middleware in apache4 Proxy to updates request paths before being forwarded. Read the technical documentation."
---

![AddPrefix](../../../../assets/img/middleware/addprefix.png)

The `addPrefix` middleware updates the path of a request before forwarding it.

## Configuration Examples

```yaml tab="Structured (YAML)"
# Prefixing with /foo
http:
  middlewares:
    add-foo:
      addPrefix:
        prefix: "/foo"
```

```toml tab="Structured (TOML)"
# Prefixing with /foo
[http.middlewares]
  [http.middlewares.add-foo.addPrefix]
    prefix = "/foo"
```

```yaml tab="Labels"
# Prefixing with /foo
labels:
  - "apache4.http.middlewares.add-foo.addprefix.prefix=/foo"
```

```json tab="Tags"
// Prefixing with /foo
{
  // ...
  "Tags": [
    "apache4.http.middlewares.add-foo.addprefix.prefix=/foo"
  ]
}
```

```yaml tab="Kubernetes"
# Prefixing with /foo
apiVersion: apache4.io/v1alpha1
kind: Middleware
metadata:
  name: add-foo
spec:
  addPrefix:
    prefix: /foo
```

## Configuration Options

| Field  | Description                                                                                                                                                                                                | Default | Required |
|:-----------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------|:---------|
| `prefix` | String to add **before** the current path in the requested URL. It should include a leading slash (`/`). | "" | Yes |
