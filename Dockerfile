FROM golang:1.19.2-alpine

COPY main.go /main.go
EXPOSE 8080
ENTRYPOINT go run /main.go
