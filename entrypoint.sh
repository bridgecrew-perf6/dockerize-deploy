#!/bin/bash
export AWS_ACCESS_KEY_ID=x 
export AWS_SECRET_ACCESS_KEY=x
export AWS_DEFAULT_REGION=eu-west-2

remap() {

    echo "no ns file: sleep \n"
    while ! [ -f ./ns-records.txt ];
    do
    sleep 2 # or less like 0.2
    done;
    echo "found ns file \n"
    sleep 1

    ns=`cat ns-records.txt`
    echo $ns

    ./remap gingerbreadtemplate.uk $ns
}

terraform init
terraform apply -auto-approve &
remap &
wait