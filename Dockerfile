FROM golang:1.19.2-alpine

COPY main.go /main.go
WORKDIR /
RUN go mod init  hpa-example
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/collectors
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
RUN go build /main.go
EXPOSE 8000:8000
ENTRYPOINT go run /main.go
