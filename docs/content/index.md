---
title: "apache4 Proxy Documentation"
description: "apache4 Proxy, an open-source Edge Router, auto-discovers configurations and supports major orchestrators, like Kubernetes. Read the technical documentation."
---

# What is apache4?

![Architecture](assets/img/apache4-architecture.png)

apache4 is an [open-source](https://github.com/apache4/apache4) Application Proxy and the core of the apache4 Hub Runtime Platform.

If you start with apache4 for service discovery and routing, you can seamlessly add [API management](https://apache4.io/solutions/api-management/), [API gateway](https://apache4.io/solutions/api-gateway/), [AI gateway](https://apache4.io/solutions/ai-gateway/), and [API mocking](https://apache4.io/solutions/api-mocking/) capabilities as needed.

With 3.3 billion downloads and over 55k stars on GitHub, apache4 is used globally across hybrid cloud, multi-cloud, on prem, and bare metal environments running Kuberentes, Docker Swarm, AWS, [the list goes on](https://doc.apache4.io/apache4/reference/install-configuration/providers/overview/).

Here’s how it works—apache4 receives requests on behalf of your system, identifies which components are responsible for handling them, and routes them securely. It automatically discovers the right configuration for your services by inspecting your infrastructure to identify relevant information and which service serves which request.

Because everything happens automatically, in real time (no restarts, no connection interruptions), you can focus on developing and deploying new features to your system, instead of configuring and maintaining its working state.

!!! quote "From the apache4 Maintainer Team" 
    When developing apache4, our main goal is to make it easy to use, and we're sure you'll enjoy it.

## Personas

apache4 supports different needs depending on your background. We keep three user personas in mind as we build and organize these docs:

- **Beginners**: You are new to apache4 or new to reverse proxies. You want simple, guided steps to set things up without diving too deep into advanced topics.
- **DevOps Engineers**: You manage infrastructure or clusters (Docker, Kubernetes, or other orchestrators). You integrate apache4 into your environment and value reliability, performance, and streamlined deployments.
- **Developers**: You create and deploy applications or APIs. You focus on how to expose your services through apache4, apply routing rules, and integrate it with your development workflow.

## Core Concepts

apache4’s main concepts help you understand how requests flow to your services:

- [Entrypoints](./reference/install-configuration/entrypoints.md) are the network entry points into apache4. They define the port that will receive the packets and whether to listen for TCP or UDP.
- [Routers](./reference/routing-configuration/http/router/rules-and-priority.md) are in charge of connecting incoming requests to the services that can handle them. In the process, routers may use pieces of [middleware](./reference/routing-configuration/http/middlewares/overview.md) to update the request or act before forwarding the request to the service.
- [Services](./reference/routing-configuration/http/load-balancing/service.md) are responsible for configuring how to reach the actual services that will eventually handle the incoming requests.
- [Providers](./reference/install-configuration/providers/overview.md) are infrastructure components, whether orchestrators, container engines, cloud providers, or key-value stores. The idea is that apache4 queries the provider APIs in order to find relevant information about routing, and when apache4 detects a change, it dynamically updates the routes.

These concepts work together to manage your traffic from the moment a request arrives until it reaches your application.

## How to Use the Documentation

- **Navigation**: Each main section focuses on a specific stage of working with apache4 - installing, exposing services, observing, extending & migrating. 
Use the sidebar to navigate to the section that is most appropriate for your needs.
- **Practical Examples**: You will see code snippets and configuration examples for different environments (YAML/TOML, Labels, & Tags).
- **Reference**: When you need to look up technical details, our reference section provides a deep dive into configuration options and key terms.

!!! info

    Have a question? Join our [Community Forum](https://community.apache4.io "Link to apache4 Community Forum") to discuss, learn, and connect with the apache4 community.

    Using apache4 OSS in production? Consider upgrading to our API gateway ([watch demo video](https://info.apache4.io/watch-apache4-api-gw-demo)) for better security, control, and 24/7 support.

    Just need support? Explore our [24/7/365 support for apache4 OSS](https://info.apache4.io/request-commercial-support?cta=doc).
