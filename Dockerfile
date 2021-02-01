FROM golang:alpine
WORKDIR /go/src/github.com/golang-encurtador-url
COPY . .
RUN go build
ENTRYPOINT [ "./golang-encurtador-url" ]
EXPOSE 8888