# ==================== [START] Global Variable Declaration =================== #
SHELL := /bin/bash
BASE_DIR := $(shell pwd)
UNAME_S := $(shell uname -s)

TF_PLANS_DIR := $(BASE_DIR)/terraform/plans
FLOOD_ALERT_TF_PLAN_FILE := $(TF_PLANS_DIR)/flood-alert.tfplan
FLOOD_ALERT_TF_DIR := $(BASE_DIR)/terraform/projects/flood-alert

FLOOD_ALERT_HELM_CHART := $(BASE_DIR)/helm/flood-alert-app

#Cluster info to connect gcloud to newly created cluster
GCP_PROJECT_NAME= "flood-alert-app"
PRIMARY_CLUSTER_NAME= "development"
PRIMARY_CLUSTER_ZONE = "us-east1-c"

export
# ===================== [END] Global Variable Declaration ==================== #

magic_button: tf_init tf_plan tf_apply helm_init helm_install 
	

tf_init:
	@cd $(FLOOD_ALERT_TF_DIR) && \
		terraform init -input=false

tf_plan: 
	@cd $(FLOOD_ALERT_TF_DIR) && \
		terraform plan -out=$(FLOOD_ALERT_TF_PLAN_FILE)

tf_apply: 
	@cd $(FLOOD_ALERT_TF_DIR) && \
		terraform apply $(FLOOD_ALERT_TF_PLAN_FILE)  
	@gcloud container clusters get-credentials $(PRIMARY_CLUSTER_NAME) --zone $(PRIMARY_CLUSTER_ZONE) --project $(GCP_PROJECT_NAME)

tf_destroy:
	@cd $(FLOOD_ALERT_TF_DIR) && \
		terraform destroy

tf_quick_destroy:
	@cd $(FLOOD_ALERT_TF_DIR) && \
		terraform destroy -auto-approve

# usage: make helm_install ip_address='321.123.456.12' version='latest'
helm_init: 
	@kubectl create clusterrolebinding permissive-binding --clusterrole=cluster-admin --user=admin --user=kubelet --group=system:serviceaccounts;
	@helm init --wait  

helm_install: 
	@helm install --name flood-app $(FLOOD_ALERT_HELM_CHART) \
		--set "floodAlert.service.loadBalancerIP=$$ip_address" \
		--set "floodAlert.image.version=$$version"

# usage: make helm_upgrade ip_address='321.123.456.12' version='latest'
helm_upgrade: 
	@helm upgrade flood-app $(FLOOD_ALERT_HELM_CHART) \
		--set "floodAlert.service.loadBalancerIP=$$ip_address" \
		--set "floodAlert.image.version=$$version"

helm_delete:
	@helm delete --purge flood-app 
	@kubectl delete pods --all
