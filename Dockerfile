# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine

WORKDIR /app

COPY main/go.mod ./
COPY main/go.sum ./
RUN go mod download && go mod verify

COPY main/*.go ./
COPY .env ./

RUN go build -o /eververse

CMD [ "/eververse"]
