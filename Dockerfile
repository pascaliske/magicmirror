# server
FROM --platform=${BUILDPLATFORM} golang:alpine as server
LABEL maintainer="info@pascaliske.dev"
WORKDIR /go/src/app

# environment
ARG TARGETOS
ARG TARGETARCH
ENV CGO_ENABLED=0

# install dependencies
COPY server/go.mod /go/src/app
COPY server/go.sum /go/src/app
RUN go mod download

# build binary
COPY server /go/src/app
RUN go build -v -ldflags="-w -s" -o ./magicmirror main.go

# app
FROM --platform=${BUILDPLATFORM} node:16-alpine AS app
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
FROM alpine
LABEL maintainer="info@pascaliske.dev"

# environment
ENV MM_PORT=9000

# install tini
RUN apk add --no-cache curl tini

# inject built files
COPY --from=server /go/src/app/magicmirror /magicmirror
COPY --from=app /build/dist/magicmirror /public

HEALTHCHECK CMD curl -f http://localhost:${MM_PORT}/health || exit 1

ENTRYPOINT [ "/sbin/tini", "--" ]
CMD [ "/magicmirror" ]
