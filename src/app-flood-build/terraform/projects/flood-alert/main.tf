################################################################################
#                        GCP INFRASTRUCTURE DEFINITIONS                        #
################################################################################

provider "google" {
  project     = "${var.gcp_project_name}"
  region      = "${var.gcp_project_region}"
  credentials = "${file("../../../secrets/flood-alert-app/credentials.json")}"
}

resource "google_container_cluster" "primary" {
  name               = "${var.primary_cluster_name}"
  network            = "${var.primary_cluster_network}"
  zone               = "${var.primary_cluster_zone}"
  initial_node_count = "${var.primary_cluster_initial_node_count}"
}
