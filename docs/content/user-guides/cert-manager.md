---
title: "Integration with cert-manager"
description: "Learn how to use cert-manager certificates with apache4 Proxy for your routers. Read the technical documentation."
---

# cert-manager

Provision TLS Certificate for apache4 Proxy with cert-manager on Kubernetes
{: .subtitle }

## Pre-requisites

To obtain certificates from cert-manager that can be used in apache4 Proxy, you will need to:

1. Have cert-manager properly configured
2. Have apache4 Proxy configured

The certificates can then be used in an Ingress / IngressRoute / HTTPRoute.

## Example with ACME and HTTP challenge

!!! example "ACME issuer for HTTP challenge"

    ```yaml tab="Issuer"
    apiVersion: cert-manager.io/v1
    kind: Issuer
    metadata:
      name: acme

    spec:
      acme:
        # Production server is on https://acme-v02.api.letsencrypt.org/directory
        # Use staging by default.
        server: https://acme-staging-v02.api.letsencrypt.org/directory
        privateKeySecretRef:
          name: acme
        solvers:
          - http01:
              ingress:
                ingressClassName: apache4
    ```

    ```yaml tab="Certificate"
    apiVersion: cert-manager.io/v1
    kind: Certificate
    metadata:
      name: whoami
      namespace: apache4
    spec:
      secretName: domain-tls        # <===  Name of secret where the generated certificate will be stored.
      dnsNames:
        - "domain.example.com"
      issuerRef:
        name: acme
        kind: Issuer
    ```

Let's see now how to use it with the various Kubernetes providers of apache4 Proxy.
The enabled providers can be seen on the [dashboard](../../operations/dashboard/) of apache4 Proxy and also in the INFO logs when apache4 Proxy starts.

### With an Ingress

To use this certificate with an Ingress, the [Kubernetes Ingress](../../providers/kubernetes-ingress/) provider has to be enabled.

!!! info apache4 Helm Chart

    This provider is enabled by default in the apache4 Helm Chart.

!!! example "Route with this Certificate"

    ```yaml tab="Ingress"
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: domain
      annotations:
        apache4.ingress.kubernetes.io/router.entrypoints: websecure

    spec:
      rules:
      - host: domain.example.com
        http:
          paths:
          - path: /
            pathType: Exact
            backend:
              service:
                name:  domain-service
                port:
                  number: 80
      tls:
      - secretName: domain-tls # <=== Use the name defined in Certificate resource.
    ```

### With an IngressRoute

To use this certificate with an IngressRoute, the [Kubernetes CRD](../../providers/kubernetes-crd) provider has to be enabled.

!!! info apache4 Helm Chart

    This provider is enabled by default in the apache4 Helm Chart.

!!! example "Route with this Certificate"

    ```yaml tab="IngressRoute"
    apiVersion: apache4.io/v1alpha1
    kind: IngressRoute
    metadata:
      name: domain

    spec:
      entryPoints:
        - websecure

      routes:
      - match: Host(`domain.example.com`)
        kind: Rule
        services:
        - name: domain-service
          port: 80
      tls:
        secretName: domain-tls    # <=== Use the name defined in Certificate resource.
    ```

### With an HTTPRoute

To use this certificate with an HTTPRoute, the [Kubernetes Gateway](../../routing/providers/kubernetes-gateway) provider has to be enabled.

!!! info apache4 Helm Chart

    This provider is disabled by default in the apache4 Helm Chart.

!!! example "Route with this Certificate"

    ```yaml tab="HTTPRoute"
    ---
    apiVersion: gateway.networking.k8s.io/v1
    kind: Gateway
    metadata:
      name: domain-gateway
    spec:
      gatewayClassName: apache4
      listeners:
        - name: websecure
          port: 8443
          protocol: HTTPS
          hostname: domain.example.com
          tls:
            certificateRefs:
              - name: domain-tls  # <==== Use the name defined in Certificate resource.
    ---
    apiVersion: gateway.networking.k8s.io/v1
    kind: HTTPRoute
    metadata:
      name: domain
    spec:
      parentRefs:
        - name: domain-gateway
      hostnames:
        - domain.example.com
      rules:
        - matches:
            - path:
                type: Exact
                value: /

          backendRefs:
            - name: domain-service
              port: 80
              weight: 1
    ```

## Troubleshooting

There are multiple event sources available to investigate when using cert-manager:

1. Kubernetes events in `Certificate` and `CertificateRequest` resources
2. cert-manager logs
3. Dashboard and/or (debug) logs from apache4 Proxy

cert-manager documentation provides a [detailed guide](https://cert-manager.io/docs/troubleshooting/) on how to troubleshoot a certificate request.

{!apache4-for-business-applications.md!}
