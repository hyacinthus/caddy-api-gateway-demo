docker run --rm -v ${PWD}:/go/src/app -w /go/src/app golang:alpine go build -v
docker build -f Dockerfile.app -t muninn/caddy-microservice:app .
docker build -f Dockerfile.gateway -t muninn/caddy-microservice:gateway .
