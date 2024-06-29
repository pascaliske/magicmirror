# --- tini
FROM --platform=${BUILDPLATFORM} alpine:3.20 as tini
LABEL maintainer="info@pascaliske.dev"

# environment
ENV TINI_VERSION=v0.19.0
ARG TARGETPLATFORM

# install tini
RUN case ${TARGETPLATFORM} in \
    "linux/amd64")  TINI_ARCH=amd64  ;; \
    "linux/arm64")  TINI_ARCH=arm64  ;; \
    "linux/arm/v7") TINI_ARCH=armhf  ;; \
    esac \
    && wget -q https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static-${TINI_ARCH} -O /tini \
    && chmod +x /tini

# --- api
FROM --platform=${BUILDPLATFORM} golang:1.22-alpine as api
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

# build binary
COPY api /build
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -ldflags="-w -s -X ${MM_PKG}.Version=${MM_VERSION} -X ${MM_PKG}.GitCommit=${MM_GIT_COMMIT} -X ${MM_PKG}.BuildTime=${MM_BUILD_TIME}" -o ./magicmirror main.go

# --- app
FROM --platform=${BUILDPLATFORM} node:20-alpine AS app
LABEL maintainer="info@pascaliske.dev"
WORKDIR /build

# install dependencies
COPY package.json /build
COPY yarn.lock /build
RUN yarn install --frozen-lockfile --ignore-scripts

# build app
COPY . /build
RUN yarn run build

# final image
FROM alpine:3.20
LABEL maintainer="info@pascaliske.dev"

# create non-root user
RUN addgroup -S -g 911 unknown && adduser -S -u 911 -G unknown -s /bin/false unknown

# environment
ENV MM_PORT=9000

# install curl
RUN apk add --no-cache curl shadow su-exec

# inject built files
COPY --from=tini /tini /sbin/tini
COPY --from=api /build/magicmirror /magicmirror
COPY --from=app /build/dist/magicmirror/browser /public

# inject entrypoint
COPY docker-entrypoint.sh /docker-entrypoint.sh

# health check
HEALTHCHECK CMD curl -f http://localhost:${MM_PORT}/health || exit 1

# let's go!
ENTRYPOINT [ "/sbin/tini", "--", "/docker-entrypoint.sh" ]
CMD [ "/magicmirror" ]
