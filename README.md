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

The JSON response of visiting [this](http://localhost:8080/api/v1/overview/cluster) is:
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

## Reminder
1. The gcloud ssh command is still in beta

2. Kubectl handles locating and authentication of the apiserver.


