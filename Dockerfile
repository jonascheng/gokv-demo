FROM golang:1.16.6-alpine3.14 as build-env

ENV GOOS linux

RUN apk add --no-cache git build-base && \
  echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
  echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
  echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories

WORKDIR /go/src/github.com/jonascheng/gokv-demo/
COPY . .

RUN make build

FROM alpine:3.14

WORKDIR /app

COPY --from=build-env /go/src/github.com/jonascheng/gokv-demo/bin/gokv-demo .

ENTRYPOINT [ "./gokv-demo" ]