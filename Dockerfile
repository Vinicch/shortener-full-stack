FROM golang:1.18.2-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go mod tidy
RUN go build -o shortener main.go

FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/shortener /usr/local/bin/shortener
ENTRYPOINT [ "shortener" ]