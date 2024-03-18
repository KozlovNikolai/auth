LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate-grpc:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux_auth cmd/grpc-server/main.go

copy-to-server:
	scp service_linux_auth root@5.159.101.123:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/msc/auth:v0.0.1 .
	docker login -u token -p CRgAAAAAD6l6BhyNZkQouCVYue4xorwBot6D5eZ6 cr.selcloud.ru/msc
	docker push cr.selcloud.ru/msc/auth:v0.0.1

run-into-server:
	docker pull cr.selcloud.ru/msc/auth:v0.0.1
	docker run -p 50061:50061 cr.selcloud.ru/msc/auth:v0.0.1

docker-local-build-and-run:
	-docker stop auth
	-docker rm auth
	-docker rmi auth:v0.0.1
	docker buildx build --no-cache --platform linux/amd64 -t auth:v0.0.1 .
	docker run --name auth -d -p 50061:50051 auth:v0.0.1

