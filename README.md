# [`pascaliske/magicmirror`](https://pascaliske.github.io/magicmirror/)

> Go + Angular based smart mirror platform - packaged as single docker image.

[![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/pascaliske/magicmirror/latest?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![Docker Image Size (tag)](https://img.shields.io/docker/image-size/pascaliske/magicmirror/latest?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![Docker Pulls](https://img.shields.io/docker/pulls/pascaliske/magicmirror?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![GitHub Tag](https://img.shields.io/github/v/tag/pascaliske/magicmirror?style=flat-square)](https://github.com/pascaliske/magicmirror) [![Linting Status](https://img.shields.io/github/actions/workflow/status/pascaliske/magicmirror/linting.yml?branch=master&label=linting&style=flat-square)](https://github.com/pascaliske/magicmirror/actions/workflows/linting.yml) [![Build Status](https://img.shields.io/github/actions/workflow/status/pascaliske/magicmirror/image.yml?branch=master&label=build&style=flat-square)](https://github.com/pascaliske/magicmirror/actions/workflows/image.yml) [![GitHub Last Commit](https://img.shields.io/github/last-commit/pascaliske/magicmirror?style=flat-square)](https://github.com/pascaliske/magicmirror) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/MIT) [![Awesome Badges](https://img.shields.io/badge/badges-awesome-green.svg?style=flat-square)](https://github.com/Naereen/badges)

## Images

| Registry   | Image                                                                                                                             |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------- |
| GitHub     | [`ghcr.io/pascaliske/magicmirror`](https://github.com/pascaliske/docker-magicmirror/pkgs/container/magicmirror)                   |
| Docker Hub | [`pascaliske/magicmirror`](https://hub.docker.com/r/pascaliske/magicmirror)                                                       |

The following platforms are available for the images:

```bash
$ docker run --rm mplatform/mquery ghcr.io/pascaliske/magicmirror:latest
Image: pascaliske/magicmirror:latest
 * Manifest List: Yes
 * Supported platforms:
   - linux/amd64
   - linux/arm/v7
   - linux/arm64
```

## Usage

Visit the [documentation](https://pascaliske.github.io/magicmirror/) for a complete setup guide.

To run a quick demo, you can just pull the image from one of the registries:

```bash
# github container registry
$ docker pull ghcr.io/pascaliske/magicmirror

# docker hub
$ docker pull pascaliske/magicmirror
```

Now you can simply run the server component using the following command:

```bash
$ docker run \
    --detach \
    --name magicmirror \
    --publish 9000:9000 \
    --volume $(pwd)/config.yml:/etc/magicmirror/config.yml \
    ghcr.io/pascaliske/magicmirror
```

## Configuration

The project is configurable using a YAML file. You can find all possible values on the [configuration page](https://pascaliske.github.io/magicmirror/configuration/).

## License

[MIT](LICENSE.md) – © 2023 [Pascal Iske](https://pascaliske.dev)
