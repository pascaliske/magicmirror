# Configuration

## Introduction

Both, the server and app components can be fully customized using a single YAML configuration file but none of the keys are actually required. The file can be placed in one of the following destinations inside the container where it will be automatically picked up:

- `/etc/magicmirror/config.yml` (preferred)
- `/config.yml`

Alternatively you could explicitly pass in the configuration file via flag:

```shell
$ /magicmirror --config /foo/bar/baz.yml
```

??? tldr "Full example with all possible keys"

    ```yaml title="/etc/magicmirror/config.yml"
    # optional: http listener port for the server
    port: 9000

    # optional: configuration for logging, possible values: 'debug', 'info', 'warn', 'error', 'fatal'
    log:
      level: debug

    # optional: configuration for prometheus metrics endpoint
    metrics:
      enabled: true
      path: /metrics

    # optional: geo location, only required for various cards and apis, e.g. weather card
    location:
      latitude: 1.2345678
      longitude: 1.234567

    # optional: rss feeds displayed inside news card
    feeds:
      - https://my-news.org/rss/feed

    # optional: api keys for the respective data apis
    apiKeys:
      # optional: open weather api key, required if weather card is used
      openWeather: my-api-key-123
    ```

## Check-only mode

You can run the server in a check-only mode to validate the configuration file. Simply pass the `--check` flag like so:

```shell
# default config path
$ /magicmirror --check

# custom config path
$ /magicmirror --config /foo/bar/baz.yml --check
```

## Sections

To keep things simple the configuration file is divided into multiple sections. Almost all of them can be omitted as they are optional.

### Basic

| Key    | Type  | Mandatory | Default | Description                                   |
| ------ | ----- | --------- | ------- | --------------------------------------------- |
| `port` | `int` | no        | `9000`  | Port for the HTTP server to serve the web UI. |

??? example

    ```yaml
    port: 9000
    ```

### Log

| Key     | Type                                               | Mandatory | Default | Description                         |
| ------- | -------------------------------------------------- | --------- | ------- | ----------------------------------- |
| `level` | `enum` (`debug`, `info`, `warn`, `error`, `fatal`) | no        | `info`  | Log level for the server component. |

??? example

    ```yaml
    log:
      level: debug
    ```

### Metrics

The server is capable of exposing some basic metrics for [Prometheus](https://prometheus.io/). To enable them the following configuration values can be used:

| Key       | Type     | Mandatory | Default    | Description                                                                             |
| --------- | -------- | --------- | ---------- | --------------------------------------------------------------------------------------- |
| `enabled` | `bool`   | no        | `true`     | Enable/Disable Prometheus metrics endpoint.                                             |
| `path`    | `string` | no        | `/metrics` | Path of the metrics endpoint. The default value is ready for Prometheus to be consumed. |

??? example

    ```yaml
    metrics:
      enabled: true
      path: /metrics
    ```

### Location

Customize your geo location. It is needed for various cards, such as the weather card.

> Note: If you skip this section those cards will be disabled automatically.

| Key         | Type    | Mandatory | Default | Description                               |
| ----------- | ------- | --------- | ------- | ----------------------------------------- |
| `latitude`  | `float` | no        | -       | Latitude as float for your geo location.  |
| `longitude` | `float` | no        | -       | Longitude as float for your geo location. |

??? example

    ```yaml
    location:
      latitude: 1.2345678
      longitude: 1.234567
    ```

### Feeds

The news card requires RSS feeds to display them on the mirror. This section enables you to pass in one or more URLs inside a top-level list configuration item.

??? example

    ```yaml
    feeds:
      - https://my-news.org/rss/feed
    ```

### API Keys

Some cards (e.g. the weather card) need an API key to authenticate you against their data APIs.

> Note: If you skip this section those cards will be disabled automatically.

| Key           | Type     | Mandatory | Default | Description                                        |
| ------------- | -------- | --------- | ------- | -------------------------------------------------- |
| `openWeather` | `string` | no        | -       | OpenWeather API key required for the weather card. |

??? example

    ```yaml
    apiKeys:
      openWeather: my-api-key-123
    ```
