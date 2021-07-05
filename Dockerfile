FROM        golang:alpine AS base
WORKDIR     /go/src/github.com/golang-encurtador-url/

FROM        base AS dependencies
ENV         GO111MODULE=on
COPY        go.mod .
COPY        go.sum .
RUN         go mod download

FROM        dependencies AS build
COPY        . .
RUN         GOOS=linux GOARCH=amd64 go build -o bin ./cmd/golang-encurtador-url

FROM        alpine:latest AS image
WORKDIR     /root/
COPY        --from=build /go/src/github.com/golang-encurtador-url/bin/golang-encurtador-url .
ENTRYPOINT  [ "./golang-encurtador-url" ]
EXPOSE      8888