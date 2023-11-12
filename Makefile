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
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	protoc --proto_path=staging/src:. --proto_path="${shell go env GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis":. --go_out=. --go-grpc_out=./ --grpc-gateway_out=logtostderr=true:. pkg/api/runtime/display/api.proto
	protoc --proto_path=staging/src:. --proto_path="${shell go env GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis":. --go_out=. --go-grpc_out=./ pkg/api/runtime/vcuda/api.proto

run:
	docker build --pull --no-cache . -t k8s:v1
	docker run --security-opt seccomp=unconfined --rm -it -v /etc/gpu-manager/vm:/etc/gpu-manager/vm k8s:v1 bash

#lint:
#  https://linksaas.pro/download
#  @revive -config revive.toml -exclude vendor/... -exclude pkg/api/runtime/... ./...

multi_arch:
	docker buildx build -f Dockerfile --no-cache=true --pull=false --platform=linux/amd64,linux/arm64 -t registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5 . --push


push:
	docker build -f Dockerfile --no-cache -t registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5 .
	docker push registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5
#	go build -o ./cmd/nvml/nvml ./cmd/nvml
#
#	docker run --name=cen --rm -it -v /usr/:/usr/local/gpu/host/ -v /etc/gpu-manager/vdriver:/etc/gpu-manager/vdriver registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5 /usr/bin/copy-bin-lib
#
#	docker run --privileged --name=cen --rm -it -v `pwd`:/data -v /etc/gpu-manager/vdriver/nvidia/:/usr/local/nvidia/ -e LD_LIBRARY_PATH=/usr/local/nvidia/lib64/  registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5 /data/cmd/nvml/nvml
#	docker run --privileged --name=cen --rm -it -v `pwd`:/data -v /etc/gpu-manager/vdriver/nvidia/:/etc/gpu-manager/vdriver/nvidia/ -e LD_LIBRARY_PATH=/etc/gpu-manager/vdriver/nvidia/lib64  registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5 /usr/bin/nvml
#	docker run --privileged --name=cen --rm -it -v `pwd`:/data -v /etc/gpu-manager/vdriver/nvidia/:/usr/local/nvidia/  registry.cn-hangzhou.aliyuncs.com/acejilam/tkestack-gpu-manager:v1.1.5 bash
