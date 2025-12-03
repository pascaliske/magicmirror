# --- tini
FROM --platform=${BUILDPLATFORM} alpine:3.23 AS tini
LABEL maintainer="info@pascaliske.dev"

# environment
ENV TINI_VERSION=v0.19.0
ARG TARGETPLATFORM

# install tini
RUN case "${TARGETPLATFORM}" in \
    "linux/amd64")  TINI_ARCH=amd64  ;; \
    "linux/arm64")  TINI_ARCH=arm64  ;; \
    "linux/arm/v7") TINI_ARCH=armhf  ;; \
    esac \
    && wget -q https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static-${TINI_ARCH} -O /tini \
    && chmod +x /tini

# --- app
FROM --platform=${BUILDPLATFORM} node:22-alpine AS app
LABEL maintainer="info@pascaliske.dev"
WORKDIR /build

# install dependencies
COPY package.json /build
COPY yarn.lock /build
RUN yarn install --frozen-lockfile --ignore-scripts

# inject source code
COPY . /build

# build app
RUN yarn run build

# --- api
FROM --platform=${BUILDPLATFORM} golang:1.25-alpine AS api
LABEL maintainer="info@pascaliske.dev"
WORKDIR /build

# environment
ARG TARGETOS
ARG TARGETARCH
ENV CGO_ENABLED=0
ARG MM_PKG="github.com/pascaliske/magicmirror/version"
ARG MM_VERSION
ARG MM_GIT_COMMIT
ARG MM_BUILD_TIME

# install dependencies
COPY api/go.mod /build
COPY api/go.sum /build
RUN go mod download

# inject source code
COPY api /build

# inject built app
COPY --from=app /build/dist/magicmirror/browser /build/public/static

# build binary
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -ldflags="-w -s -X ${MM_PKG}.Version=${MM_VERSION} -X ${MM_PKG}.GitCommit=${MM_GIT_COMMIT} -X ${MM_PKG}.BuildTime=${MM_BUILD_TIME}" -o ./magicmirror main.go

# --- final image
FROM scratch
LABEL maintainer="info@pascaliske.dev"

# inject binaries
COPY --from=tini /tini /sbin/tini
COPY --from=api /build/magicmirror /magicmirror

# health check
HEALTHCHECK CMD ["/magicmirror", "health"]

# let's go!
ENTRYPOINT [ "/sbin/tini", "--", "/magicmirror" ]
CMD [ "serve" ]
