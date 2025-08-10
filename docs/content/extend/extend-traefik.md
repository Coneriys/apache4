---
title: Extend apache4
description: Extend apache4 with custom plugins using Yaegi and WebAssembly.
---

# Extend apache4

Plugins are a powerful feature for extending apache4 with custom features and behaviors. The [Plugin Catalog](https://plugins.apache4.io/) is a software-as-a-service (SaaS) platform that provides an exhaustive list of the existing plugins.

??? note "Plugin Catalog Access"
    You can reach the [Plugin Catalog](https://plugins.apache4.io/) from the apache4 Dashboard using the `Plugins` menu entry.

## Add a new plugin to a apache4 instance

To add a new plugin to a apache4 instance, you must change that instance's install (static) configuration. Each plugin's **Install** section provides an install (static) configuration example. Many plugins have their own section in the apache4 routing (dynamic) configuration.

!!! danger "Experimental Features"
    Plugins can change the behavior of apache4 in unforeseen ways. Exercise caution when adding new plugins to production apache4 instances.

To learn more about how to add a new plugin to a apache4 instance, please refer to the [developer documentation](https://plugins.apache4.io/install).

## Plugin Systems

apache4 supports two different plugin systems, each designed for different use cases and developer preferences.

### Yaegi Plugin System

apache4 [Yaegi](https://github.com/apache4/yaegi) plugins are developed using the Go language. It is essentially a Go package. Unlike pre-compiled plugins, Yaegi plugins are executed on the fly by Yaegi, a Go interpreter embedded in apache4.

This approach eliminates the need for compilation and a complex toolchain, making plugin development as straightforward as creating web browser extensions. Yaegi plugins support both middleware and provider functionality.

#### Key characteristics

- Written in Go language
- No compilation required
- Executed by embedded interpreter
- Supports full Go feature set
- Hot-reloadable during development

### WebAssembly (WASM) Plugin System

apache4 WASM plugins can be developed using any language that compiles to WebAssembly (WASM). This method is based on [http-wasm](https://http-wasm.io/).

WASM plugins compile to portable binary modules that execute with near-native performance while maintaining security isolation.

#### Key characteristics

- Multi-language support (Go, Rust, C++, etc.)
- Compiled to WebAssembly binary
- Near-native performance
- Strong security isolation
- Currently supports middleware only

## Build Your Own Plugins

apache4 users can create their own plugins and share them with the community using the [Plugin Catalog](https://plugins.apache4.io/). To learn more about apache4 plugin creation, please refer to the [developer documentation](https://plugins.apache4.io/create).
