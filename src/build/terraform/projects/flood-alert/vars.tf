################################################################################
#                           GCP VARIABLE DEFINITIONS                           #
################################################################################

variable "gcp_project_name" {
  type        = "string"
  description = "Name of Google Cloud Platform project."
}

variable "gcp_project_region" {
  type        = "string"
  description = "Primary region for Google Cloud Platform project."
  default     = "us-east1"
}

variable "primary_cluster_name" {
  type        = "string"
  description = "Primary region for Google Cloud Platform project."
  default     = "development"
}
variable "primary_cluster_network" {
  type        = "string"
  description = "Primary region for Google Cloud Platform project."
  default     = "default"
}
variable "primary_cluster_zone" {
  type        = "string"
  description = "Primary region for Google Cloud Platform project."
  default     = "us-east1-c"
}
variable "primary_cluster_initial_node_count" {
  type        = "string"
  description = "Initial node count for primary cluster."
  default     = "2"
}
