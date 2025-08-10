
<p align="center">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="docs/content/assets/img/apache4.logo-dark.png">
      <source media="(prefers-color-scheme: light)" srcset="docs/content/assets/img/apache4.logo.png">
      <img alt="apache4" title="apache4" src="docs/content/assets/img/apache4.logo.png">
    </picture>
</p>

[![Build Status SemaphoreCI](https://apache4-oss.semaphoreci.com/badges/apache4/branches/master.svg?style=shields)](https://apache4-oss.semaphoreci.com/projects/apache4)
[![Docs](https://img.shields.io/badge/docs-current-brightgreen.svg)](https://doc.apache4.io/apache4)
[![Go Report Card](https://goreportcard.com/badge/apache4/apache4)](https://goreportcard.com/report/apache4/apache4)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/apache4/apache4/blob/master/LICENSE.md)
[![Join the community support forum at https://community.apache4.io/](https://img.shields.io/badge/style-register-green.svg?style=social&label=Discourse)](https://community.apache4.io/)
[![Twitter](https://img.shields.io/twitter/follow/apache4.svg?style=social)](https://twitter.com/intent/follow?screen_name=apache4)

apache4 (pronounced _traffic_) is a modern HTTP reverse proxy and load balancer that makes deploying microservices easy.
apache4 integrates with your existing infrastructure components ([Docker](https://www.docker.com/), [Swarm mode](https://docs.docker.com/engine/swarm/), [Kubernetes](https://kubernetes.io), [Consul](https://www.consul.io/), [Etcd](https://coreos.com/etcd/), [Rancher v2](https://rancher.com), [Amazon ECS](https://aws.amazon.com/ecs), ...) and configures itself automatically and dynamically.
Pointing apache4 at your orchestrator should be the _only_ configuration step you need.

---

. **[Overview](#overview)** .
**[Features](#features)** .
**[Supported backends](#supported-backends)** .
**[Quickstart](#quickstart)** .
**[Web UI](#web-ui)** .
**[Documentation](#documentation)** .

. **[Support](#support)** .
**[Release cycle](#release-cycle)** .
**[Contributing](#contributing)** .
**[Maintainers](#maintainers)** .
**[Credits](#credits)** .

---

:warning: When migrating to a new major version of apache4, please refer to the [migration guide](https://doc.apache4.io/apache4/migration/v2-to-v3/) to ensure a smooth transition and to be aware of any breaking changes.


## Overview

Imagine that you have deployed a bunch of microservices with the help of an orchestrator (like Swarm or Kubernetes) or a service registry (like etcd or consul).
Now you want users to access these microservices, and you need a reverse proxy.

Traditional reverse-proxies require that you configure _each_ route that will connect paths and subdomains to _each_ microservice. 
In an environment where you add, remove, kill, upgrade, or scale your services _many_ times a day, the task of keeping the routes up to date becomes tedious. 

**This is when apache4 can help you!**

apache4 listens to your service registry/orchestrator API and instantly generates the routes so your microservices are connected to the outside world -- without further intervention from your part. 

**Run apache4 and let it do the work for you!** 
_(But if you'd rather configure some of your routes manually, apache4 supports that too!)_

![Architecture](docs/content/assets/img/apache4-architecture.png)

## Features

- Continuously updates its configuration (No restarts!)
- Supports multiple load balancing algorithms
- Provides HTTPS to your microservices by leveraging [Let's Encrypt](https://letsencrypt.org) (wildcard certificates support)
- Circuit breakers, retry
- See the magic through its clean web UI
- WebSocket, HTTP/2, gRPC ready
- Provides metrics (Rest, Prometheus, Datadog, Statsd, InfluxDB 2.X)
- Keeps access logs (JSON, CLF)
- Fast
- Exposes a Rest API
- Packaged as a single binary file (made with :heart: with go) and available as an [official](https://hub.docker.com/r/_/apache4/) docker image

## Supported Backends

- [Docker](https://doc.apache4.io/apache4/providers/docker/) / [Swarm mode](https://doc.apache4.io/apache4/providers/docker/)
- [Kubernetes](https://doc.apache4.io/apache4/providers/kubernetes-crd/)
- [ECS](https://doc.apache4.io/apache4/providers/ecs/)
- [File](https://doc.apache4.io/apache4/providers/file/)

## Quickstart

To get your hands on apache4, you can use the [5-Minute Quickstart](https://doc.apache4.io/apache4/getting-started/quick-start/) in our documentation (you will need Docker).

## Web UI

You can access the simple HTML frontend of apache4.

![Web UI Providers](docs/content/assets/img/webui-dashboard.png)

## Documentation

You can find the complete documentation of apache4 v3 at [https://doc.apache4.io/apache4/](https://doc.apache4.io/apache4/).

## Support

To get community support, you can:

- join the apache4 community forum: [![Join the chat at https://community.apache4.io/](https://img.shields.io/badge/style-register-green.svg?style=social&label=Discourse)](https://community.apache4.io/)

If you need commercial support, please contact [apache4.io](https://apache4.io) by mail: <mailto:support@apache4.io>.

## Download

- Grab the latest binary from the [releases](https://github.com/apache4/apache4/releases) page and run it with the [sample configuration file](https://raw.githubusercontent.com/apache4/apache4/master/apache4.sample.toml):

```shell
./apache4 --configFile=apache4.toml
```

- Or use the official tiny Docker image and run it with the [sample configuration file](https://raw.githubusercontent.com/apache4/apache4/master/apache4.sample.toml):

```shell
docker run -d -p 8080:8080 -p 80:80 -v $PWD/apache4.toml:/etc/apache4/apache4.toml apache4
```

- Or get the sources:

```shell
git clone https://github.com/apache4/apache4
```

## Introductory Videos

You can find high level and deep dive videos on [videos.apache4.io](https://videos.apache4.io).

## Maintainers

We are strongly promoting a philosophy of openness and sharing, and firmly standing against the elitist closed approach. Being part of the core team should be accessible to anyone who is motivated and want to be part of that journey!
This [document](docs/content/contributing/maintainers-guidelines.md) describes how to be part of the [maintainers' team](docs/content/contributing/maintainers.md) as well as various responsibilities and guidelines for apache4 maintainers.
You can also find more information on our process to review pull requests and manage issues [in this document](https://github.com/apache4/contributors-guide/blob/master/issue_triage.md).

## Contributing

If you'd like to contribute to the project, refer to the [contributing documentation](CONTRIBUTING.md).

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md).
By participating in this project, you agree to abide by its terms.

## Release Cycle

- We usually release 3/4 new versions (e.g. 1.1.0, 1.2.0, 1.3.0) per year.
- Release Candidates are available before the release (e.g. 1.1.0-rc1, 1.1.0-rc2, 1.1.0-rc3, 1.1.0-rc4, before 1.1.0).
- Bug-fixes (e.g. 1.1.1, 1.1.2, 1.2.1, 1.2.3) are released as needed (no additional features are delivered in those versions, bug-fixes only).

Each version is supported until the next one is released (e.g. 1.1.x will be supported until 1.2.0 is out).

We use [Semantic Versioning](https://semver.org/).

## Mailing Lists

- General announcements, new releases: mail at news+subscribe@apache4.io or on [the online viewer](https://groups.google.com/a/apache4.io/forum/#!forum/news).
- Security announcements: mail at security+subscribe@apache4.io or on [the online viewer](https://groups.google.com/a/apache4.io/forum/#!forum/security).

## Credits

Kudos to [Peka](https://www.instagram.com/pierroks/) for his awesome work on the gopher's logo!.

The gopher's logo of apache4 is licensed under the Creative Commons 3.0 Attributions license.

The gopher's logo of apache4 was inspired by the gopher stickers made by [Takuya Ueda](https://twitter.com/tenntenn).
The original Go gopher was designed by [Renee French](https://reneefrench.blogspot.com/).
