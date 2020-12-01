## Node

Here is a sample output of the details of 3 nodes. These nodes belong to the same nodepool.

This is basically a mirror of the 'Node details' page of the GCP console.

## Output
```
---------------------Node 1---------------------

### Object metadata ###
Node name: gke-cluster-1-default-pool-c6af14e8-0zl9
Namespace: 
UID: 7c094e7b-bda6-4c50-8490-9c74f1ee9add
Creation time: 2020-11-25 09:11:33 +0800 HKT

### Node Spec ###
Pod CIDR: 10.4.0.0/24
Node ID: gce://fyp-gcp-296605/us-central1-c/gke-cluster-1-default-pool-c6af14e8-0zl9

### Node System info ###
Machine ID: 2229178c881ccb7883cedb01d7c9175c
Kernel version: 4.19.112+
OS Image: Container-Optimized OS from Google
Operating system: linux
Container runtime version: docker://19.3.1
Kubelet Version: v1.16.13-gke.401
Kube-proxy Version: v1.16.13-gke.401

### Node resources capacity ###
CPU: 2 CPU
Memory: 4.14 GB
Pod: 110 pods
Ephemeral storage: 101.24 GB
Storage: 0 B

### Node resources allocatable ###
CPU: 1 CPU
Memory: 2.97 GB
Pod: 110 pods
Ephemeral storage: 47.09 GB
Storage: 0 B

### Node conditions ###
Condition 0 
 Condition name: FrequentKubeletRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: kubelet is functioning properly

Condition 1 
 Condition name: FrequentDockerRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: docker is functioning properly

Condition 2 
 Condition name: FrequentContainerdRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: containerd is functioning properly

Condition 3 
 Condition name: FrequentUnregisterNetDevice 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: node is functioning properly

Condition 4 
 Condition name: KernelDeadlock 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: kernel has no deadlock

Condition 5 
 Condition name: ReadonlyFilesystem 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: Filesystem is not read-only

Condition 6 
 Condition name: CorruptDockerOverlay2 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:20 +0800 HKT 
 Last transition: 2020-11-25 09:11:36 +0800 HKT 
 Message: docker overlay2 is functioning properly

Condition 7 
 Condition name: NetworkUnavailable 
 Status: False 
 Last heartbeat: 2020-11-25 09:11:34 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: NodeController create implicit route

Condition 8 
 Condition name: MemoryPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:50 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has sufficient memory available

Condition 9 
 Condition name: DiskPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:50 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has no disk pressure

Condition 10 
 Condition name: PIDPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:50 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has sufficient PID available

Condition 11 
 Condition name: Ready 
 Status: True 
 Last heartbeat: 2020-12-01 21:02:50 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: kubelet is posting ready status. AppArmor enabled


### Node config ###

---------------------Node 2---------------------

### Object metadata ###
Node name: gke-cluster-1-default-pool-c6af14e8-9fj1
Namespace: 
UID: f916d621-5315-4f4a-b722-71f7899f537e
Creation time: 2020-11-25 09:11:33 +0800 HKT

### Node Spec ###
Pod CIDR: 10.4.2.0/24
Node ID: gce://fyp-gcp-296605/us-central1-c/gke-cluster-1-default-pool-c6af14e8-9fj1

### Node System info ###
Machine ID: 33e842ad113db39ef6713764e180aa82
Kernel version: 4.19.112+
OS Image: Container-Optimized OS from Google
Operating system: linux
Container runtime version: docker://19.3.1
Kubelet Version: v1.16.13-gke.401
Kube-proxy Version: v1.16.13-gke.401

### Node resources capacity ###
CPU: 2 CPU
Memory: 4.14 GB
Pod: 110 pods
Ephemeral storage: 101.24 GB
Storage: 0 B

### Node resources allocatable ###
CPU: 1 CPU
Memory: 2.97 GB
Pod: 110 pods
Ephemeral storage: 47.09 GB
Storage: 0 B

### Node conditions ###
Condition 0 
 Condition name: KernelDeadlock 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: kernel has no deadlock

Condition 1 
 Condition name: ReadonlyFilesystem 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: Filesystem is not read-only

Condition 2 
 Condition name: CorruptDockerOverlay2 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: docker overlay2 is functioning properly

Condition 3 
 Condition name: FrequentUnregisterNetDevice 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: node is functioning properly

Condition 4 
 Condition name: FrequentKubeletRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: kubelet is functioning properly

Condition 5 
 Condition name: FrequentDockerRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: docker is functioning properly

Condition 6 
 Condition name: FrequentContainerdRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:03 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: containerd is functioning properly

Condition 7 
 Condition name: NetworkUnavailable 
 Status: False 
 Last heartbeat: 2020-11-25 09:11:34 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: NodeController create implicit route

Condition 8 
 Condition name: MemoryPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:02 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has sufficient memory available

Condition 9 
 Condition name: DiskPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:02 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has no disk pressure

Condition 10 
 Condition name: PIDPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:03:02 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has sufficient PID available

Condition 11 
 Condition name: Ready 
 Status: True 
 Last heartbeat: 2020-12-01 21:03:02 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: kubelet is posting ready status. AppArmor enabled


### Node config ###

---------------------Node 3---------------------

### Object metadata ###
Node name: gke-cluster-1-default-pool-c6af14e8-vcnw
Namespace: 
UID: 207ff7af-2d64-41c8-b0de-96e0873bae58
Creation time: 2020-11-25 09:11:33 +0800 HKT

### Node Spec ###
Pod CIDR: 10.4.1.0/24
Node ID: gce://fyp-gcp-296605/us-central1-c/gke-cluster-1-default-pool-c6af14e8-vcnw

### Node System info ###
Machine ID: e3aec2adaff98274f6bb4f234a14e46a
Kernel version: 4.19.112+
OS Image: Container-Optimized OS from Google
Operating system: linux
Container runtime version: docker://19.3.1
Kubelet Version: v1.16.13-gke.401
Kube-proxy Version: v1.16.13-gke.401

### Node resources capacity ###
CPU: 2 CPU
Memory: 4.14 GB
Pod: 110 pods
Ephemeral storage: 101.24 GB
Storage: 0 B

### Node resources allocatable ###
CPU: 1 CPU
Memory: 2.97 GB
Pod: 110 pods
Ephemeral storage: 47.09 GB
Storage: 0 B

### Node conditions ###
Condition 0 
 Condition name: ReadonlyFilesystem 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: Filesystem is not read-only

Condition 1 
 Condition name: CorruptDockerOverlay2 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: docker overlay2 is functioning properly

Condition 2 
 Condition name: FrequentUnregisterNetDevice 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: node is functioning properly

Condition 3 
 Condition name: FrequentKubeletRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: kubelet is functioning properly

Condition 4 
 Condition name: FrequentDockerRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: docker is functioning properly

Condition 5 
 Condition name: FrequentContainerdRestart 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: containerd is functioning properly

Condition 6 
 Condition name: KernelDeadlock 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:38 +0800 HKT 
 Last transition: 2020-11-25 09:11:37 +0800 HKT 
 Message: kernel has no deadlock

Condition 7 
 Condition name: NetworkUnavailable 
 Status: False 
 Last heartbeat: 2020-11-25 09:11:34 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: NodeController create implicit route

Condition 8 
 Condition name: MemoryPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:57 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has sufficient memory available

Condition 9 
 Condition name: DiskPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:57 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has no disk pressure

Condition 10 
 Condition name: PIDPressure 
 Status: False 
 Last heartbeat: 2020-12-01 21:02:57 +0800 HKT 
 Last transition: 2020-11-25 09:11:33 +0800 HKT 
 Message: kubelet has sufficient PID available

Condition 11 
 Condition name: Ready 
 Status: True 
 Last heartbeat: 2020-12-01 21:02:57 +0800 HKT 
 Last transition: 2020-11-25 09:11:34 +0800 HKT 
 Message: kubelet is posting ready status. AppArmor enabled
 ```