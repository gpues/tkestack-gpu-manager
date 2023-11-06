ROOT=${shell pwd }
all:
	go build -o "${ROOT}/bin/gpu-manager" -ldflags "-s -w" ./cmd/manager
	go build -o "${ROOT}/bin/gpu-client" -ldflags "-s -w" ./cmd/client

clean:
	rm -rf ./bin

test:
	go fmt ./...
	go test -timeout=1m -bench=. -cover -v ./...
#
#proto:
#	  docker run --rm \
#    -v ${ROOT}/pkg/api:/tmp/pkg/api \
#    -v ${ROOT}/staging/src:/tmp/staging/src \
#    -u $(id -u) \
#    devsu/grpc-gateway \
#      bash -c "cd /tmp && protoc \\
#        --proto_path=staging/src:. \\
#        --proto_path=/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis:. \\
#        --go_out=plugins=grpc:. \\
#        --grpc-gateway_out=logtostderr=true:. \\
#        pkg/api/runtime/display/api.proto"
#
#  docker run --rm \
#    -v ${ROOT}/pkg/api:/tmp/pkg/api \
#    -u $(id -u) \
#    devsu/grpc-gateway \
#      bash -c "cd /tmp && protoc \\
#        --go_out=plugins=grpc:. \\
#        pkg/api/runtime/vcuda/api.proto"
#

img:
	docker build --pull --no-cache . -t k8s:v1
	docker run --security-opt seccomp=unconfined --rm -it -v /etc/gpu-manager/vm:/etc/gpu-manager/vm k8s:v1 bash




lint:
#  https://linksaas.pro/download
#  @revive -config revive.toml -exclude vendor/... -exclude pkg/api/runtime/... ./...
