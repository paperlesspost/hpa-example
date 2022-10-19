FROM golang:1 AS builder
WORKDIR ${GOPATH}/pkg/mod/github.com/prometheus/client_golang
COPY . .
WORKDIR ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/prometheus
RUN go install paperlesspost.net/hpa-example@latest
WORKDIR ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/examples/random
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
WORKDIR ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/examples/simple
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
WORKDIR ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/examples/gocollector
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

# Final image.
FROM quay.io/prometheus/busybox:latest
LABEL maintainer="The Prometheus Authors <prometheus-developers@googlegroups.com>"
COPY --from=builder ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/examples/random \
    ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/examples/simple \
    ${GOPATH}/pkg/mod/github.com/prometheus/client_golang@v1.13.0/examples/gocollector ./
EXPOSE 8080
CMD ["echo", "Please run an example. Either /random, /simple or /gocollector"]
