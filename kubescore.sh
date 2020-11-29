#!/usr/bin/env bash

# get all yaml files
while read -r namespace
do
    echo "scanning namespace '${namespace}'"
    mkdir -p "${HOME}/cluster-backup/${namespace}"
    while read -r resource
    do
        echo "  scanning resource '${resource}'"
        mkdir -p "${HOME}/cluster-backup/${namespace}/${resource}"
        while read -r item
        do
            echo "    exporting item '${item}'"
            kubectl get "$resource" -n "$namespace" "$item" -o yaml > "${HOME}/cluster-backup/${namespace}/${resource}/$item.yaml"
        done < <(kubectl get "$resource" -n "$namespace" 2>&1 | tail -n +2 | awk '{print $1}')
    done < <(kubectl api-resources --namespaced=true 2>/dev/null | tail -n +2 | awk '{print $1}')
done < <(kubectl get namespaces | tail -n +2 | awk '{print $1}')


# run this script in the root dir (e.g. fyp/)
# put all .yaml files in $OBJECT_FOLDER_NAME

RESULT_FILE="kubescore.json"
OBJECT_FOLDER_NAME="fixtures"
FORMAT="json"

sudo docker run -it -v $(pwd)/$OBJECT_FOLDER_NAME:/$OBJECT_FOLDER_NAME zegl/kube-score:v1.7.0 score $OBJECT_FOLDER_NAME/*.yaml -o $FORMAT > $RESULT_FILE