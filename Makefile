compose:
	 docker-compose build && docker-compose up --remove-orphans

build:
	 docker-compose build

up:
	 docker-compose up

test:
	echo "testing.."
	go clean -testcache
	go test -v -cover ./...

proto:
	protoc --go_out=. --go-grpc_out=. api/rpc/*.proto

.DEFAULT_GOAL := compose