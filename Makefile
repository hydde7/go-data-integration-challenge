#Challenge Makefile

start:
  go run main.go

check:
  go test ./src/database

setup:
  go get -d ./...
  docker-compose up -d
