FROM golang:alpine
WORKDIR /go/src/github.com/golang-encurtador-url
COPY . .
RUN apk update
RUN apk add --no-cache git 
RUN go get -u github.com/gin-gonic/gin
EXPOSE 8888