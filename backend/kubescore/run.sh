#!/bin/bash

# consts
pwd="$(pwd)"
objectDir="./backend/kubescore/objs"
resultFile="./backend/kubescore/res.json"
format="json"

rm -r ${objectDir}
mkdir -p ${objectDir}

# get all yaml files

while read namespace
do
    # skip auto-generated namespaces
    if [[ ${namespace} = "kube-system" || ${namespace} = "kube-node-lease" || ${namespace} = "kube-public" ]]; 
    then
        continue
    else
        echo "scanning namespace '${namespace}'"
        while read -r resource
        do
            echo "  scanning resource '${resource}'"
            if [ ${resource} = "events" ]
            then
                continue
            else
                while read -r item
                do
                    echo "    exporting item '${item}'"
                    kubectl get "$resource" -n "$namespace" "$item" -o yaml > "${objectDir}/${namespace}-${resource}-${item}.yaml"
                done < <(kubectl get "$resource" -n "$namespace" 2>&1 | tail -n +2 | awk '{print $1}')
            fi
        done < <(kubectl api-resources --namespaced=true 2>/dev/null | tail -n +2 | awk '{print $1}')
    fi
done < <(kubectl get namespaces | tail -n +2 | awk '{print $1}')

# run kubescore

cd ${objectDir}
sudo docker run -v --rm $(pwd):/project zegl/kube-score:v1.10.0 score *.yaml -o ${format} >  ../res.json
cd ${pwd}