## Deployment
An example of printing the name of all deployments and their corresponding containers.

Link to client library: https://pkg.go.dev/k8s.io/client-go@v0.19.4/kubernetes#Clientset.AppsV1

## Usage
If kubectl is already configured, simply change the package name to main and run it.

## Sample output
```
--------Deployment 0: frontend--------
Name: php-redis ; Image: gcr.io/google-samples/gb-frontend:v4

--------Deployment 1: redis-master--------
Name: master ; Image: k8s.gcr.io/redis:e2e

--------Deployment 2: redis-slave--------
Name: slave ; Image: gcr.io/google_samples/gb-redisslave:v1

--------Deployment 3: event-exporter-gke--------
Name: event-exporter ; Image: gke.gcr.io/event-exporter:v0.3.3-gke.0
Name: prometheus-to-sd-exporter ; Image: gke.gcr.io/prometheus-to-sd:v0.10.0-gke.0

--------Deployment 4: fluentd-gke-scaler--------
Name: fluentd-gke-scaler ; Image: k8s.gcr.io/fluentd-gcp-scaler:0.5.2

--------Deployment 5: kube-dns--------
Name: kubedns ; Image: gke.gcr.io/k8s-dns-kube-dns-amd64:1.15.13
Name: dnsmasq ; Image: gke.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.15.13
Name: sidecar ; Image: gke.gcr.io/k8s-dns-sidecar-amd64:1.15.13
Name: prometheus-to-sd ; Image: gke.gcr.io/prometheus-to-sd:v0.4.2

--------Deployment 6: kube-dns-autoscaler--------
Name: autoscaler ; Image: gke.gcr.io/cluster-proportional-autoscaler-amd64:1.7.1-gke.0

--------Deployment 7: l7-default-backend--------
Name: default-http-backend ; Image: k8s.gcr.io/ingress-gce-404-server-with-metrics-amd64:v1.6.0

--------Deployment 8: metrics-server-v0.3.6--------
Name: metrics-server ; Image: k8s.gcr.io/metrics-server-amd64:v0.3.6
Name: metrics-server-nanny ; Image: gke.gcr.io/addon-resizer:1.8.8-gke.1

--------Deployment 9: stackdriver-metadata-agent-cluster-level--------
Name: metadata-agent ; Image: gcr.io/stackdriver-agents/metadata-agent-go:1.2.0
Name: metadata-agent-nanny ; Image: gke.gcr.io/addon-resizer:1.8.11-gke.1
```