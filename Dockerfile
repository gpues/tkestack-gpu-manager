FROM registry.cn-hangzhou.aliyuncs.com/acejilam/nvidia-cuda:12.2.2-devel-centos7 as build
WORKDIR /gpu
RUN yum install git -y
ENV GOLANG_VERSION 1.21.0
RUN curl -sSL https://dl.google.com/go/go${GOLANG_VERSION}.linux-$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/).tar.gz \
    | tar -C /usr/local -xz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY go.mod ./
COPY go.sum ./
COPY ./staging ./staging/
COPY ./cmd ./cmd/
COPY ./pkg ./pkg/

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOFLAGS="-buildvcs=false"
RUN go mod tidy
RUN go build -o /tmp/gpu-manager ./cmd/manager
RUN go build -o /tmp/gpu-client ./cmd/client
RUN go build -o /tmp/copy-bin-lib ./cmd/copy-bin-lib
RUN go build -o /tmp/nvml ./cmd/nvml
COPY ./build/* /tmp
RUN mv /tmp/jq-linux-$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) /tmp/jq && chmod +x /tmp/jq
COPY build/libvgpu.so /tmp/
RUN chmod 777 /tmp/libvgpu.so && chown root:root /tmp/libvgpu.so

#FROM registry.cn-hangzhou.aliyuncs.com/acejilam/centos-cuda:7.11.08
FROM registry.cn-hangzhou.aliyuncs.com/acejilam/centos:7

# kubelet
VOLUME ["/var/lib/kubelet/device-plugins"]

# gpu manager storage
VOLUME ["/etc/gpu-manager/vm"]

# /usr:/usr/local/gpu/host
# 容器(/usr/local/gpu拷贝到/usr/local/nvidia) -> 宿主机(/etc/gpu-manager/vdriver:/usr/local/nvidia)
VOLUME ["/etc/gpu-manager/vdriver"]

# /etc/gpu-manager/log:/var/log/gpu-manager
VOLUME ["/var/log/gpu-manager"]

# nvidia library search location
VOLUME ["/usr/local/gpu/"]

#RUN echo "/usr/local/nvidia/lib" > /etc/ld.so.conf.d/nvidia.conf && \
#    echo "/usr/local/nvidia/lib64" >> /etc/ld.so.conf.d/nvidia.conf

RUN echo "/etc/gpu-manager/vdriver/nvidia/lib" > /etc/ld.so.conf.d/nvidia.conf && \
    echo "/etc/gpu-manager/vdriver/nvidia/lib64" >> /etc/ld.so.conf.d/nvidia.conf

# cgroup
VOLUME ["/sys/fs/cgroup"]

# display
EXPOSE 5678

COPY build/start.sh /
COPY build/copy-bin-lib.sh /


COPY build/volume.json /etc/gpu-manager/
COPY build/extra-config.json /etc/gpu-manager/

COPY --from=build /tmp/gpu-manager /usr/bin/
COPY --from=build /tmp/gpu-client /usr/bin/
COPY --from=build /tmp/copy-bin-lib /usr/bin/
COPY --from=build /tmp/jq /usr/bin/jq
COPY --from=build /tmp/libvgpu.so /usr/local/gpu/
COPY --from=build /tmp/nvml /usr/bin/nvml
CMD ["/start.sh"]
