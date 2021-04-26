#!/bin/bash

format="json"

sudo docker run -v $(pwd):/project zegl/kube-score:v1.10.0 score "./backend/kubescore/interactive-in.yaml" --enable-optional-test strings -o ${format} > "./backend/kubescore/interactive-out.json"