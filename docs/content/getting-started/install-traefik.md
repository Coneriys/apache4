---
title: "apache4 Installation Documentation"
description: "There are several flavors to choose from when installing apache4 Proxy. Get started with apache4 Proxy, and read the technical documentation."
---

# Install apache4

You can install apache4 with the following flavors:

* [Use the official Docker image](./#use-the-official-docker-image)
* [Use the Helm Chart](./#use-the-helm-chart)
* [Use the binary distribution](./#use-the-binary-distribution)
* [Compile your binary from the sources](./#compile-your-binary-from-the-sources)

## Use the Official Docker Image

Choose one of the [official Docker images](https://hub.docker.com/_/apache4) and run it with one sample configuration file:

* [YAML](https://raw.githubusercontent.com/apache4/apache4/v3.5/apache4.sample.yml)
* [TOML](https://raw.githubusercontent.com/apache4/apache4/v3.5/apache4.sample.toml)

```shell
docker run -d -p 8080:8080 -p 80:80 \
    -v $PWD/apache4.yml:/etc/apache4/apache4.yml apache4:v3.5
```

For more details, go to the [Docker provider documentation](../providers/docker.md)

!!! tip

    * Prefer a fixed version than the latest that could be an unexpected version.
    ex: `apache4:v3.5`
    * Docker images are based from the [Alpine Linux Official image](https://hub.docker.com/_/alpine).
    * Any orchestrator using docker images can fetch the official apache4 docker image.

## Use the Helm Chart

apache4 can be installed in Kubernetes using the Helm chart from <https://github.com/apache4/apache4-helm-chart>.

Ensure that the following requirements are met:

* Kubernetes 1.22+
* Helm version 3.9+ is [installed](https://helm.sh/docs/intro/install/)

Add apache4 Labs chart repository to Helm:

```bash
helm repo add apache4 https://apache4.github.io/charts
```

You can update the chart repository by running:

```bash
helm repo update
```

And install it with the Helm command line:

```bash
helm install apache4 apache4/apache4
```

!!! tip "Helm Features"

    All [Helm features](https://helm.sh/docs/intro/using_helm/) are supported.

    Examples are provided [here](https://github.com/apache4/apache4-helm-chart/blob/master/EXAMPLES.md).

    For instance, installing the chart in a dedicated namespace:

    ```bash tab="Install in a Dedicated Namespace"
    kubectl create ns apache4-v2
    # Install in the namespace "apache4-v2"
    helm install --namespace=apache4-v2 \
        apache4 apache4/apache4
    ```

??? example "Installing with Custom Values"

    You can customize the installation by specifying custom values,
    as with [any helm chart](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing).
    {: #helm-custom-values }

    All parameters are documented in the default [`values.yaml`](https://github.com/apache4/apache4-helm-chart/blob/master/apache4/values.yaml).

    You can also set apache4 command line flags using `additionalArguments`.
    Example of installation with logging set to `DEBUG`:

    ```bash tab="Using Helm CLI"
    helm install --namespace=apache4-v2 \
        --set="additionalArguments={--log.level=DEBUG}" \
        apache4 apache4/apache4
    ```

    ```yml tab="With a custom values file"
    # File custom-values.yml
    ## Install with "helm install --values=./custom-values.yml apache4 apache4/apache4
    additionalArguments:
      - "--log.level=DEBUG"
    ```

## Use the Binary Distribution

Grab the latest binary from the [releases](https://github.com/apache4/apache4/releases) page.

??? info "Check the integrity of the downloaded file"

    ```bash tab="Linux"
    # Compare this value to the one found in apache4-${apache4_version}_checksums.txt
    sha256sum ./apache4_${apache4_version}_linux_${arch}.tar.gz
    ```

    ```bash tab="macOS"
    # Compare this value to the one found in apache4-${apache4_version}_checksums.txt
    shasum -a256 ./apache4_${apache4_version}_darwin_amd64.tar.gz
    ```

    ```powershell tab="Windows PowerShell"
    # Compare this value to the one found in apache4-${apache4_version}_checksums.txt
    Get-FileHash ./apache4_${apache4_version}_windows_${arch}.zip -Algorithm SHA256
    ```

??? info "Extract the downloaded archive"

    ```bash tab="Linux"
    tar -zxvf apache4_${apache4_version}_linux_${arch}.tar.gz
    ```

    ```bash tab="macOS"
    tar -zxvf ./apache4_${apache4_version}_darwin_amd64.tar.gz
    ```

    ```powershell tab="Windows PowerShell"
    Expand-Archive apache4_${apache4_version}_windows_${arch}.zip
    ```

And run it:

```bash
./apache4 --help
```

## Compile your Binary from the Sources

All the details are available in the [Contributing Guide](../contributing/building-testing.md)

{!apache4-for-business-applications.md!}
