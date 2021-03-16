#!/bin/sh

gcloud beta cloud-shell ssh --authorize-session --command="cd ~"
gcloud beta cloud-shell ssh --authorize-session --command="docker run --rm -v ~:/host aquasec/kube-bench:latest install"

gcloud beta cloud-shell ssh --authorize-session --command="./kube-bench --benchmark gke-1.0 --json > kb_output.json"

gcloud beta cloud-shell scp cloudshell:~/kb_output.json  localhost:./backend/kubebench/kb_output.json

