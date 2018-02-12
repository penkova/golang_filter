set -x
set -e
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main
docker build -t mongo-filter .
docker-compose up