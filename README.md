# `pascaliske/magicmirror`

> Go + Angular based smart mirror platform - packaged as single docker image.

[![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/pascaliske/magicmirror/latest?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![Docker Image Size (tag)](https://img.shields.io/docker/image-size/pascaliske/magicmirror/latest?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![Docker Pulls](https://img.shields.io/docker/pulls/pascaliske/magicmirror?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![GitHub Tag](https://img.shields.io/github/v/tag/pascaliske/magicmirror?style=flat-square)](https://github.com/pascaliske/magicmirror) [![Linting Status](https://img.shields.io/github/workflow/status/pascaliske/magicmirror/Linting/master?label=linting&style=flat-square)](https://github.com/pascaliske/magicmirror/actions/workflows/linting.yml) [![Build Status](https://img.shields.io/github/workflow/status/pascaliske/magicmirror/Image/master?label=build&style=flat-square)](https://github.com/pascaliske/magicmirror/actions/workflows/image.yml) [![GitHub Last Commit](https://img.shields.io/github/last-commit/pascaliske/magicmirror?style=flat-square)](https://github.com/pascaliske/magicmirror) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/MIT) [![Awesome Badges](https://img.shields.io/badge/badges-awesome-green.svg?style=flat-square)](https://github.com/Naereen/badges)

## Image

| Registry   | Image                                                                                                                             |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------- |
| GitHub     | [`ghcr.io/pascaliske/magicmirror`](https://github.com/pascaliske/docker-magicmirror/pkgs/container/magicmirror)                   |
| Docker Hub | [`pascaliske/magicmirror`](https://hub.docker.com/r/pascaliske/magicmirror)                                                       |

The following platforms are available for this image:

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

To use this image pull it from one of the following registries:

```bash
# github container registry
docker pull ghcr.io/pascaliske/magicmirror

# docker hub
docker pull pascaliske/magicmirror
```

Now you can run the server using the following command:

```bash
docker run \
    --detach \
    --name magicmirror \
    --publish 9000:9000 \
    --volume $(pwd)/config.yml:/config.yml \
    ghcr.io/pascaliske/magicmirror
```

## License

[MIT](LICENSE.md) – © 2022 [Pascal Iske](https://pascaliske.dev)
