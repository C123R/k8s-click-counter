FROM golang:alpine

WORKDIR /go/src

COPY src .
WORKDIR /go/src/app

RUN go install app

EXPOSE 9000

ENTRYPOINT /go/bin/app
