docker run --rm -v ${PWD}:/go/src/client -w /go/src/client golang:alpine go build -v
docker build -f Dockerfile.client -t muninn/caddy-microservice:client .
