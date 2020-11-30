#!/usr/bin/env bash

# consts
OBJECT_DIR="objects"
RESULT_FILE="result.json"
RESULT_DIR="kubescore"
FORMAT="json"

mkdir -p ${OBJECT_DIR}

# get all yaml files

# pass the name of namespace to scan as 1st parameter
namespace=$1

# while read namespace
# do
    echo "scanning namespace '${namespace}'"
    # mkdir -p "${OBJECT_DIR}/${namespace}"
    while read -r resource
    do
        echo "  scanning resource '${resource}'"
        # mkdir -p "${OBJECT_DIR}/${resource}"
        while read -r item
        do
            echo "    exporting item '${item}'"
            kubectl get "$resource" -n "$namespace" "$item" -o yaml > "${OBJECT_DIR}/${resource}-${item}.yaml"
        done < <(kubectl get "$resource" -n "$namespace" 2>&1 | tail -n +2 | awk '{print $1}')
    done < <(kubectl api-resources --namespaced=true 2>/dev/null | tail -n +2 | awk '{print $1}')
# done < <(kubectl get namespaces | tail -n +2 | awk '{print $1}')

# run kubescore

sudo docker run -it -v $(pwd)/${OBJECT_DIR}:/${OBJECT_DIR} zegl/kube-score:v1.7.0 score ${OBJECT_DIR}/*.yaml -o $FORMAT > $(pwd)/$RESULT_FILE