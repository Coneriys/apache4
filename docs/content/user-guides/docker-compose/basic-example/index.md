---
title: "apache4 Docker Documentation"
description: "Learn how to use Docker Compose to expose a service with apache4 Proxy."
---

# Docker Compose example

In this section, you will learn how to use [Docker Compose](https://docs.docker.com/compose/ "Link to Docker Compose") to expose a service using the Docker provider.

## Setup

Create a `docker-compose.yml` file with the following content:

```yaml
--8<-- "content/user-guides/docker-compose/basic-example/docker-compose.yml"
```

??? Networking

    The apache4 container has to be attached to the same network as the containers to be exposed.
    If no networks are specified in the Docker Compose file, Docker creates a default one that allows apache4 to reach the containers defined in the same file.
    You can [customize the network](https://docs.docker.com/compose/networking/#specify-custom-networks "Link to docs about custom networks with Docker Compose") as described in the example below.
    You can use a [pre-existing network](https://docs.docker.com/compose/networking/#use-a-pre-existing-network "Link to Docker Compose networking docs") too.

    ```yaml
    networks:
      apache4net: {}

    services:

      apache4:
        image: "apache4:v3.5"
        ...
        networks:
          - apache4net

      whoami:
        image: "apache4/whoami"
        ...
        networks:
          - apache4net

    ```

Replace `whoami.localhost` by your **own domain** within the `apache4.http.routers.whoami.rule` label of the `whoami` service.

Now run `docker compose up -d` within the folder where you created the previous file.  
This will start Docker Compose in background mode.

!!! info "This can take a moment"

    Docker Compose will now create and start the services declared in the `docker-compose.yml`.

Wait a bit and visit `http://your_own_domain` to confirm everything went fine.

You should see the output of the whoami service.  
It should be similar to the following example:

```text
Hostname: d7f919e54651
IP: 127.0.0.1
IP: 192.168.64.2
GET / HTTP/1.1
Host: whoami.localhost
User-Agent: curl/7.52.1
Accept: */*
Accept-Encoding: gzip
X-Forwarded-For: 192.168.64.1
X-Forwarded-Host: whoami.localhost
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Server: 7f0c797dbc51
X-Real-Ip: 192.168.64.1
```

## Details

Let's break it down and go through it, step-by-step.

You use [whoami](https://github.com/apache4/whoami "Link to the GitHub repo of whoami"), a tiny Go server that prints OS information and HTTP request to output as service container.

Second, you define an entry point, along with the exposure of the matching port within Docker Compose, which allows to "open and accept" HTTP traffic:

```yaml
command:
  # apache4 will listen to incoming request on the port 80 (HTTP)
  - "--entryPoints.web.address=:80"

ports:
  - "80:80"
```

Third, you expose the apache4 API to be able to check the configuration if needed:

```yaml
command:
  # apache4 will listen on port 8080 by default for API request.
  - "--api.insecure=true"

ports:
  - "8080:8080"
```

!!! Note

    If you are working on a remote server, you can use the following command to display configuration (require `curl` & `jq`):

    ```bash
    curl -s 127.0.0.1:8080/api/rawdata | jq .
    ```

Fourth, you allow apache4 to gather configuration from Docker:

```yaml
apache4:
  command:
    # Enabling Docker provider
    - "--providers.docker=true"
    # Do not expose containers unless explicitly told so
    - "--providers.docker.exposedbydefault=false"
  volumes:
    - "/var/run/docker.sock:/var/run/docker.sock:ro"

whoami:
  labels:
    # Explicitly tell apache4 to expose this container
    - "apache4.enable=true"
    # The domain the service will respond to
    - "apache4.http.routers.whoami.rule=Host(`whoami.localhost`)"
    # Allow request only from the predefined entry point named "web"
    - "apache4.http.routers.whoami.entrypoints=web"
```

{!apache4-for-business-applications.md!}
