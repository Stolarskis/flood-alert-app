#!/bin/bash

#Start cluster via terraform 
cd ~/go/src/gitlab.com/app-flood-scripts/helper-scripts
./start-cluster.sh

#Give tiller the right roles
kubectl create clusterrolebinding permissive-binding --clusterrole=cluster-admin --user=admin --user=kubelet --group=system:serviceaccounts;

#Setup helm install tiller 
helm init --wait
helm update

#Create Pods
cd ~/go/src/gitlab.com/flood-app-terraform
helm install --name flood-alert-app ./flood-alert-app