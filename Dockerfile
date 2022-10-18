FROM golang:1.19.2-alpine

COPY main.go /main.go
ENTRYPOINT go run /main.go
