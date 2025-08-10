---
title: "apache4 Routing Configuration with KV stores"
description: "Read the technical documentation to learn the apache4 Routing Configuration with KV stores."
---

# apache4 & KV Stores

## Routing Configuration

!!! info "Keys"

    Keys are case-insensitive.

### Routers

!!! warning "The character `@` is not authorized in the router name `<router_name>`."

??? info "`apache4/http/routers/<router_name>/rule`"

    See [rule](../http/router/rules-and-priority.md#rules) for more information.
    
    | Key (Path)                           | Value                      |
    |--------------------------------------|----------------------------|
    | `apache4/http/routers/myrouter/rule` | ```Host(`example.com`)```  |

??? info "`apache4/http/routers/<router_name>/ruleSyntax`"

    !!! warning

        RuleSyntax option is deprecated and will be removed in the next major version.
        Please do not use this field and rewrite the router rules to use the v3 syntax.

    See [rule](../http/router/rules-and-priority.md#rulesyntax) for more information.
    
    | Key (Path)                           | Value                      |
    |--------------------------------------|----------------------------|
    | `apache4/http/routers/myrouter/ruleSyntax` | `v3`  |

??? info "`apache4/http/routers/<router_name>/entrypoints`"

    See [entry points](../../install-configuration/entrypoints.md) for more information.

    | Key (Path)                                    | Value       |
    |-----------------------------------------------|-------------|
    | `apache4/http/routers/myrouter/entrypoints/0` | `web`       |
    | `apache4/http/routers/myrouter/entrypoints/1` | `websecure` |

??? info "`apache4/http/routers/<router_name>/middlewares`"

    See [middlewares overview](../http/middlewares/overview.md) for more information.

    | Key (Path)                                    | Value       |
    |-----------------------------------------------|-------------|
    | `apache4/http/routers/myrouter/middlewares/0` | `auth`      |
    | `apache4/http/routers/myrouter/middlewares/1` | `prefix`    |
    | `apache4/http/routers/myrouter/middlewares/2` | `cb`        |

??? info "`apache4/http/routers/<router_name>/service`"

    See [service](../http/load-balancing/service.md) for more information.

    | Key (Path)                              | Value       |
    |-----------------------------------------|-------------|
    | `apache4/http/routers/myrouter/service` | `myservice` |

??? info "`apache4/http/routers/<router_name>/tls`"

    See [tls](../http/tls/overview.md) for more information.

    | Key (Path)                          | Value  |
    |-------------------------------------|--------|
    | `apache4/http/routers/myrouter/tls` | `true` |
    
??? info "`apache4/http/routers/<router_name>/tls/certresolver`"

    See [certResolver](../../install-configuration/tls/certificate-resolvers/overview.md) for more information.

    | Key (Path)                                       | Value        |
    |--------------------------------------------------|--------------|
    | `apache4/http/routers/myrouter/tls/certresolver` | `myresolver` |    

??? info "`apache4/http/routers/<router_name>/tls/domains/<n>/main`"

    See [domains](../../install-configuration/tls/certificate-resolvers/acme.md#domain-definition) for more information.

    | Key (Path)                                         | Value         |
    |----------------------------------------------------|---------------|
    | `apache4/http/routers/myrouter/tls/domains/0/main` | `example.org` |
    
??? info "`apache4/http/routers/<router_name>/tls/domains/<n>/sans/<n>`"

    See [domains](../../install-configuration/tls/certificate-resolvers/acme.md#domain-definition) for more information.

    | Key (Path)                                           | Value              |
    |------------------------------------------------------|--------------------|
    | `apache4/http/routers/myrouter/tls/domains/0/sans/0` | `test.example.org` |
    | `apache4/http/routers/myrouter/tls/domains/0/sans/1` | `dev.example.org`  |
    
??? info "`apache4/http/routers/<router_name>/tls/options`"

    See [TLS](../http/tls/overview.md) for more information.

    | Key (Path)                                  | Value    |
    |---------------------------------------------|----------|
    | `apache4/http/routers/myrouter/tls/options` | `foobar` |

??? info "`apache4/http/routers/<router_name>/observability/accesslogs`"

    The accessLogs option controls whether the router will produce access-logs.
      
    | Key (Path)                                               | Value  |
    |----------------------------------------------------------|--------|
    | `apache4/http/routers/myrouter/observability/accesslogs` | `true` |

??? info "`apache4/http/routers/<router_name>/observability/metrics`"

    The metrics option controls whether the router will produce metrics.

    | Key (Path)                                            | Value  |
    |-------------------------------------------------------|--------|
    | `apache4/http/routers/myrouter/observability/metrics` | `true` |

??? info "`apache4/http/routers/<router_name>/observability/tracing`"

    The tracing option controls whether the router will produce traces.
    
    | Key (Path)                                            | Value  |
    |-------------------------------------------------------|--------|
    | `apache4/http/routers/myrouter/observability/tracing` | `true` |

??? info "`apache4/http/routers/<router_name>/priority`"

    See [domains](../../install-configuration/tls/certificate-resolvers/acme.md#domain-definition) for more information.

    | Key (Path)                               | Value |
    |------------------------------------------|-------|
    | `apache4/http/routers/myrouter/priority` | `42`  |

### Services

!!! warning "The character `@` is not authorized in the service name `<service_name>`."

??? info "`apache4/http/services/<service_name>/loadbalancer/servers/<n>/url`"

    See [servers](../http/load-balancing/service.md#servers) for more information.

    | Key (Path)                                                      | Value                                   |
    |-----------------------------------------------------------------|-----------------------------------------|
    | `apache4/http/services/myservice/loadbalancer/servers/0/url`    | `http://<ip-server-1>:<port-server-1>/` |

??? info "`apache4/http/services/<service_name>/loadbalancer/servers/<n>/preservePath`"

    See [servers](../http/load-balancing/service.md#servers) for more information.

    | Key (Path)                                                      | Value                                   |
    |-----------------------------------------------------------------|-----------------------------------------|
    | `apache4/http/services/myservice/loadbalancer/servers/0/preservePath`    | `true` |

??? info "`apache4/http/services/<service_name>/loadbalancer/servers/<n>/weight`"

    See [servers](../http/load-balancing/service.md#servers) for more information.

    | Key (Path)                                                      | Value                                   |
    |-----------------------------------------------------------------|-----------------------------------------|
    | `apache4/http/services/myservice/loadbalancer/servers/0/weight`    | `1` |

??? info "`apache4/http/services/<service_name>/loadbalancer/serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../http/load-balancing/serverstransport.md) for more information.

    | Key (Path)                                                      | Value         |
    |-----------------------------------------------------------------|---------------|
    | `apache4/http/services/myservice/loadbalancer/serverstransport` | `foobar@file` |

??? info "`apache4/http/services/<service_name>/loadbalancer/passhostheader`"

    | Key (Path)                                                      | Value  |
    |-----------------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/passhostheader`   | `true` |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/headers/<header_name>`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                               | Value    |
    |--------------------------------------------------------------------------|----------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/headers/X-Foo` | `foobar` |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/hostname`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                          | Value         |
    |---------------------------------------------------------------------|---------------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/hostname` | `example.org` |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/interval`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                          | Value |
    |---------------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/interval` | `10`  |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/unhealthyinterval`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                                   | Value |
    |------------------------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/unhealthyinterval` | `10`  |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/path`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                      | Value  |
    |-----------------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/path` | `/foo` |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/method`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/method` | `foobar` |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/status`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                        | Value |
    |-------------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/status` | `42`  |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/port`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                      | Value |
    |-----------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/port` | `42`  |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/scheme`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                        | Value  |
    |-------------------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/scheme` | `http` |

??? info "`apache4/http/services/<service_name>/loadbalancer/healthcheck/timeout`"

    See [health check](../http/load-balancing/service.md#health-check) for more information.

    | Key (Path)                                                         | Value |
    |--------------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/healthcheck/timeout` | `10`  |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky`"

    | Key (Path)                                            | Value  |
    |-------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/sticky` | `true` |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky/cookie/httponly`"

    | Key (Path)                                                            | Value  |
    |-----------------------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/sticky/cookie/httponly` | `true` |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky/cookie/name`"

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `apache4/http/services/myservice/loadbalancer/sticky/cookie/name` | `foobar` |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky/cookie/path`"

    | Key (Path)                                                        | Value     |
    |-------------------------------------------------------------------|-----------|
    | `apache4/http/services/myservice/loadbalancer/sticky/cookie/path` | `/foobar` |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky/cookie/secure`"

    | Key (Path)                                                          | Value  |
    |---------------------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/sticky/cookie/secure` | `true` |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky/cookie/samesite`"

    | Key (Path)                                                            | Value  |
    |-----------------------------------------------------------------------|--------|
    | `apache4/http/services/myservice/loadbalancer/sticky/cookie/samesite` | `none` |

??? info "`apache4/http/services/<service_name>/loadbalancer/sticky/cookie/maxage`"

    | Key (Path)                                                          | Value |
    |---------------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/sticky/cookie/maxage` | `42`  |

??? info "`apache4/http/services/<service_name>/loadbalancer/responseforwarding/flushinterval`"

    | Key (Path)                                                                      | Value |
    |---------------------------------------------------------------------------------|-------|
    | `apache4/http/services/myservice/loadbalancer/responseforwarding/flushinterval` | `10`  |

??? info "`apache4/http/services/<service_name>/mirroring/service`"

    | Key (Path)                                               | Value    |
    |----------------------------------------------------------|----------|
    | `apache4/http/services/<service_name>/mirroring/service` | `foobar` |

??? info "`apache4/http/services/<service_name>/mirroring/mirrors/<n>/name`"

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `apache4/http/services/<service_name>/mirroring/mirrors/<n>/name` | `foobar` |

??? info "`apache4/http/services/<service_name>/mirroring/mirrors/<n>/percent`"

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `apache4/http/services/<service_name>/mirroring/mirrors/<n>/percent` | `42`  |

??? info "`apache4/http/services/<service_name>/weighted/services/<n>/name`"

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `apache4/http/services/<service_name>/weighted/services/<n>/name` | `foobar` |

??? info "`apache4/http/services/<service_name>/weighted/services/<n>/weight`"

    | Key (Path)                                                          | Value |
    |---------------------------------------------------------------------|-------|
    | `apache4/http/services/<service_name>/weighted/services/<n>/weight` | `42`  |

??? info "`apache4/http/services/<service_name>/weighted/sticky/cookie/name`"

    | Key (Path)                                                         | Value    |
    |--------------------------------------------------------------------|----------|
    | `apache4/http/services/<service_name>/weighted/sticky/cookie/name` | `foobar` |

??? info "`apache4/http/services/<service_name>/weighted/sticky/cookie/secure`"

    | Key (Path)                                                           | Value  |
    |----------------------------------------------------------------------|--------|
    | `apache4/http/services/<service_name>/weighted/sticky/cookie/secure` | `true` |

??? info "`apache4/http/services/<service_name>/weighted/sticky/cookie/samesite`"

    | Key (Path)                                                             | Value  |
    |------------------------------------------------------------------------|--------|
    | `apache4/http/services/<service_name>/weighted/sticky/cookie/samesite` | `none` |

??? info "`apache4/http/services/<service_name>/weighted/sticky/cookie/httpOnly`"

    | Key (Path)                                                             | Value  |
    |------------------------------------------------------------------------|--------|
    | `apache4/http/services/<service_name>/weighted/sticky/cookie/httpOnly` | `true` |

??? info "`apache4/http/services/<service_name>/weighted/sticky/cookie/maxage`"

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `apache4/http/services/<service_name>/weighted/sticky/cookie/maxage` | `42`  |

??? info "`apache4/http/services/<service_name>/failover/fallback`"

    See [Failover](../http/load-balancing/service.md#failover) for more information

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `apache4/http/services/<service_name>/failover/fallback` | `backup`  |

??? info "`apache4/http/services/<service_name>/failover/healthcheck`"

    See [Failover](../http/load-balancing/service.md#failover) for more information

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `apache4/http/services/<service_name>/failover/healthcheck` | `{}`  |

??? info "`apache4/http/services/<service_name>/failover/service`"

    See [Failover](../http/load-balancing/service.md#failover) for more information

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `apache4/http/services/<service_name>/failover/service` | `main`  |

### Middleware

More information about available middlewares in the dedicated [middlewares section](../http/middlewares/overview.md).

!!! warning "The character `@` is not authorized in the middleware name."

!!! warning "Conflicts in Declaration"

    If you declare multiple middleware with the same name but with different parameters, the middleware fails to be declared.

### TCP

You can declare TCP Routers and/or Services using KV.

#### TCP Routers

??? info "`apache4/tcp/routers/<router_name>/entrypoints`"

    See [entry points](../../install-configuration/entrypoints.md) for more information.

    | Key (Path)                                      | Value |
    |-------------------------------------------------|-------|
    | `apache4/tcp/routers/mytcprouter/entrypoints/0` | `ep1` |
    | `apache4/tcp/routers/mytcprouter/entrypoints/1` | `ep2` |
    
??? info "`apache4/tcp/routers/<router_name>/rule`"

    See [entry points](../../install-configuration/entrypoints.md) for more information.

    | Key (Path)                           | Value                        |
    |--------------------------------------|------------------------------|
    | `apache4/tcp/routers/my-router/rule` | ```HostSNI(`example.com`)``` |  

??? info "`apache4/tcp/routers/<router_name>/service`"

    See [service](../tcp/service.md) for more information.
    
    | Key (Path)                                | Value       |
    |-------------------------------------------|-------------|
    | `apache4/tcp/routers/mytcprouter/service` | `myservice` |

??? info "`apache4/tcp/routers/<router_name>/tls`"

    See [TLS](../tcp/tls.md) for more information.

    | Key (Path)                            | Value  |
    |---------------------------------------|--------|
    | `apache4/tcp/routers/mytcprouter/tls` | `true` |

??? info "`apache4/tcp/routers/<router_name>/tls/certresolver`"

    See [certResolver](../tcp/tls.md#configuration-options) for more information.

    | Key (Path)                                         | Value        |
    |----------------------------------------------------|--------------|
    | `apache4/tcp/routers/mytcprouter/tls/certresolver` | `myresolver` |

??? info "`apache4/tcp/routers/<router_name>/tls/domains/<n>/main`"

    See [TLS](../tcp/tls.md) for more information.

    | Key (Path)                                           | Value         |
    |------------------------------------------------------|---------------|
    | `apache4/tcp/routers/mytcprouter/tls/domains/0/main` | `example.org` |
        
??? info "`apache4/tcp/routers/<router_name>/tls/domains/<n>/sans`"

    See [TLS](../tcp/tls.md) for more information.

    | Key (Path)                                             | Value              |
    |--------------------------------------------------------|--------------------|
    | `apache4/tcp/routers/mytcprouter/tls/domains/0/sans/0` | `test.example.org` |
    | `apache4/tcp/routers/mytcprouter/tls/domains/0/sans/1` | `dev.example.org`  |
    
??? info "`apache4/tcp/routers/<router_name>/tls/options`"

    See [TLS](../tcp/tls.md) for more information.

    | Key (Path)                                    | Value    |
    |-----------------------------------------------|----------|
    | `apache4/tcp/routers/mytcprouter/tls/options` | `foobar` |

??? info "`apache4/tcp/routers/<router_name>/tls/passthrough`"

    See [TLS](../tcp/tls.md) for more information.

    | Key (Path)                                        | Value  |
    |---------------------------------------------------|--------|
    | `apache4/tcp/routers/mytcprouter/tls/passthrough` | `true` |

??? info "`apache4/tcp/routers/<router_name>/priority`"

    See [priority](../tcp/router/rules-and-priority.md#priority) for more information.

    | Key (Path)                               | Value |
    |------------------------------------------|-------|
    | `apache4/tcp/routers/mytcprouter/priority`  | `42`  |

#### TCP Services

??? info "`apache4/tcp/services/<service_name>/loadbalancer/servers/<n>/address`"

    See [servers](../tcp/service.md#servers-load-balancer) for more information.

    | Key (Path)                                                         | Value            |
    |--------------------------------------------------------------------|------------------|
    | `apache4/tcp/services/mytcpservice/loadbalancer/servers/0/address` | `xx.xx.xx.xx:xx` |
    
??? info "`apache4/tcp/services/<service_name>/loadbalancer/servers/<n>/tls`"

    See [servers](../tcp/service.md#servers-load-balancer) for more information.

    | Key (Path)                                                         | Value            |
    |--------------------------------------------------------------------|------------------|
    | `apache4/tcp/services/mytcpservice/loadbalancer/servers/0/tls` | `true` |

??? info "`apache4/tcp/services/<service_name>/loadbalancer/proxyprotocol/version`"

    See [PROXY protocol](../tcp/service.md#proxy-protocol) for more information.

    | Key (Path)                                                             | Value |
    |------------------------------------------------------------------------|-------|
    | `apache4/tcp/services/mytcpservice/loadbalancer/proxyprotocol/version` | `1`   |

??? info "`apache4/tcp/services/<service_name>/loadbalancer/serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../tcp/serverstransport.md) for more information.

    | Key (Path)                                                      | Value         |
    |-----------------------------------------------------------------|---------------|
    | `apache4/tcp/services/myservice/loadbalancer/serverstransport` | `foobar@file` |

??? info "`apache4/tcp/services/<service_name>/weighted/services/<n>/name`"

    | Key (Path)                                                          | Value    |
    |---------------------------------------------------------------------|----------|
    | `apache4/tcp/services/<service_name>/weighted/services/0/name`      | `foobar` |

??? info "`apache4/tcp/services/<service_name>/weighted/services/<n>/weight`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `apache4/tcp/services/<service_name>/weighted/services/0/weight` | `42`  |

#### TCP Middleware

You can declare pieces of middleware using tags starting with `apache4/tcp/middlewares/{name-of-your-choice}.`, followed by the middleware type/options.

For example, to declare a middleware [`InFlightConn`](../tcp/middlewares/inflightconn.md) named `test-inflightconn`, you'd write `apache4/tcp/middlewares/test-inflightconn/inflightconn/amount=10`.

More information about available middlewares in the dedicated [middlewares section](../tcp/middlewares/overview.md).

??? example "Declaring and Referencing a Middleware"
    
    ```bash
    # ...
    # Declaring a middleware
    apache4/tcp/middlewares/test-inflightconn/amount=10
    # Referencing a middleware
    apache4/tcp/routers.my-service/middlewares=test-inflightconn
    ```

!!! warning "Conflicts in Declaration"

    If you declare multiple middleware with the same name but with different parameters, the middleware fails to be declared.

### UDP

You can declare UDP Routers and/or Services using KV.

#### UDP Routers

??? info "`apache4/udp/routers/<router-name>/entrypoints/<n>`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `apache4/udp/routers/myudprouter/entrypoints/0` | `foobar`  |

??? info "`apache4/udp/routers/<router-name>/service`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `apache4/udp/routers/myudprouter/service` | `foobar`  |

#### UDP Services

??? info "`apache4/udp/services/loadBalancer/servers/<n>/address`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `apache4/udp/services/loadBalancer/servers/<n>/address` | `foobar`  |

??? info "`apache4/udp/services/weighted/services/<n>/name`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `apache4/udp/services/weighted/services/0/name` | `foobar`  |

??? info "`apache4/udp/services/weighted/services/<n>/name`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `apache4/udp/services/weighted/servers/0/weight` | `42`  |

## TLS

### TLS Options

With the KV provider, you configure some parameters of the TLS connection using the `tls/options` key. For example, you can define a basic setup like this:

| Key (Path)                                           | Value    |
|------------------------------------------------------|----------|
| `apache4/tls/options/Options0/alpnProtocols/0`       | `foobar` |
| `apache4/tls/options/Options0/cipherSuites/0`        | `foobar` |
| `apache4/tls/options/Options0/clientAuth/caFiles/0`  | `foobar` |
| `apache4/tls/options/Options0/disableSessiontickets` | `true`   |

For more information on the available TLS options that can be configured, please refer to the [TLS Options](../http/tls/tls-options.md) page.

### TLS Default Generated Certificates

You can configure apache4 to use an ACME provider (like Let's Encrypt) to generate the default certificate. The configuration to resolve the default certificate should be defined in a TLS store:

| Key (Path)                                                     | Value    |
|----------------------------------------------------------------|----------|
| `apache4/tls/stores/Store0/defaultGeneratedCert/domain/main`   | `foobar` |
| `apache4/tls/stores/Store0/defaultGeneratedCert/domain/sans/0` | `foobar` |
| `apache4/tls/stores/Store0/defaultGeneratedCert/domain/sans/1` | `foobar` |
| `apache4/tls/stores/Store0/defaultGeneratedCert/resolver`      | `foobar` |
