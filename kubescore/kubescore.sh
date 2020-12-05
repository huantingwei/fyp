#!/usr/bin/env bash

# consts
objectDir="./objects"
resultFile="./result.json"
resultDir="kubescore"
format="json"

mkdir -p ${objectDir}

# get all yaml files

# pass the name of namespace to scan as 1st parameter
namespace=$1

# while read namespace
# do
    echo "scanning namespace '${namespace}'"
    # mkdir -p "${objectDir}/${namespace}"
    while read -r resource
    do
        echo "  scanning resource '${resource}'"
        # mkdir -p "${objectDir}/${resource}"
        while read -r item
        do
            echo "    exporting item '${item}'"
            kubectl get "$resource" -n "$namespace" "$item" -o yaml > "${objectDir}/${resource}-${item}.yaml"
        done < <(kubectl get "$resource" -n "$namespace" 2>&1 | tail -n +2 | awk '{print $1}')
    done < <(kubectl api-resources --namespaced=true 2>/dev/null | tail -n +2 | awk '{print $1}')
# done < <(kubectl get namespaces | tail -n +2 | awk '{print $1}')

# run kubescore

run_kubescore = "sudo docker run -it -v $(pwd)/${objectDir}:~/${objectDir} zegl/kube-score:v1.7.0 score ~/${objectDir}/*.yaml -o $format > $(pwd)/$resultDir/$resultFile
