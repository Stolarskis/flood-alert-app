# ==================== [START] Global Variable Declaration =================== #
SHELL := /bin/bash
BASE_DIR := $(shell pwd)
CONTAINER_REG := gcr.io/flood-alert-app/app-flood-ctl
TEST_ENV := whatsup


push_latest:
	@docker build -t $(CONTAINER_REG):latest .
	@docker push $(CONTAINER_REG):latest
	
#TODO: Add a rollback target to delete the latest image and then add the latest tag to most recent
#https://ahmet.im/blog/google-container-registry-tips

