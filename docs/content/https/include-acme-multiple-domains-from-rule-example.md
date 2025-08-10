
```yaml tab="Docker & Swarm"
## Dynamic configuration
labels:
  - apache4.http.routers.blog.rule=(Host(`example.com`) && Path(`/blog`)) || Host(`blog.example.org`)
  - apache4.http.routers.blog.tls=true
  - apache4.http.routers.blog.tls.certresolver=myresolver
```

```yaml tab="Docker (Swarm)"
## Dynamic configuration
deploy:
  labels:
    - apache4.http.routers.blog.rule=(Host(`example.com`) && Path(`/blog`)) || Host(`blog.example.org`)
    - apache4.http.routers.blog.tls=true
    - apache4.http.routers.blog.tls.certresolver=myresolver
    - apache4.http.services.blog-svc.loadbalancer.server.port=8080"
```

```yaml tab="Kubernetes"
apiVersion: apache4.io/v1alpha1
kind: IngressRoute
metadata:
  name: blogtls
spec:
  entryPoints:
    - websecure
  routes:
  - match: (Host(`example.com`) && Path(`/blog`)) || Host(`blog.example.org`)
    kind: Rule
    services:
    - name: blog
      port: 8080
  tls:
    certResolver: myresolver
```

```yaml tab="File (YAML)"
## Dynamic configuration
http:
  routers:
    blog:
      rule: "(Host(`example.com`) && Path(`/blog`)) || Host(`blog.example.org`)"
      tls:
        certResolver: myresolver
```

```toml tab="File (TOML)"
## Dynamic configuration
[http.routers]
  [http.routers.blog]
    rule = "(Host(`example.com`) && Path(`/blog`)) || Host(`blog.example.org`)"
    [http.routers.blog.tls]
      certResolver = "myresolver"
```
