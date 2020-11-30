## main.go
contain a sample of using the Go client library for Kubernetes API (talking to the API server of master node)

Link to client library: https://pkg.go.dev/k8s.io/client-go@v0.19.4/kubernetes

Link to GitHub repo: https://github.com/kubernetes/client-go

Version v0.19.4 is used

This library is sufficient for retrieving info of most kubernetes objects

## usage
Before using the library, configure your kubectl using the command below
```
gcloud container clusters get-credentials ${cluster name} --zone ${zone name} --project ${project name}
```

## Reminder
1. The gcloud ssh command is still in beta

2. Kubectl handles locating and authentication of the apiserver.


