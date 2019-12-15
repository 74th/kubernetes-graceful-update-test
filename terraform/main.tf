provider "google" {
  project = "nnyn-dev"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}
resource "google_compute_address" "lb_ip" {
  name = "graceful-test-server"
}
resource "google_container_cluster" "cluster" {
  name     = "graceful-cluster"
  location = "asia-northeast1-a"

  initial_node_count = 3

  node_config {
      preemptible = true
      machine_type = "n1-standard-2"

      oauth_scopes = [
        "https://www.googleapis.com/auth/cloud-platform"
      ]
  }
}
