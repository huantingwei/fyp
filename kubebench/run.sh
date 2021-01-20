#!/bin/sh

installkubebench="docker run --rm -v `pwd`:/host aquasec/kube-bench:latest install"
runkubebench="./kube-bench --benchmark gke-1.0 --json > kb_output.json"

gcloud beta cloud-shell ssh --command="$installkubebench"

gcloud beta cloud-shell ssh --command="$runkubebench"

gcloud beta cloud-shell scp cloudshell:~/kb_output.json  localhost:./kubebench/kb_output.json

