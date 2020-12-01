## Deployment
An example of printing the metadata, spec, and status of all deployments.

Can look at how to filter out system object later.

## Sample output
```
-------------Deployment 1------------

### Deployment metadata ###
Name: frontend
Namespace: default
UID: ae654228-ed3f-4835-8316-ddc64ce81304
Creation time: 2020-11-25 09:30:51 +0800 HKT

### Deployment spec ###
Number of desired pods: 824636421628
Label selector - match labels: [app: guestbook] [tier: frontend] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: php-redis
Container image: gcr.io/google-samples/gb-frontend:v4
Container pull policy: IfNotPresent
Container ports: [containerPort: 80 ; Protocol: TCP] 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 3
Ready replicas: 3
Available replicas: 3
Unavailable replicas: 0

-------------Deployment 2------------

### Deployment metadata ###
Name: redis-master
Namespace: default
UID: c1140dec-1f37-496f-9af1-d0a862ef3e36
Creation time: 2020-11-25 09:22:58 +0800 HKT

### Deployment spec ###
Number of desired pods: 824636422108
Label selector - match labels: [app: redis] [role: master] [tier: backend] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: master
Container image: k8s.gcr.io/redis:e2e
Container pull policy: IfNotPresent
Container ports: [containerPort: 6379 ; Protocol: TCP] 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0

-------------Deployment 3------------

### Deployment metadata ###
Name: redis-slave
Namespace: default
UID: ece67caa-b25b-4780-83e0-6ec7cc788919
Creation time: 2020-11-25 09:26:58 +0800 HKT

### Deployment spec ###
Number of desired pods: 824636422588
Label selector - match labels: [app: redis] [role: slave] [tier: backend] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: slave
Container image: gcr.io/google_samples/gb-redisslave:v1
Container pull policy: IfNotPresent
Container ports: [containerPort: 6379 ; Protocol: TCP] 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 2
Ready replicas: 2
Available replicas: 2
Unavailable replicas: 0

-------------Deployment 4------------

### Deployment metadata ###
Name: event-exporter-gke
Namespace: kube-system
UID: f1fefe45-798f-40fe-99b9-b8a800785a42
Creation time: 2020-11-25 09:11:25 +0800 HKT

### Deployment spec ###
Number of desired pods: 824636423200
Label selector - match labels: [k8s-app: event-exporter] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: event-exporter
Container image: gke.gcr.io/event-exporter:v0.3.3-gke.0
Container pull policy: IfNotPresent
Container ports: 
---------------
Container name: prometheus-to-sd-exporter
Container image: gke.gcr.io/prometheus-to-sd:v0.10.0-gke.0
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0

-------------Deployment 5------------

### Deployment metadata ###
Name: fluentd-gke-scaler
Namespace: kube-system
UID: fb1965d0-4150-46f6-8969-f286476e0e3f
Creation time: 2020-11-25 09:11:29 +0800 HKT

### Deployment spec ###
Number of desired pods: 824636424040
Label selector - match labels: [k8s-app: fluentd-gke-scaler] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: fluentd-gke-scaler
Container image: k8s.gcr.io/fluentd-gcp-scaler:0.5.2
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0

-------------Deployment 6------------

### Deployment metadata ###
Name: kube-dns
Namespace: kube-system
UID: 5d0c3905-3829-4abb-bc6b-5897ccf2cc63
Creation time: 2020-11-25 09:11:25 +0800 HKT

### Deployment spec ###
Number of desired pods: 824634860920
Label selector - match labels: [k8s-app: kube-dns] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: kubedns
Container image: gke.gcr.io/k8s-dns-kube-dns-amd64:1.15.13
Container pull policy: IfNotPresent
Container ports: [containerPort: 10053 ; Protocol: UDP] [containerPort: 10053 ; Protocol: TCP] [containerPort: 10055 ; Protocol: TCP] 
---------------
Container name: dnsmasq
Container image: gke.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.15.13
Container pull policy: IfNotPresent
Container ports: [containerPort: 53 ; Protocol: UDP] [containerPort: 53 ; Protocol: TCP] 
---------------
Container name: sidecar
Container image: gke.gcr.io/k8s-dns-sidecar-amd64:1.15.13
Container pull policy: IfNotPresent
Container ports: [containerPort: 10054 ; Protocol: TCP] 
---------------
Container name: prometheus-to-sd
Container image: gke.gcr.io/prometheus-to-sd:v0.4.2
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: Default
Restart policy: Always

### Deployment status ###
Updated replicas: 2
Ready replicas: 2
Available replicas: 2
Unavailable replicas: 0

-------------Deployment 7------------

### Deployment metadata ###
Name: kube-dns-autoscaler
Namespace: kube-system
UID: 685197e2-f480-4dec-aecc-47efcb9ce78e
Creation time: 2020-11-25 09:11:24 +0800 HKT

### Deployment spec ###
Number of desired pods: 824634862648
Label selector - match labels: [k8s-app: kube-dns-autoscaler] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: autoscaler
Container image: gke.gcr.io/cluster-proportional-autoscaler-amd64:1.7.1-gke.0
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0

-------------Deployment 8------------

### Deployment metadata ###
Name: l7-default-backend
Namespace: kube-system
UID: 013998d0-b23a-4981-a1ff-36b8d04956ae
Creation time: 2020-11-25 09:11:24 +0800 HKT

### Deployment spec ###
Number of desired pods: 824634863256
Label selector - match labels: [k8s-app: glbc] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: default-http-backend
Container image: k8s.gcr.io/ingress-gce-404-server-with-metrics-amd64:v1.6.0
Container pull policy: IfNotPresent
Container ports: [containerPort: 8080 ; Protocol: TCP] 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0

-------------Deployment 9------------

### Deployment metadata ###
Name: metrics-server-v0.3.6
Namespace: kube-system
UID: 0316ccd4-17e4-4586-98f0-d536bffd4567
Creation time: 2020-11-25 09:11:25 +0800 HKT

### Deployment spec ###
Number of desired pods: 824634863872
Label selector - match labels: [k8s-app: metrics-server] [version: v0.3.6] 

### Pod specification of this deployment ###
---Containers in the pod---
Container name: metrics-server
Container image: k8s.gcr.io/metrics-server-amd64:v0.3.6
Container pull policy: IfNotPresent
Container ports: [containerPort: 443 ; Protocol: TCP] 
---------------
Container name: metrics-server-nanny
Container image: gke.gcr.io/addon-resizer:1.8.8-gke.1
Container pull policy: IfNotPresent
Container ports: 
---------------
DNS policy: ClusterFirst
Restart policy: Always

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0

-------------Deployment 10------------

### Deployment metadata ###
Name: stackdriver-metadata-agent-cluster-level
Namespace: kube-system
UID: ff31f3c8-dec7-48bb-80eb-c49bc71008c5
Creation time: 2020-11-25 09:11:25 +0800 HKT

### Deployment spec ###
Number of desired pods: 824634864872
Label selector - match labels: [app: stackdriver-metadata-agent] [cluster-level: true] 

### Pod specification of this deployment ###
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

### Deployment status ###
Updated replicas: 1
Ready replicas: 1
Available replicas: 1
Unavailable replicas: 0
```