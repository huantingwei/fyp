# Client libraries which will be used in this project

## 1. Go client library for talking to the Kubernetes master node API server

Link to client library: https://pkg.go.dev/k8s.io/client-go@v0.19.4/kubernetes

Link to GitHub repo: https://github.com/kubernetes/client-go

Version v0.19.4 is used

This library is sufficient for retrieving info of most kubernetes objects (e.g. node, pod, service...)

### Usage

Before using the library, configure your kubectl using the command below
```
gcloud container clusters get-credentials ${cluster name} --zone ${zone name} --project ${project name}
```

running kubebench.sh in the kubebench folder will do this configuration for you



## 2. Go client library for talking to GCP to retrieve cluster info (e.g. cluster, node pool)

Link to the client library: https://pkg.go.dev/cloud.google.com/go@v0.72.0/container/apiv1

Link to Github repo: https://github.com/googleapis/google-cloud-go

### usage

Follow the [Authorisation](https://github.com/googleapis/google-cloud-go) section of the Github repo.

I used the JSON key file method.

To generate the JSON key file, I followed this [guide](https://cloud.google.com/docs/authentication/production#manually)



## Sample output

### /api/v1/overview/cluster

```
{
    "Name": "cluster-1",
    "CreationTime": "2020-11-25T01:08:11+00:00",
    "MasterVersion": "1.16.13-gke.401",
    "IPendpoint": "104.155.167.236",
    "Location": "us-central1-c",
    "ReleaseChannel": 0,
    "Status": "RUNNING",
    "Network": "default",
    "NetworkConfig": "projects/fyp-gcp-296605/global/networks/default",
    "Subnet": "projects/fyp-gcp-296605/regions/us-central1/subnetworks/default",
    "IntranodeVisibility": false,
    "NetworkPolicyEnabled": false,
    "MasterAuthNetworkEnabled": false,
    "ShieldedNodeEnabled": false,
    "BinaryAuthorisationEnabled": false,
    "ClientCertificateEnabled": false
}
```

### /api/v1/overview/nodepool

```
{
    "count": 1,
    "data": [
        {
            "Name": "default-pool",
            "Version": "1.16.13-gke.401",
            "Location": "us-central1-c",
            "Status": 2,
            "AutoscalingEnabled": false,
            "InitialNodeCount": 3,
            "ImageType": "COS",
            "MachineType": "e2-medium",
            "DiskType": "pd-standard",
            "DiskSize": 100,
            "AutoUpgrade": true,
            "AutoRepair": true,
            "ServiceAccount": "default",
            "SecureBoot": false,
            "IntegrityMonitoring": true
        }
    ],
    "type": "nodepool"
}
```

### /api/v1/overview/deployment
```
{
    "count": 2,
    "data": [
        {
            "ObjectMeta": {
                "Name": "frontend",
                "Namespace": "default",
                "Uid": "ae654228-ed3f-4835-8316-ddc64ce81304",
                "CreationTime": "2020-11-25 09:30:51 +0800 HKT"
            },
            "DesiredPods": 3,
            "MatchLabels": {
                "app": "guestbook",
                "tier": "frontend"
            },
            "Containers": [
                {
                    "Name": "php-redis",
                    "Image": "gcr.io/google-samples/gb-frontend:v4",
                    "ImagePullPolicy": "IfNotPresent",
                    "ContainerPorts": {
                        "80": "TCP"
                    }
                }
            ],
            "DnsPolicy": "ClusterFirst",
            "RestartPolicy": "Always",
            "NodeName": "",
            "UpdatedReplicas": 3,
            "ReadyReplicas": 3,
            "AvailableReplicas": 3,
            "UnavailableReplicas": 0
        },
        {
            "ObjectMeta": {
                "Name": "redis-master",
                "Namespace": "default",
                "Uid": "c1140dec-1f37-496f-9af1-d0a862ef3e36",
                "CreationTime": "2020-11-25 09:22:58 +0800 HKT"
            },
            "DesiredPods": 1,
            "MatchLabels": {
                "app": "redis",
                "role": "master",
                "tier": "backend"
            },
            "Containers": [
                {
                    "Name": "master",
                    "Image": "k8s.gcr.io/redis:e2e",
                    "ImagePullPolicy": "IfNotPresent",
                    "ContainerPorts": {
                        "6379": "TCP"
                    }
                }
            ],
            "DnsPolicy": "ClusterFirst",
            "RestartPolicy": "Always",
            "NodeName": "",
            "UpdatedReplicas": 1,
            "ReadyReplicas": 1,
            "AvailableReplicas": 1,
            "UnavailableReplicas": 0
        }
    ],    
    "type": "deployment"
}
```

### /api/v1/overview/node

```
{
    "count": 2,
    "data": [
        {
            "ObjectMeta": {
                "Name": "gke-cluster-1-default-pool-c6af14e8-0zl9",
                "Namespace": "",
                "Uid": "7c094e7b-bda6-4c50-8490-9c74f1ee9add",
                "CreationTime": "2020-11-25 09:11:33 +0800 HKT"
            },
            "PodCIDR": "10.4.0.0/24",
            "NodeID": "gce://fyp-gcp-296605/us-central1-c/gke-cluster-1-default-pool-c6af14e8-0zl9",
            "MachineID": "2229178c881ccb7883cedb01d7c9175c",
            "KernelVersion": "4.19.112+",
            "OsImage": "Container-Optimized OS from Google",
            "Os": "linux",
            "ContainerRuntime": "docker://19.3.1",
            "KubeletVersion": "v1.16.13-gke.401",
            "KubeProxyVersion": "v1.16.13-gke.401",
            "CpuCap": 2,
            "MemoryCap": 4.140924928,
            "PodsCap": 110,
            "EphemeralStorageCap": 101.241290752,
            "StorageCap": 0,
            "CpuAllocatable": 1,
            "MemoryAllocatable": 2.967568384,
            "PodsAllocatable": 110,
            "EphemeralStorageAllocatable": 47.093746742,
            "StorageAllocatable": 0,
            "Conditions": [
                {
                    "ConditionName": "CorruptDockerOverlay2",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "docker overlay2 is functioning properly"
                },
                {
                    "ConditionName": "FrequentKubeletRestart",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "kubelet is functioning properly"
                },
                {
                    "ConditionName": "FrequentDockerRestart",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "docker is functioning properly"
                },
                {
                    "ConditionName": "FrequentContainerdRestart",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "containerd is functioning properly"
                },
                {
                    "ConditionName": "FrequentUnregisterNetDevice",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "node is functioning properly"
                },
                {
                    "ConditionName": "KernelDeadlock",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "kernel has no deadlock"
                },
                {
                    "ConditionName": "ReadonlyFilesystem",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:08 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:36 +0800 HKT",
                    "Message": "Filesystem is not read-only"
                },
                {
                    "ConditionName": "NetworkUnavailable",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-11-25 09:11:34 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "NodeController create implicit route"
                },
                {
                    "ConditionName": "MemoryPressure",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:16:27 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:33 +0800 HKT",
                    "Message": "kubelet has sufficient memory available"
                },
                {
                    "ConditionName": "DiskPressure",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:16:27 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:33 +0800 HKT",
                    "Message": "kubelet has no disk pressure"
                },
                {
                    "ConditionName": "PIDPressure",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:16:27 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:33 +0800 HKT",
                    "Message": "kubelet has sufficient PID available"
                },
                {
                    "ConditionName": "Ready",
                    "Status": "True",
                    "LastHeartbeatTime": "2020-12-04 15:16:27 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "kubelet is posting ready status. AppArmor enabled"
                }
            ]
        },
        {
            "ObjectMeta": {
                "Name": "gke-cluster-1-default-pool-c6af14e8-9fj1",
                "Namespace": "",
                "Uid": "f916d621-5315-4f4a-b722-71f7899f537e",
                "CreationTime": "2020-11-25 09:11:33 +0800 HKT"
            },
            "PodCIDR": "10.4.2.0/24",
            "NodeID": "gce://fyp-gcp-296605/us-central1-c/gke-cluster-1-default-pool-c6af14e8-9fj1",
            "MachineID": "33e842ad113db39ef6713764e180aa82",
            "KernelVersion": "4.19.112+",
            "OsImage": "Container-Optimized OS from Google",
            "Os": "linux",
            "ContainerRuntime": "docker://19.3.1",
            "KubeletVersion": "v1.16.13-gke.401",
            "KubeProxyVersion": "v1.16.13-gke.401",
            "CpuCap": 2,
            "MemoryCap": 4.140924928,
            "PodsCap": 110,
            "EphemeralStorageCap": 101.241290752,
            "StorageCap": 0,
            "CpuAllocatable": 1,
            "MemoryAllocatable": 2.967568384,
            "PodsAllocatable": 110,
            "EphemeralStorageAllocatable": 47.093746742,
            "StorageAllocatable": 0,
            "Conditions": [
                {
                    "ConditionName": "KernelDeadlock",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "kernel has no deadlock"
                },
                {
                    "ConditionName": "ReadonlyFilesystem",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "Filesystem is not read-only"
                },
                {
                    "ConditionName": "CorruptDockerOverlay2",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "docker overlay2 is functioning properly"
                },
                {
                    "ConditionName": "FrequentUnregisterNetDevice",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "node is functioning properly"
                },
                {
                    "ConditionName": "FrequentKubeletRestart",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "kubelet is functioning properly"
                },
                {
                    "ConditionName": "FrequentDockerRestart",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "docker is functioning properly"
                },
                {
                    "ConditionName": "FrequentContainerdRestart",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:14:54 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "containerd is functioning properly"
                },
                {
                    "ConditionName": "NetworkUnavailable",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-11-25 09:11:34 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "NodeController create implicit route"
                },
                {
                    "ConditionName": "MemoryPressure",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:15:44 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:33 +0800 HKT",
                    "Message": "kubelet has sufficient memory available"
                },
                {
                    "ConditionName": "DiskPressure",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:15:44 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:33 +0800 HKT",
                    "Message": "kubelet has no disk pressure"
                },
                {
                    "ConditionName": "PIDPressure",
                    "Status": "False",
                    "LastHeartbeatTime": "2020-12-04 15:15:44 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:33 +0800 HKT",
                    "Message": "kubelet has sufficient PID available"
                },
                {
                    "ConditionName": "Ready",
                    "Status": "True",
                    "LastHeartbeatTime": "2020-12-04 15:15:44 +0800 HKT",
                    "LastTransitionTime": "2020-11-25 09:11:34 +0800 HKT",
                    "Message": "kubelet is posting ready status. AppArmor enabled"
                }
            ]
        },
    ],
    "type": "node"
}
```

## Reminder
1. The gcloud ssh command is still in beta

2. Kubectl handles locating and authentication of the apiserver.


