# Flood Alert App

Ops Dependencies:

- Terraform
- Helm

## Directory Structure

- [Terraform Infrastructure Definitions](terraform/projects/flood-alert/main.tf)
- [Terraform Variable Definitions](terraform/projects/flood-alert/vars.tf)
- [Terraform Values](terraform/projects/flood-alert/flood-alert.tf)

```
.
├── [ 102]  helm/
│   └── [ 204]  flood-alert-app/
│       ├── [ 170]  templates/
│       │   ├── [ 449]  01.flood-alert-app.svc.yaml
│       │   ├── [ 656]  02.flood-alert-app.deployment.yaml
│       │   └── [ 700]  03.app-flood-ctl.cronjob.yaml
│       ├── [ 136]  Chart.yaml
│       └── [ 611]  values.yaml
├── [ 136]  secrets/
│   ├── [ 102]  examples/
│   │   └── [  40]  credentials.json
│   └── [  68]  production/
├── [ 170]  terraform/
│   ├── [  68]  plans/
│   ├── [ 102]  projects/
│   │   └── [ 170]  flood-alert/
│   │       ├── [ 474]  flood-alert.tfvars
│   │       ├── [ 708]  main.tf
│   │       └── [1.1K]  vars.tf
│   └── [ 102]  state/
│       └── [ 18K]  terraform.tfstate
├── [1.2K]  Makefile
└── [1.0K]  README.md
```

## Usage

A [Makefile](Makefile) is included to simplify deployment management and prepare for automated deployments (CI/CD) in the future. It assumes you already have `helm` and `terraform` installed on your host machine.

### Create the K8s Cluster with Terraform

```sh
make tf_init
make tf_plan
make tf_apply
```

### Destroy the K8s Cluster with Terraform

```sh
make tf_destroy
```

### Install Flood App into the Cluster

**Be sure to use the real static IP address you have reserved in GCP:** [Reserve Static Address](https://console.cloud.google.com/networking/addresses/list)

#### Initial installation

**NOTE:** In the commands below, `version` should match an image tag in [Google's container registry](https://console.cloud.google.com/gcr/images/) or you will get image pull errors. Do not use sha's, use normal [Semantic Versioning](https://semver.org/) for tagging images.

```sh
make helm_install ip_address='321.123.456.12' version='0.0.1'
```

#### Upgrade

```sh
make helm_upgrade ip_address='321.123.456.12' version='0.0.2'
```

### Remove Flood App from Cluster

```sh
make helm_delete
```
