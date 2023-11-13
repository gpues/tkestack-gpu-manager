#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset



/usr/sbin/ldconfig

/usr/bin/gpu-manager \
--v=${LOG_LEVEL} \
--hostname-override=${NODE_NAME} \
--share-mode=true \
--extra-config=/etc/gpu-manager/extra-config.json \
--log-dir=/var/log/gpu-manager ${EXTRA_FLAGS:-""}

# 宿主机目录 /usr/ -> 宿主机目录 /usr/local/gpu/host

# -v /sys/fs/cgroup:/sys/fs/cgroup
# -v /usr:/usr/local/host:
# -v /var/run:/var/run
# -v /etc/gpu-manager/checkpoint:/etc/gpu-manager/checkpoint
# -v /etc/gpu-manager/log:/var/log/gpu-manager
# -v /etc/gpu-manager/vdriver:/etc/gpu-manager/vdriver
# -v /etc/gpu-manager/vm:/etc/gpu-manager/vm
# -v /var/lib/kubelet/device-plugins:/var/lib/kubelet/device-plugins

# container
# -v /etc/gpu-manager/vdriver/nvidia:/usr/local/nvidia/lib64
# -v /etc/gpu-manager/vm/containerID/:/etc/vcuda
#    /etc/gpu-manager/vm/containerID/vcuda.sock