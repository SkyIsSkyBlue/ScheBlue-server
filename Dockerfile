FROM golang:1.19

ENV GO111MODULE=on
ENV DOCKERIZE_VERSION v0.6.1
RUN go install github.com/jwilder/dockerize@${DOCKERIZE_VERSION}

# setting workdir
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

# wait for sql
ENTRYPOINT dockerize -timeout 60s -wait tcp://mysql:3306 air -c .air.toml