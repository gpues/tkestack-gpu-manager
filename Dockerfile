FROM nvidia/cuda:12.2.2-devel-centos7 as build
WORKDIR /gpu
ENV GOLANG_VERSION 1.21.0
RUN curl -sSL https://dl.google.com/go/go${GOLANG_VERSION}.linux-amd64.tar.gz \
    | tar -C /usr/local -xz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOFLAGS="-buildvcs=false"
RUN yum install git -y
COPY . .

RUN go mod tidy
RUN go build -o /tmp/gpu-manager ./cmd/manager
RUN go build -o /tmp/gpu-client ./cmd/client

FROM registry.cn-hangzhou.aliyuncs.com/acejilam/centos:7

COPY --from=build /tmp/gpu-manager /usr/bin/
COPY --from=build /tmp/gpu-client /usr/bin/


# kubelet
VOLUME ["/var/lib/kubelet/device-plugins"]

# gpu manager storage
VOLUME ["/etc/gpu-manager/vm"]
VOLUME ["/etc/gpu-manager/vdriver"]
VOLUME ["/var/log/gpu-manager"]

# nvidia library search location
VOLUME ["/usr/local/host"]

RUN echo "/usr/local/nvidia/lib" > /etc/ld.so.conf.d/nvidia.conf && \
    echo "/usr/local/nvidia/lib64" >> /etc/ld.so.conf.d/nvidia.conf

ENV PATH=$PATH:/usr/local/nvidia/bin

# cgroup
VOLUME ["/sys/fs/cgroup"]

# display
EXPOSE 5678

COPY build/start.sh /
COPY build/copy-bin-lib.sh /

CMD ["/start.sh"]
