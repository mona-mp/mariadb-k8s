# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /usr/local/go/src/user-management

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go /usr/local/go/src/user-management/
COPY controllers /usr/local/go/src/user-management/controllers
COPY database /usr/local/go/src/user-management/database
COPY entity /usr/local/go/src/user-management/entity
COPY .env /usr/local/go/src/user-management/

RUN env GOOS=linux GOARCH=amd64 go build main.go
RUN chmod +x  main

EXPOSE 18080

CMD [ "./main" ]