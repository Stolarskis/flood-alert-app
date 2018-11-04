#!/bin/bash

cd ~/go/src/gitlab.com/flood-app-terraform
terraform init
terraform plan -out myplan
terraform apply -auto-approve  myplan
gcloud container clusters get-credentials terraform-test-custer --zone us-central1-a --project flood-alert-app

