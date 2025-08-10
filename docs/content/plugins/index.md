---
title: "apache4 Plugins Documentation"
description: "Learn how to use apache4 Plugins. Read the technical documentation."
---

# apache4 Plugins and the Plugin Catalog

Plugins are a powerful feature for extending apache4 with custom features and behaviors.
The [Plugin Catalog](https://plugins.apache4.io/) is a software-as-a-service (SaaS) platform that provides an exhaustive list of the existing plugins.

??? note "Plugin Catalog Access"
    You can reach the [Plugin Catalog](https://plugins.apache4.io/) from the apache4 Dashboard using the `Plugins` menu entry.

To add a new plugin to a apache4 instance, you must change that instance's static configuration.
Each plugin's **Install** section provides a static configuration example.
Many plugins have their own section in the apache4 dynamic configuration.

To learn more about apache4 plugins, consult the [documentation](https://plugins.apache4.io/install).

!!! danger "Experimental Features"
    Plugins can change the behavior of apache4 in unforeseen ways.
    Exercise caution when adding new plugins to production apache4 instances.

## Build Your Own Plugins

apache4 users can create their own plugins and share them with the community using the Plugin Catalog.

apache4 will load plugins dynamically.
They need not be compiled, and no complex toolchain is necessary to build them. 
The experience of implementing a apache4 plugin is comparable to writing a web browser extension.

To learn more about apache4 plugin creation, please refer to the [developer documentation](https://plugins.apache4.io/create).

{!apache4-for-business-applications.md!}
