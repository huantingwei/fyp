#!/bin/bash

# consts
pwd="$(pwd)"
objectDir="kubescore/objs"
resultFile="kubescore/res.json"
resultDir="kubescore"
format="json"

rm -r ${objectDir}
mkdir -p ${objectDir}

# get all yaml files

# pass the name of namespace to scan as 1st parameter
#namespace=$1

# scan default namespace
namespace="default"

# while read namespace
# do

    echo "scanning namespace '${namespace}'"
    #mkdir -p "${objectDir}/${namespace}"
    while read -r resource
    do
        echo "  scanning resource '${resource}'"
        #mkdir -p "${objectDir}/${resource}"
        while read -r item
        do
            echo "    exporting item '${item}'"
            kubectl get "$resource" -n "$namespace" "$item" -o yaml > "${objectDir}/${resource}-${item}.yaml"
        done < <(kubectl get "$resource" -n "$namespace" 2>&1 | tail -n +2 | awk '{print $1}')
    done < <(kubectl api-resources --namespaced=true 2>/dev/null | tail -n +2 | awk '{print $1}')

# done < <(kubectl get namespaces | tail -n +2 | awk '{print $1}')

# run kubescore

cd ${objectDir}
sudo docker run --rm -v $(pwd):/project zegl/kube-score:v1.10.0 score *.yaml -o ${format} >  ../res.json
cd ${pwd}

# clean up
rm -r ${objectDir}
rm ${resultFile}