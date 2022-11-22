FROM golang:1.19.2-alpine

COPY main.go /main.go
WORKDIR /
RUN go build /main.go
EXPOSE 8000:8000
ENTRYPOINT go run /main.go
