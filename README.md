# hpa-example

Docker image for starting a simple web server in Go to test Kubernetes Horizontal Pod Autoscaler.

we are adapting:
https://github.com/prometheus/client_golang

for this

### build/test  notes
```
go mod init  hpa-example
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/collectors
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
go mod tidy
go build ./main.go
go run ./main.go
# from another terminal:
curl "http://127.0.0.1:8080/metrics"
```
### publish to quay
```
docker build -t hpa-example .
export MYTAG="v0.3.0"
docker tag hpa-example:${MYTAG} quay.io/paperlesspost/hpa-example:${MYTAG}
docker login -u="XXX" -p="XXX quay.io
docker tag hpa-example:${MYTAG} quay.io/paperlesspost/hpa-example:${MYTAG}
docker push quay.io/paperlesspost/hpa-example:${MYTAG}

```
