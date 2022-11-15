FROM golang:1.17-alpine3.13 AS build-env

LABEL maintainer="adi khoiron hasan<adikhoironhasan@gmail.com>"

ENV APP_NAME=privy-cake-store

RUN apk update && apk upgrade && apk add git

RUN ls -ls

RUN mkdir -p /src/privy-cake-store

COPY . /src/privy-cake-store

WORKDIR /src/privy-cake-store

RUN go get all

RUN go mod tidy

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build

EXPOSE 8000

CMD "./privy-cake-store"