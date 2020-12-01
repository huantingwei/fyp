## Cluster
cluster.go contains an example of using the GCP client library to retrieve info about a cluster.

It is basically the cluster detail page on GCP console

## Sample output
```
------------Cluster general info------------
Name: cluster-1
Creation time: 2020-11-25T01:08:11+00:00
Version: 1.16.13-gke.401
End point: 104.155.167.236
Location: us-central1-c
Release channel: 0
Status: RUNNING
------------Cluster networking config------------
Network: default
Network config: projects/fyp-gcp-296605/global/networks/default
Subnet: projects/fyp-gcp-296605/regions/us-central1/subnetworks/default
Intranode visibility: false
Network policy enabled: false
Master authorised network enabled: false
------------Cluster security config------------
Shielded node enabled: false
Binary authorisation enabled: false
Client certificate enabled: false
```