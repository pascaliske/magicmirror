# `pascaliske/magicmirror`

> Go + Angular based smart mirror platform - packaged as single docker image.

[![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/pascaliske/magicmirror/latest?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![Docker Image Size (tag)](https://img.shields.io/docker/image-size/pascaliske/magicmirror/latest?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![Docker Pulls](https://img.shields.io/docker/pulls/pascaliske/magicmirror?style=flat-square)](https://hub.docker.com/r/pascaliske/magicmirror) [![GitHub Tag](https://img.shields.io/github/v/tag/pascaliske/magicmirror?style=flat-square)](https://github.com/pascaliske/magicmirror) [![Build Status](https://img.shields.io/github/workflow/status/pascaliske/magicmirror/Image/master?label=build&style=flat-square)](https://github.com/pascaliske/magicmirror/actions) [![GitHub Last Commit](https://img.shields.io/github/last-commit/pascaliske/magicmirror?style=flat-square)](https://github.com/pascaliske/magicmirror) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/MIT) [![Awesome Badges](https://img.shields.io/badge/badges-awesome-green.svg?style=flat-square)](https://github.com/Naereen/badges)

## Usage

To use this image pull it from one of the following registries:

```bash
# docker hub
docker pull pascaliske/magicmirror

# github container registry
docker pull ghcr.io/pascaliske/magicmirror
```

To run the server use the following command:

```bash
docker run \
    --detach \
    --name magicmirror
    --publish 9000:9000 \
    --volume $(pwd)/config.yml:/config.yml \
    ghcr.io/pascaliske/magicmirror
```

## License

[MIT](LICENSE.md) – © 2022 [Pascal Iske](https://pascaliske.dev)
