FROM golang:1.17-alpine as builder
WORKDIR /root/go/src/github.com/einsier/ustc_melee_game
COPY . /root/go/src/github.com/einsier/ustc_melee_game
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go build -o db-proxy main.go

FROM alpine:latest
# environment variable for mongoDB connection
ARG DB_USER
ARG DB_PWD
ARG DB_HOST
ARG DB_PORT
ENV ENV_DB_USER=$DB_USER \
    ENV_DB_PWD=$DB_PWD \
    ENV_DB_HOST=$DB_HOST \
    ENV_DB_PORT=$DB_PORT
WORKDIR  /root/go/src/github.com/einsier/ustc_melee_game
COPY --from=builder  /root/go/src/github.com/einsier/ustc_melee_game/db-proxy .
EXPOSE 1234/tcp
ENTRYPOINT ./db-proxy -DBUser $ENV_DB_USER -DBPassword $ENV_DB_PWD -DBHost $ENV_DB_HOST -DBPort $ENV_DB_PORT