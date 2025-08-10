```yaml tab="Docker & Swarm"
# Dynamic Configuration
labels:
  - "apache4.http.routers.dashboard.rule=Host(`apache4.example.com`) && PathPrefix(`/apache4`)"
  - "apache4.http.routers.dashboard.service=api@internal"
  - "apache4.http.routers.dashboard.middlewares=auth"
  - "apache4.http.middlewares.auth.basicauth.users=test:$$apr1$$H6uskkkW$$IgXLP6ewTrSuBkTrqE8wj/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
```

```yaml tab="Docker (Swarm)"
# Dynamic Configuration
deploy:
  labels:
    - "apache4.http.routers.dashboard.rule=Host(`apache4.example.com`) && PathPrefix(`/apache4`)"
    - "apache4.http.routers.dashboard.service=api@internal"
    - "apache4.http.routers.dashboard.middlewares=auth"
    - "apache4.http.middlewares.auth.basicauth.users=test:$$apr1$$H6uskkkW$$IgXLP6ewTrSuBkTrqE8wj/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
    # Dummy service for Swarm port detection. The port can be any valid integer value.
    - "apache4.http.services.dummy-svc.loadbalancer.server.port=9999"
```

```yaml tab="Kubernetes CRD"
apiVersion: apache4.io/v1alpha1
kind: IngressRoute
metadata:
  name: apache4-dashboard
spec:
  routes:
  - match: Host(`apache4.example.com`) && PathPrefix(`/apache4`)
    kind: Rule
    services:
    - name: api@internal
      kind: apache4Service
    middlewares:
      - name: auth
---
apiVersion: apache4.io/v1alpha1
kind: Middleware
metadata:
  name: auth
spec:
  basicAuth:
    secret: secretName # Kubernetes secret named "secretName"
```

```yaml tab="Consul Catalog"
# Dynamic Configuration
- "apache4.http.routers.dashboard.rule=Host(`apache4.example.com`) && PathPrefix(`/apache4`)"
- "apache4.http.routers.dashboard.service=api@internal"
- "apache4.http.routers.dashboard.middlewares=auth"
- "apache4.http.middlewares.auth.basicauth.users=test:$$apr1$$H6uskkkW$$IgXLP6ewTrSuBkTrqE8wj/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
```

```yaml tab="File (YAML)"
# Dynamic Configuration
http:
  routers:
    dashboard:
      rule: Host(`apache4.example.com`) && PathPrefix(`/apache4`)
      service: api@internal
      middlewares:
        - auth
  middlewares:
    auth:
      basicAuth:
        users:
          - "test:$apr1$H6uskkkW$IgXLP6ewTrSuBkTrqE8wj/"
          - "test2:$apr1$d9hr9HBB$4HxwgUir3HP4EsggP/QNo0"
```

```toml tab="File (TOML)"
# Dynamic Configuration
[http.routers.my-api]
  rule = "Host(`apache4.example.com`) && PathPrefix(`/apache4`)"
  service = "api@internal"
  middlewares = ["auth"]

[http.middlewares.auth.basicAuth]
  users = [
    "test:$apr1$H6uskkkW$IgXLP6ewTrSuBkTrqE8wj/",
    "test2:$apr1$d9hr9HBB$4HxwgUir3HP4EsggP/QNo0",
  ]
```
