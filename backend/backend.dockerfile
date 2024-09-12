FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN env GOOS=linux CGO_ENABLED=0 go build -o backendApp .

RUN chmod +x /app/backendApp

## build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/backendApp /app

CMD ["/app/backendApp"]