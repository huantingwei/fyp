#!/bin/bash

sudo docker run -v $(pwd):/project zegl/kube-score:v1.10.0 score "./backend/kubescore/interactive-in.yaml" -o json > "./backend/kubescore/interactive-out.json"