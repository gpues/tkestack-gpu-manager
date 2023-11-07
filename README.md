## Pod template example

There is nothing special to submit a Pod except the description of GPU resource is no longer 1
. The GPU
resources are described as that 100 `tencent.com/vcuda-core` for 1 GPU and N `tencent.com/vcuda-memory` for GPU memory (1 tencent.com/vcuda-memory means 256Mi
GPU memory). And because of the limitation of extend resource validation of Kubernetes, to support
GPU utilization limitation, you should add `tencent.com/vcuda-core-limit: XX` in the annotation
 field of a Pod.
 
 **Notice: the value of `tencent.com/vcuda-core` is either the multiple of 100 or any value
smaller than 100.For example, 100, 200 or 20 is valid value but 150 or 250 is invalid**

- Submit a Pod with 0.3 GPU utilization and 7680MiB GPU memory with 0.5 GPU utilization limit

```
apiVersion: v1
kind: Pod
metadata:
  name: vcuda
  annotations:
    tencent.com/vcuda-core-limit: 50
spec:
  restartPolicy: Never
  containers:
  - image: <test-image>
    name: nvidia
    command:
    - /usr/local/nvidia/bin/nvidia-smi
    - pmon
    - -d
    - 10
    resources:
      requests:
        tencent.com/vcuda-core: 50
        tencent.com/vcuda-memory: 30
      limits:
        tencent.com/vcuda-core: 50
        tencent.com/vcuda-memory: 30
```

- Submit a Pod with 2 GPU card

```
apiVersion: v1
kind: Pod
metadata:
  name: vcuda
spec:
  restartPolicy: Never
  containers:
  - image: <test-image>
    name: nvidia
    command:
    - /usr/local/nvidia/bin/nvidia-smi
    - pmon
    - -d
    - 10
    resources:
      requests:
        tencent.com/vcuda-core: 200
        tencent.com/vcuda-memory: 60
      limits:
        tencent.com/vcuda-core: 200
        tencent.com/vcuda-memory: 60
```

## FAQ

If you have some questions about this project, you can first refer to [FAQ](./docs/faq.md) to find a solution.
