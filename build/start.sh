#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

source copy-bin-lib.sh

echo "rebuild ldcache"
/usr/sbin/ldconfig

echo "launch gpu manager"
/usr/bin/gpu-manager \
--v=${LOG_LEVEL} \
--hostname-override=${NODE_NAME} \
--share-mode=true \
--volume-config=/etc/gpu-manager/volume.json \
--extra-config=/etc/gpu-manager/extra-config.json \
--extra-config=/etc/gpu-manager/extra-config.json \
--log-dir=/var/log/gpu-manager \
--query-addr=0.0.0.0 \
--addr=/var/run/gpu-manager.sock \
--logtostderr



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