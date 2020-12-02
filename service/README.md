## Services
A sample output containing the metadata, spec and status of each individual service.

May look at how to filter out system obejct later on.


## Sample output
```
---------------------Service 1---------------------

### Object metadata ###
Service name: frontend
Namespace: default
UID: 9f861630-efb2-4714-a124-6f318a98cd65
Creation time: 2020-11-25 09:35:09 +0800 HKT

### Service spec ###
Type: LoadBalancer
Cluster IP: 10.8.1.74
Label selector: [app: guestbook] [tier: frontend] 
List of service ports: 
[Port: 80 ; Node port: 30279 ; Target port: 80 ; Protocol: TCP]

### Load balancer status (if applicable) ###
Load balancer Ingress IP: 
Load balancer IP: 35.239.181.9

---------------------Service 2---------------------

### Object metadata ###
Service name: kubernetes
Namespace: default
UID: ed603b88-d5d3-407f-9f17-bc8a8441ca05
Creation time: 2020-11-25 09:11:10 +0800 HKT

### Service spec ###
Type: ClusterIP
Cluster IP: 10.8.0.1
Label selector: 
List of service ports: 
[Port: 443 ; Node port: 0 ; Target port: 443 ; Protocol: TCP]

### Load balancer status (if applicable) ###
Load balancer Ingress IP: 

---------------------Service 3---------------------

### Object metadata ###
Service name: redis-master
Namespace: default
UID: 5207ab9e-754d-4fc3-a0d6-adcac3218b98
Creation time: 2020-11-25 09:25:09 +0800 HKT

### Service spec ###
Type: ClusterIP
Cluster IP: 10.8.15.224
Label selector: [app: redis] [role: master] [tier: backend] 
List of service ports: 
[Port: 6379 ; Node port: 0 ; Target port: 6379 ; Protocol: TCP]

### Load balancer status (if applicable) ###
Load balancer Ingress IP: 

---------------------Service 4---------------------

### Object metadata ###
Service name: default-http-backend
Namespace: kube-system
UID: 5cee31df-99b5-4200-adcc-0ad0ad92da08
Creation time: 2020-11-25 09:11:24 +0800 HKT

### Service spec ###
Type: NodePort
Cluster IP: 10.8.15.28
Label selector: [k8s-app: glbc] 
List of service ports: 
[Port: 80 ; Node port: 31323 ; Target port: 8080 ; Protocol: TCP]

### Load balancer status (if applicable) ###
Load balancer Ingress IP: 

---------------------Service 5---------------------

### Object metadata ###
Service name: kube-dns
Namespace: kube-system
UID: aeae8edc-4f41-41c5-a559-e3bd70dd3b6c
Creation time: 2020-11-25 09:11:25 +0800 HKT

### Service spec ###
Type: ClusterIP
Cluster IP: 10.8.0.10
Label selector: [k8s-app: kube-dns] 
List of service ports: 
[Port: 53 ; Node port: 0 ; Target port: 53 ; Protocol: UDP]
[Port: 53 ; Node port: 0 ; Target port: 53 ; Protocol: TCP]

### Load balancer status (if applicable) ###
Load balancer Ingress IP: 

---------------------Service 6---------------------

### Object metadata ###
Service name: metrics-server
Namespace: kube-system
UID: fa70a0a5-77e9-46b0-8dd5-97d4b9e18455
Creation time: 2020-11-25 09:11:25 +0800 HKT

### Service spec ###
Type: ClusterIP
Cluster IP: 10.8.13.36
Label selector: [k8s-app: metrics-server] 
List of service ports: 
[Port: 443 ; Node port: 0 ; Target port: 0 ; Protocol: TCP]

### Load balancer status (if applicable) ###
Load balancer Ingress IP: 
```