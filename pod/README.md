## Pod
A sample output showing the metadata, spec, and status (status of the containers inside as well) of all pods on GKE.

## Sample output
```
---------------------Pod 26---------------------

### Object metadata ###
Node name: prometheus-to-sd-xw5xc
Namespace: kube-system
UID: 637600b0-7284-44b4-8354-2223e84e6cfe
Creation time: 2020-11-25 09:11:34 +0800 HKT

### Object metadata - extra ###
Labels: [controller-revision-hash: 7c676f4db4] [k8s-app: prometheus-to-sd] [pod-template-generation: 1] 
Owner reference: 
[Name: prometheus-to-sd]
[UID: 54c97bfc-2d9b-4e32-9deb-2df5db414e37]
[Kind: DaemonSet]


### Pod spec ###
---Containers in the pod---
Container name: prometheus-to-sd
Container image: k8s.gcr.io/prometheus-to-sd:v0.8.2
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: ClusterFirst
Restart policy: Always
Node name: gke-cluster-1-default-pool-c6af14e8-9fj1

### Pod status ###
Host IP: 10.128.0.4
Pod IP: 10.128.0.4
Phase: Running
Container status: 
--- Containers 1 ---
Container name: prometheus-to-sd
Container Image: k8s.gcr.io/prometheus-to-sd:v0.8.2
Ready: true
Container restart count: 0

---------------------Pod 27---------------------

### Object metadata ###
Node name: stackdriver-metadata-agent-cluster-level-5c4f6c65c6-l7pq7
Namespace: kube-system
UID: aed97291-ec3f-41a5-a4fa-44c6f3482d1b
Creation time: 2020-11-25 09:11:59 +0800 HKT

### Object metadata - extra ###
Labels: [app: stackdriver-metadata-agent] [cluster-level: true] [pod-template-hash: 5c4f6c65c6] 
Owner reference: 
[Name: stackdriver-metadata-agent-cluster-level-5c4f6c65c6]
[UID: 28eeb66a-f3ff-4a7c-9895-583f186cb3be]
[Kind: ReplicaSet]


### Pod spec ###
---Containers in the pod---
Container name: metadata-agent
Container image: gcr.io/stackdriver-agents/metadata-agent-go:1.2.0
Container pull policy: IfNotPresent
Container ports: 
---------------
Container name: metadata-agent-nanny
Container image: gke.gcr.io/addon-resizer:1.8.11-gke.1
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: ClusterFirst
Restart policy: Always
Node name: gke-cluster-1-default-pool-c6af14e8-9fj1

### Pod status ###
Host IP: 10.128.0.4
Pod IP: 10.4.2.4
Phase: Running
Container status: 
--- Containers 1 ---
Container name: metadata-agent
Container Image: gcr.io/stackdriver-agents/metadata-agent-go:1.2.0
Ready: true
Container restart count: 0
--- Containers 2 ---
Container name: metadata-agent-nanny
Container Image: gke.gcr.io/addon-resizer:1.8.11-gke.1
Ready: true
Container restart count: 0
```