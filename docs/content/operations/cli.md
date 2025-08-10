---
title: "apache4 CLI Documentation"
description: "Learn the basics of the apache4 Proxy command line interface (CLI). Read the technical documentation."
---

# CLI

The apache4 Command Line
{: .subtitle }

## General

```bash
apache4 [command] [flags] [arguments]
```

Use `apache4 [command] --help` for help on any command.

Commands:

- `healthcheck` Calls apache4 `/ping` to check the health of apache4 (the API must be enabled).
- `version` Shows the current apache4 version.

Flag's usage:

```bash
# set flag_argument to flag(s)
apache4 [--flag=flag_argument] [-f [flag_argument]]

# set true/false to boolean flag(s)
apache4 [--flag[=true|false| ]] [-f [true|false| ]]
```

All flags are documented in the [(static configuration) CLI reference](../reference/static-configuration/cli.md).

!!! info "Flags are case-insensitive."

### `healthcheck`

Calls apache4 `/ping` to check the health of apache4.
Its exit status is `0` if apache4 is healthy and `1` otherwise.

This can be used with Docker [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) instruction
or any other health check orchestration mechanism.

!!! info
    The [`ping` endpoint](../operations/ping.md) must be enabled to allow the `healthcheck` command to call `/ping`.

Usage:

```bash
apache4 healthcheck [command] [flags] [arguments]
```

Example:

```bash
$ apache4 healthcheck
OK: http://:8082/ping
```

### `version`

Shows the current apache4 version.

Usage:

```bash
apache4 version
```
