# Configuration

## Introduction

The Magic Mirror can be fully customized using a YAML configuration file. The file can be placed in one of the following destinations where it will be automatically picked up by the server:

- `/etc/magicmirror/config.yml` (preferred)
- `/config.yml`

Alternatively you could explicitly pass in the configuration file via flag:

```shell
$ /magicmirror --config /foo/bar/baz.yml
```

## Sections

To keep things simple the configuration file is divided into multiple sections. Many of them can be omitted as they are optional.

### Basic

| Key  | Type  | Mandatory | Default | Description                                   |
| ---- | ----- | --------- | ------- | --------------------------------------------- |
| port | `int` | no        | `9000`  | Port for the HTTP server to serve the web UI. |

??? example

    ```yaml
    port: 9000
    ```

### Log

| Key   | Type                                               | Mandatory | Default | Description                         |
| ----- | -------------------------------------------------- | --------- | ------- | ----------------------------------- |
| level | `enum` (`debug`, `info`, `warn`, `error`, `fatal`) | no        | `info`  | Log level for the server component. |

??? example

    ```yaml
    log:
      level: debug
    ```

### Metrics

The server is capable of exposing some basic metrics for [Prometheus](https://prometheus.io/). To enable them the following configuration values can be used:

| Key     | Type     | Mandatory | Default    | Description                                                                             |
| ------- | -------- | --------- | ---------- | --------------------------------------------------------------------------------------- |
| enabled | `bool`   | no        | `true`     | Enable/Disable Prometheus metrics endpoint.                                             |
| path    | `string` | no        | `/metrics` | Path of the metrics endpoint. The default value is ready for Prometheus to be consumed. |

??? example

    ```yaml
    metrics:
      enabled: true
      path: /metrics
    ```

### Location

Customize your geo location. It is needed for various cards, such as the weather card.

> Note: If you skip this section those cards will be disabled automatically.

| Key       | Type    | Mandatory | Default | Description                               |
| --------- | ------- | --------- | ------- | ----------------------------------------- |
| latitude  | `float` | no        | -       | Latitude as float for your geo location.  |
| longitude | `float` | no        | -       | Longitude as float for your geo location. |

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

| Key         | Type     | Mandatory | Default | Description                                        |
| ----------- | -------- | --------- | ------- | -------------------------------------------------- |
| openWeather | `string` | no        | -       | OpenWeather API key required for the weather card. |

??? example

    ```yaml
    apiKeys:
      openWeather: my-api-key-123
    ```
