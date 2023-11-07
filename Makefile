ROOT=${shell pwd }
all:
	go build -o "${ROOT}/bin/gpu-manager" -ldflags "-s -w" ./cmd/manager
	go build -o "${ROOT}/bin/gpu-client" -ldflags "-s -w" ./cmd/client

clean:
	rm -rf ./bin

test:
	go fmt ./...
	go test -timeout=1m -bench=. -cover -v ./...
proto:
	go get github.com/grpc-ecosystem/grpc-gateway@v1.16.0
	protoc --proto_path=staging/src:. --proto_path="${shell go env GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis":. --go_out=. --go-grpc_out=./ --grpc-gateway_out=logtostderr=true:. pkg/api/runtime/display/api.proto
	protoc --proto_path=staging/src:. --proto_path="${shell go env GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis":. --go_out=. --go-grpc_out=./ pkg/api/runtime/vcuda/api.proto


img:
	docker build --pull --no-cache . -t k8s:v1
	docker run --security-opt seccomp=unconfined --rm -it -v /etc/gpu-manager/vm:/etc/gpu-manager/vm k8s:v1 bash

lint:
#  https://linksaas.pro/download
#  @revive -config revive.toml -exclude vendor/... -exclude pkg/api/runtime/... ./...
